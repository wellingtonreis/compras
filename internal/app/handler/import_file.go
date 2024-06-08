package handler

import (
	"compras/internal/app/models"
	"compras/internal/app/platform/database/mongodb"
	"compras/pkg/importer"
	dadosabertos "compras/pkg/service/compras/dados_abertos"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

func receiveAttachment(r *http.Request) (multipart.File, string, error) {
	r.ParseMultipartForm(10 << 20)

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		return nil, "", fmt.Errorf("falha ao receber o arquivo: %v", err)
	}
	defer file.Close()

	ext := filepath.Ext(fileHeader.Filename)

	return file, ext, nil
}

func returnNameFileTemp(ext string) (string, error) {
	var filenametmp = ""
	switch ext {
	case ".csv":
		filenametmp = "upload-*.csv"
	case ".xlsx":
		filenametmp = "upload-*.xlsx"
	default:
		return "", fmt.Errorf("formato de arquivo não suportado")
	}
	return filenametmp, nil
}

func createdFileTemp(filenametmp string, file multipart.File) (*os.File, error) {
	tempFile, err := os.CreateTemp("/tmp", filenametmp)
	if err != nil {
		return tempFile, fmt.Errorf("falha ao criar o arquivo temporário: %v", err)
	}

	_, err = io.Copy(tempFile, file)
	if err != nil {
		return tempFile, fmt.Errorf("falha ao copiar o arquivo: %v", err)
	}
	return tempFile, nil
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

func handlesData(data [][]string) *[]models.CatalogCode {
	convertedData := make([]models.CatalogCode, 0)
	api := dadosabertos.FnDadosAbertosComprasGov()

	uuid := uuid.New().String()
	for i, record := range data {
		if i == 0 {
			continue
		}

		result, err := api.ConsultarMaterial(record[0])
		if err != nil {
			log.Fatal("Error: ", err)
			panic("Erro ao tentar consultar os itens de compras")
		}

		item := models.CatalogCode{
			Catmat:       record[0],
			Apresentacao: record[1],
			Quantidade:   record[2],
			Cotacao:      uuid,
			DadosAPI:     result,
		}
		convertedData = append(convertedData, item)
	}
	return &convertedData
}

func saveData(data models.Data) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := mongodb.NewConnectionMongoDB(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer db.Close()

	var items []interface{}
	for _, item := range data.Catalog {
		items = append(items, item)
	}

	err = db.InsertData("admin", "purchases", items)
	if err != nil {
		log.Fatal(err)
	}
}

func findLatest() *[]models.CatalogCode {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := mongodb.NewConnectionMongoDB(ctx)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)
	}
	defer db.Close()
	catalogData, err := models.FilterDocumentsByPresentation(db)
	if err != nil {
		log.Fatalf("Failed to filter documents: %v", err)
	}
	return &catalogData
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {

	file, ext, err := receiveAttachment(r)
	if err != nil {
		log.Printf("Erro ao receber o arquivo: %v", err)
		http.Error(w, "Erro ao receber o arquivo", http.StatusInternalServerError)
		return
	}

	filenametmp, err := returnNameFileTemp(ext)
	if err != nil {
		log.Printf("Erro ao determinar o nome do arquivo temporário: %v", err)
		http.Error(w, "Erro ao determinar o nome do arquivo temporário", http.StatusInternalServerError)
		return
	}

	tempFile, err := createdFileTemp(filenametmp, file)
	defer os.Remove(tempFile.Name())
	defer tempFile.Close()

	if err != nil {
		log.Printf("Erro ao criar o arquivo temporário: %v", err)
		http.Error(w, "Erro ao criar o arquivo temporário", http.StatusInternalServerError)
		return
	}

	data, err := fileScanning(ext, tempFile.Name())
	if err != nil {
		log.Printf("Erro ao processar o arquivo: %v", err)
		http.Error(w, "Erro ao processar o arquivo", http.StatusInternalServerError)
		return
	}

	convertedData := handlesData(data)
	responseData := models.Data{
		Catalog: *convertedData,
	}

	saveData(responseData)
	catalogData := findLatest()

	filtered := models.Data{
		Catalog: *catalogData,
	}

	jsonData, err := json.Marshal(filtered)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
