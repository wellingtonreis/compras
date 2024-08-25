package handlers

import (
	"compras/services/govbr/config"
	"compras/services/govbr/internal/entity"
	"compras/services/govbr/internal/services"
	"compras/services/govbr/pkg/rabbitmq"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Consumer(ctx *fiber.Ctx) error {

	sequenceStr := ctx.Params("sequence")
	sequence, err := strconv.ParseInt(sequenceStr, 10, 64)
	if err != nil {
		log.Fatalf("Erro ao converter a sequência para inteiro: %v", err)
	}

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
			msgs <- msg
		}
	}()

	for {
		select {
		case msg := <-msgs:

			wg.Add(1)

			var catmat entity.Catmat
			if err := json.Unmarshal(msg.Body, &catmat); err != nil {
				log.Fatalf("Erro ao desserializar a mensagem JSON: %v", err)
			}

			if catmat.Quotation == sequence {
				go func() {
					defer wg.Done()

					svcSearchDataCatmat := services.SearchDataCatmat{
						Catmat: catmat.Catmat,
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
					log.Println("Finalizado o código de catálogo número: ", catmat.Catmat)
				}()

				msg.Ack(true)
			}
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
