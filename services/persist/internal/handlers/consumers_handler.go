package handlers

import (
	"compras/services/persist/config"
	"compras/services/persist/internal/entity"
	"compras/services/persist/internal/services"
	"compras/services/persist/pkg/rabbitmq"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Consumer(ctx *fiber.Ctx) error {

	ch, err := rabbitmq.OpenChannel(config.RabbitMQUrl)
	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup

	channel := make(chan []entity.ItemPurchase)
	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs, "catmats")

	go func() {
		for msg := range msgs {
			msg.Ack(true)
			msgs <- msg
		}
	}()

	for {
		select {
		case msg := <-msgs:

			wg.Add(1)

			var itemPurchase entity.ItemPurchase
			if err := json.Unmarshal(msg.Body, &itemPurchase); err != nil {
				log.Fatalf("Erro ao desserializar a mensagem JSON: %v", err)
			}

			log.Println("Consultando cotação número: ", itemPurchase)
			go func() {
				defer wg.Done()

				svcSearchDataCatmat := services.SearchDataCatmat{
					Catmat: catalog.Catmat,
					Retry:  3,
				}
				result, err := svcSearchDataCatmat.Search()
				if err != nil {
					log.Printf("Erro: %v", err)
				}

				svcChannelDataItemPurchase := services.ChannelDataItemPurchase{
					Channel: channel,
					Result:  result,
				}
				go svcChannelDataItemPurchase.ChannelDataItemPurchase()
				item := <-svcChannelDataItemPurchase.Channel

				jsonData, err := json.Marshal(item)
				if err != nil {
					log.Fatalf("Falha ao converter a struct em JSON %v", err)
				}

				jsonString := string(jsonData)
				err = rabbitmq.Publish(ch, jsonString, config.RabbitMQExchange, config.RabbitMQQueue)
				if err != nil {
					log.Fatalf("Falha ao publicar o JSON no RabbitMQ %v", err)
				}
				log.Println("Finalizado o código de catálogo número: ", catalog.Catmat)
			}()
		case <-time.After(time.Second * 2):

			go func() {
				wg.Wait()
				ch.Close()
				fmt.Println("Conexão com rabbitmq fechada!")
			}()

			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Consumidor iniciado!",
			})
		}
	}
}
