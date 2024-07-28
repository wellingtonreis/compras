package handler

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	models "github.com/wellingtonreis/compras/internal/app/models/purchases"

	date_custom "github.com/wellingtonreis/compras/pkg/date_custom"
	dadosabertos "github.com/wellingtonreis/compras/pkg/service/compras/dados_abertos"

	"github.com/wellingtonreis/compras/internal/app/platform/database/mongodb"
	"github.com/wellingtonreis/compras/pkg/importer"
	"github.com/wellingtonreis/compras/pkg/response"
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

func handlesData(data [][]string) (*[]models.CatalogCode, int64) {
	convertedData := make([]models.CatalogCode, 0)
	api := dadosabertos.FnDadosAbertosComprasGov()

	db := mongodb.ConnectToMongoDB()
	defer db.Close()
	sequence, err := db.GetNextSequenceValue("catalogcode")
	if err != nil {
		log.Fatal("Erro ao tentar cadastrar a sequencia de identificação:", err)
	}

	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(data) - 1)

	today := date_custom.GetToday()
	for i, record := range data {
		if i == 0 {
			continue
		}

		go func() {
			result, err := api.ConsultarMaterial(record[0])
			if err != nil {
				log.Fatal("Error: ", err)
				panic("Erro ao tentar consultar os itens de compras")
			}

			item := models.CatalogCode{
				Catmat:       record[0],
				Apresentacao: record[1],
				Quantidade:   record[2],
				Cotacao:      sequence,
				Hu:           "",
				Categoria:    "",
				Subcategoria: "",
				DataHora:     today,
				Situacao:     "Iniciada",
				ProcessoSei:  "",
				Autor:        "",
				DadosAPI:     result,
			}
			convertedData = append(convertedData, item)
			waitGroup.Done()
		}()
	}
	waitGroup.Wait()
	return &convertedData, sequence
}

func saveData(data models.Data) {
	db := mongodb.ConnectToMongoDB()
	defer db.Close()
	var items []interface{}
	for _, item := range data.Catalog {
		items = append(items, item)
	}

	db.InsertData("purchases", items)
}

func findLatest(cotacao int64) *[]models.CatalogCode {
	db := mongodb.ConnectToMongoDB()
	defer db.Close()
	catalogData, err := models.SavePurchaseItemDocuments(db, cotacao)
	if err != nil {
		log.Fatalf("Erro ao tentar cadastrar os documentos: %v", err)
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

	convertedData, sequence := handlesData(data)
	responseData := models.Data{
		Catalog: *convertedData,
	}

	saveData(responseData)
	catalogData := findLatest(sequence)

	filtered := models.Data{
		Catalog: *catalogData,
	}

	params := response.ResponseParams{
		StatusCode: 200,
		Message:    "Operação realizada com sucesso.",
		Embedded:   filtered,
		Next:       "_",
		Total:      len(filtered.Catalog),
	}

	jsonResponse := response.CreateJSONResponse(params)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}
