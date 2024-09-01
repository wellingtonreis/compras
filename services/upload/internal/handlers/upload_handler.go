package handlers

import (
	"context"
	"fmt"

	"time"

	"compras/services/upload/config"
	"compras/services/upload/internal/services"

	fiber "github.com/gofiber/fiber/v2"
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

	objectName := file.Filename
	contentType := file.Header.Get("Content-Type")
	fileSize := file.Size

	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	svcUploadMinio := &services.UploadFileMinioService{
		Ctx:         context,
		ContentType: contentType,
		ObjectName:  objectName,
		OpenedFile:  openedFile,
		FileSize:    fileSize,
	}
	err = svcUploadMinio.UploadFile()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Falha ao enviar o arquivo para o MinIO")
	}

	svcReaderFileTemp := &services.ReaderFileTempService{
		File:       file,
		OpenedFile: openedFile,
	}
	data, err := svcReaderFileTemp.ReadFileTempService(context)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Falha ao ler o arquivo temporário")
	}

	sequence, err := services.SequenceQuotation()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Falha ao criar a sequência de cotação")
	}

	svcPublishMessages := &services.PublishMessagesRabbitMQService{
		Url:       config.Env.RabbitMQUrl,
		Quotation: sequence,
	}
	_, err = svcPublishMessages.PublishMessages(data)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString("Falha ao publicar as mensagens no RabbitMQ")
	}

	route := fmt.Sprintf("%s/quotation/%d/get", config.Env.GovbrUrl, sequence)
	agent := fiber.Get(route)
	status, _, errs := agent.Bytes()
	if len(errs) > 0 {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Erro ao acessar o serviço de consulta de catálogos",
		})
	}

	return ctx.Status(status).JSON(fiber.Map{
		"message": "Arquivo enviado com sucesso!",
		"file":    file.Filename,
	})
}
