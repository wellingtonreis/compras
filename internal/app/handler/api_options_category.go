package handler

import (
	"net/http"

	"github.com/wellingtonreis/compras/pkg/response"

	models "github.com/wellingtonreis/compras/internal/app/models/classification"
)

func ListOptionsCategoryHandler(w http.ResponseWriter, r *http.Request) {

	result, err := models.ListSegments()
	if err != nil {
		w.Write([]byte("Erro ao tentar buscar os documentos: %v"))
		return
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
