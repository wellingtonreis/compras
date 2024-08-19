package handlers

import (
	"compras/services/govbr/config"
	"compras/services/govbr/internal/entity"
	"compras/services/govbr/pkg/rabbitmq"
	"encoding/json"
	"fmt"
	"log"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Consumer(ctx *fiber.Ctx) error {

	ch, err := rabbitmq.OpenChannel(config.RabbitMQUrl)
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs, "catmats")

	go func() {
		for msg := range msgs {
			msg.Ack(false)
			msgs <- msg
		}
	}()

	for {
		select {
		case msg := <-msgs: // rabbitmq
			var catalog entity.Catalog
			if err := json.Unmarshal(msg.Body, &catalog); err != nil {
				log.Fatalf("Erro ao desserializar a mensagem JSON: %v", err)
			}
			fmt.Printf("Mensagem recebida: %+v\n", catalog)

		case <-time.After(time.Second * 5):
			log.Println("Finalizando o consumo de mensagens do RabbitMQ")
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Catmats lidos com sucesso!",
			})
		}
	}
}
