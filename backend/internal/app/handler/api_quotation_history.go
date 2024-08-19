package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/wellingtonreis/compras/pkg/response"
	"go.mongodb.org/mongo-driver/bson/primitive"

	models "github.com/wellingtonreis/compras/internal/app/models/quotation_history"
	"github.com/wellingtonreis/compras/internal/app/platform/database/mongodb"

	chi "github.com/go-chi/chi/v5"
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

func UpdateClassificationSegment(w http.ResponseWriter, r *http.Request) {
	quotationStr := chi.URLParam(r, "quotation")
	quotation, err := strconv.ParseInt(quotationStr, 10, 64)
	if err != nil {
		log.Fatalf("Erro ao converter a cotação para int64: %v", err)
	}

	var params map[string]json.RawMessage
	err = json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		log.Fatalf("Erro ao decodificar o corpo da requisição: %v", err)
	}

	var quotation_str string
	err = json.Unmarshal(params["cotacao"], &quotation_str)
	if err != nil {
		log.Fatalf("Erro ao deserializar a cotação: %v", err)
	}

	var category_id string
	err = json.Unmarshal(params["categoria"], &category_id)
	if err != nil {
		log.Fatalf("Erro ao deserializar a categoria: %v", err)
	}

	var subcategory_id string
	err = json.Unmarshal(params["subcategoria"], &subcategory_id)
	if err != nil {
		log.Fatalf("Erro ao deserializar a subcategoria: %v", err)
	}

	category, err := primitive.ObjectIDFromHex(category_id)
	if err != nil {
		log.Fatalf("Erro ao converter a categoria para ObjectID: %v", err)
	}

	subcategory, err := primitive.ObjectIDFromHex(subcategory_id)
	if err != nil {
		log.Fatalf("Erro ao converter a subcategoria para ObjectID: %v", err)
	}

	quotation_number, err := strconv.ParseInt(quotation_str, 10, 64)
	if err != nil {
		log.Fatalf("Erro ao converter a cotação para int64: %v", err)
	}

	params["categoria"] = json.RawMessage(`"` + category.Hex() + `"`)
	params["subcategoria"] = json.RawMessage(`"` + subcategory.Hex() + `"`)
	params["cotacao"] = json.RawMessage(strconv.FormatInt(quotation_number, 10))

	var items models.QuotationHistory
	rawData, err := json.Marshal(params)
	if err != nil {
		log.Fatalf("Erro ao serializar os dados: %v", err)
	}

	err = json.Unmarshal(rawData, &items)
	if err != nil {
		log.Fatalf("Erro ao deserializar os dados: %v", err)
	}

	db := mongodb.ConnectToMongoDB()
	defer db.Close()

	err = models.UpdateQuotationHistorySegment(db, quotation, &items)
	if err != nil {
		log.Fatalf("Erro ao tentar atualizar os documentos: %v", err)
	}

	responseParams := response.ResponseParams{
		StatusCode: 200,
		Message:    "Operação realizada com sucesso.",
	}

	jsonResponse := response.CreateJSONResponse(responseParams)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}
