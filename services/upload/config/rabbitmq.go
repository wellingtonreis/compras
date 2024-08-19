package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var RabbitMQUrl string

func RabbitMQ() {

	path_file_env := os.Getenv("PATH_ROOT")

	err := godotenv.Load(filepath.Join(path_file_env, ".env"))
	if err != nil {
		log.Fatal("Não foi possível carregar o arquivo .env de RabbitMQ: %v ", err)
	}

	RabbitMQUrl = os.Getenv("RABBITMQ_URL")
	if RabbitMQUrl == "" {
		RabbitMQUrl = "amqp://"
		log.Println("RabbitMQUrl não configurado")
	}
}
