package handler

import (
	"context"
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
