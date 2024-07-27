package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/wellingtonreis/compras/pkg/response"

	models "github.com/wellingtonreis/compras/internal/app/models/purchases"
	"github.com/wellingtonreis/compras/internal/app/platform/database/mongodb"
)

func ListQuotationHistoryHandler(w http.ResponseWriter, r *http.Request) {

	parameter, err := io.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}
	r.Body.Close()

	var filter models.FilterQuotationHistory
	err = json.Unmarshal(parameter, &filter)
	if err != nil {
		w.Write([]byte("Erro ao deserializar o corpo da requisição"))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := mongodb.NewConnectionMongoDB(ctx)
	if err != nil {
		log.Fatal("Erro ao tentar se conectar ao mongodb:", err)
	}
	defer db.Close()
	result, err := models.SearchQuotationHistory(db, &filter)
	if err != nil {
		log.Fatalf("Erro ao tentar buscar os documentos: %v", err)
	}

	responseParams := response.ResponseParams{
		StatusCode: 200,
		Message:    "Operação realizada com sucesso.",
		Embedded:   result,
		Next:       "_",
		Total:      len(result),
	}

	jsonResponse := response.CreateJSONResponse(responseParams)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}
