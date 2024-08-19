package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"compras/services/upload/config"
	"compras/services/upload/internal/entity"

	"compras/services/upload/pkg/importer"
	"compras/services/upload/pkg/rabbitmq"

	pkgMinio "compras/services/upload/pkg/minio"

	fiber "github.com/gofiber/fiber/v2"
	minio "github.com/minio/minio-go/v7"
)

func Upload(ctx *fiber.Ctx) error {

	file, err := ctx.FormFile("file")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	openedFile, err := file.Open()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Falha ao abrir o arquivo")
	}
	defer openedFile.Close()

	ext := filepath.Ext(file.Filename)
	objectName := file.Filename
	contentType := file.Header.Get("Content-Type")
	fileSize := file.Size

	log.Println("Iniciando conexão com MinIO...")
	err = pkgMinio.InitMinio()
	if err != nil {
		log.Printf("Falha ao inicializar o MinIO: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Falha ao inicializar o MinIO ou o bucket não existe")
	}

	log.Println("Realizando upload do arquivo...")
	_, err = pkgMinio.MinioClient.PutObject(context.Background(), config.MinioBucket, objectName, openedFile, fileSize, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Printf("Falha ao enviar o arquivo para o MinIO: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Falha ao enviar o arquivo para o MinIO")
	}
	log.Println("Upload do arquivo realizado com sucesso!")

	log.Println("Abrindo conexão com o RabbitMQ...")
	ch, err := rabbitmq.OpenChannel(config.RabbitMQUrl)
	if err != nil {
		log.Printf("Falha ao abrir o canal: %v", err)
		return ctx.Status(fiber.StatusInternalServerError).SendString("Falha ao abrir o canal")
	}
	defer ch.Close()

	log.Println("Criando arquivo temporário...")
	tempFile, err := createdFileTemp(file.Filename, openedFile)
	if err != nil {
		log.Println("Falha ao criar o arquivo temporário")
		return ctx.Status(fiber.StatusInternalServerError).SendString("Falha ao criar o arquivo temporário")
	}
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	log.Println("Escaneando o arquivo...")
	data, err := fileScanning(ext, tempFile.Name())
	if err != nil {
		log.Println("Falha ao escanear o arquivo")
		return ctx.Status(fiber.StatusInternalServerError).SendString("Falha ao escanear o arquivo")
	}

	log.Println("Enviando mensagens para o RabbitMQ...")
	catmat := make(chan entity.Catmat, len(data))
	go func() {
		for i, record := range data {
			if i == 0 {
				continue
			}
			item := entity.Catmat{
				Catmat:       record[0],
				Apresentacao: record[1],
				Quantidade:   record[2],
			}
			catmat <- item
		}
	}()

	for {
		select {
		case msg := <-catmat: // rabbitmq
			jsonData, err := json.Marshal(msg)
			if err != nil {
				log.Fatalf("Falha ao converter a struct em JSON %v", err)
			}

			jsonString := string(jsonData)
			err = rabbitmq.Publish(ch, jsonString, "amq.direct")
			if err != nil {
				log.Fatalf("Falha ao publicar o JSON no RabbitMQ %v", err)
			}
			log.Printf("Mensagem enviada para o RabbitMQ: %v", jsonString)

		case <-time.After(time.Second * 5):
			log.Println("Finalizando envio de mensagens para o RabbitMQ")
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"message": "Arquivo enviado com sucesso!",
				"file":    file.Filename,
			})
		}
	}
}

func fileScanning(ext string, tempFileName string) ([][]string, error) {
	var data [][]string
	var err error

	switch ext {
	case ".csv":
		data, err = importer.ReadCSVFile(tempFileName)
	case ".xlsx":
		data, err = importer.ReadFileXLSX(tempFileName)
	default:
		err = fmt.Errorf("formato de arquivo não suportado")
	}

	if err != nil {
		return data, fmt.Errorf("falha ao copiar o arquivo: %v", err)
	}

	return data, nil
}

func createdFileTemp(filenametmp string, file multipart.File) (*os.File, error) {
	tempFile, err := os.CreateTemp("/tmp", filenametmp)
	if err != nil {
		return nil, fmt.Errorf("falha ao criar o arquivo temporário: %v", err)
	}

	_, err = io.Copy(tempFile, file)
	if err != nil {
		tempFile.Close()
		os.Remove(tempFile.Name())
		return nil, fmt.Errorf("falha ao copiar o arquivo: %v", err)
	}

	if err := tempFile.Close(); err != nil {
		return nil, fmt.Errorf("falha ao fechar o arquivo temporário: %v", err)
	}

	return tempFile, nil
}
