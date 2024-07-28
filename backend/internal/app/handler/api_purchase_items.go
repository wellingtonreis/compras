package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/wellingtonreis/compras/pkg/response"

	"github.com/wellingtonreis/compras/internal/app/platform/database/mongodb"

	models "github.com/wellingtonreis/compras/internal/app/models/purchases"

	chi "github.com/go-chi/chi/v5"
)

func ListPurchaseItemsHandler(w http.ResponseWriter, r *http.Request) {
	quotationStr := chi.URLParam(r, "quotation")
	quotation, err := strconv.ParseInt(quotationStr, 10, 64)
	if err != nil {
		log.Fatalf("Erro ao converter a cotação para int64: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := mongodb.NewConnectionMongoDB(ctx)
	if err != nil {
		log.Fatal("Erro ao tentar se conectar ao mongodb:", err)
	}

	items, err := models.ListPurchaseItems(db, quotation)
	if err != nil {
		log.Fatalf("Erro ao tentar buscar os documentos: %v", err)
	}

	params := response.ResponseParams{
		StatusCode: 200,
		Message:    "Operação realizada com sucesso.",
		Embedded:   items,
		Next:       "_",
		Total:      len(items),
	}

	jsonResponse := response.CreateJSONResponse(params)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}

func UpdatePurchaseItemsHandler(w http.ResponseWriter, r *http.Request) {
	quotationStr := chi.URLParam(r, "quotation")
	quotation, err := strconv.ParseInt(quotationStr, 10, 64)
	if err != nil {
		log.Fatalf("Erro ao converter a cotação para int64: %v", err)
	}

	var items models.ItemPurchase
	err = json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		log.Fatalf("Erro ao decodificar o corpo da requisição: %v", err)
	}

	db := mongodb.ConnectToMongoDB()
	defer db.Close()

	err = models.UpdatePurchaseItems(db, quotation, &items)
	if err != nil {
		log.Fatalf("Erro ao tentar atualizar os documentos: %v", err)
	}

	params := response.ResponseParams{
		StatusCode: 200,
		Message:    "Operação realizada com sucesso.",
	}

	jsonResponse := response.CreateJSONResponse(params)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}

func DeletePurchaseItemsHandler(w http.ResponseWriter, r *http.Request) {
	quotationStr := chi.URLParam(r, "quotation")
	quotation, err := strconv.ParseInt(quotationStr, 10, 64)
	if err != nil {
		log.Fatalf("Erro ao converter a cotação para int64: %v", err)
	}
	fmt.Println(r.Body)
	var items models.ItemPurchase
	err = json.NewDecoder(r.Body).Decode(&items)
	if err != nil {
		log.Fatalf("Erro ao decodificar o corpo da requisição: %v", err)
	}

	db := mongodb.ConnectToMongoDB()
	defer db.Close()

	err = models.DeletePurchaseItems(db, quotation, &items)
	if err != nil {
		log.Fatalf("Erro ao tentar deletar os documentos: %v", err)
	}

	params := response.ResponseParams{
		StatusCode: 200,
		Message:    "Operação realizada com sucesso.",
	}

	jsonResponse := response.CreateJSONResponse(params)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}
