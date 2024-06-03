package handler

import (
	"compras/pkg/response"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	params := response.ResponseParams{
		StatusCode: 200,
		Message:    "Operação realizada com sucesso.",
		Embedded: map[string]interface{}{
			"user": map[string]string{"name": "John Doe", "role": "admin"},
		},
		Next:  "http://api.example.com/users?page=2",
		Total: 150,
	}

	jsonResponse := response.CreateJSONResponse(params)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(jsonResponse))
}
