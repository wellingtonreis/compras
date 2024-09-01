package services

import (
	"compras/services/upload/config"
	"compras/services/upload/internal/dto"
	"compras/services/upload/pkg/rabbitmq"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

type PublishMessagesRabbitMQService struct {
	Url       string
	Quotation int64
}

func (params *PublishMessagesRabbitMQService) PublishMessages(data [][]string) (int64, error) {

	ch, err := rabbitmq.OpenChannel(params.Url)
	if err != nil {
		return 400, fmt.Errorf("falha ao abrir o canal com rabbitmq: %v", err)
	}
	var wg sync.WaitGroup

	catmat := make(chan dto.Catmat, len(data))
	go func() {
		for i, record := range data {
			if i == 0 {
				continue
			}
			item := dto.Catmat{
				Quotation:    params.Quotation,
				Catmat:       record[0],
				Apresentacao: record[1],
				Quantidade:   record[2],
			}

			CreatePurchasesQuotation(&item)
			catmat <- item
		}
	}()

	for {
		select {
		case msg := <-catmat:

			wg.Add(1)

			go func() {
				defer wg.Done()

				jsonData, err := json.Marshal(msg)
				if err != nil {
					log.Printf("falha ao converter a struct em JSON %v", err)
				}

				jsonString := string(jsonData)
				err = rabbitmq.Publish(ch, jsonString, config.Env.RabbitMQExchange, config.Env.RabbitMQQueue)
				if err != nil {
					log.Printf("falha ao publicar o JSON no RabbitMQ %v", err)
				}
				log.Printf("Mensagem enviada: %v", jsonString)
			}()
		case <-time.After(time.Second * 2):

			go func() {
				wg.Wait()
				ch.Close()
				fmt.Println("Conexão com rabbitmq fechada!")
			}()

			return 200, nil
		}
	}
}
