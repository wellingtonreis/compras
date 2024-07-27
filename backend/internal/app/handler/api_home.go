package handler

import (
	"net/http"

	"github.com/wellingtonreis/compras/pkg/response"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	params := response.ResponseParams{
		StatusCode: 200,
		Message:    "Operação realizada com sucesso.",
		Embedded:   "Teste 1",
		Next:       "http://api.example.com/users?page=2",
		Total:      0,
	}

	jsonResponse := response.CreateJSONResponse(params)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}
