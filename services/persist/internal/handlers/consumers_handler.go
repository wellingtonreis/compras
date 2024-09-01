package handlers

import (
	"compras/services/persist/configs"
	"compras/services/persist/internal/dto"
	"compras/services/persist/internal/services"
	"compras/services/persist/pkg/rabbitmq"
	"encoding/json"
	"log"
	"time"

	fiber "github.com/gofiber/fiber/v2"
	amqp "github.com/rabbitmq/amqp091-go"
)

func Consumer(ctx *fiber.Ctx) error {

	ch, err := rabbitmq.OpenChannel(configs.Env.RabbitMQUrl)
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs, "catalog")

	for {
		select {
		case msg := <-msgs:

			msg.Ack(false)
			var ItemPurchaseMessage dto.ItemPurchaseMessage
			if err := json.Unmarshal(msg.Body, &ItemPurchaseMessage); err != nil {
				log.Fatalf("Erro ao desserializar a mensagem JSON: %v", err)
			}
			services.Save(&ItemPurchaseMessage)

		case <-time.After(time.Second * 2):
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Consumidor iniciado!",
			})
		}
	}

}
