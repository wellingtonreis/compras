package response

import (
	"encoding/json"
	"log"
)

type Response struct {
	Error      string                 `json:"error,omitempty"`
	StatusCode int                    `json:"status_code"`
	Message    string                 `json:"message"`
	Embedded   map[string]interface{} `json:"embedded,omitempty"`
	Next       string                 `json:"next,omitempty"`
	Total      int                    `json:"total,omitempty"`
}

type ResponseParams struct {
	StatusCode int
	Message    string
	Error      string
	Embedded   map[string]interface{}
	Next       string
	Total      int
}

func CreateJSONResponse(params ResponseParams) string {
	response := Response{
		StatusCode: params.StatusCode,
		Message:    params.Message,
		Error:      params.Error,
		Embedded:   params.Embedded,
		Next:       params.Next,
		Total:      params.Total,
	}

	jsonData, err := json.Marshal(response)
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err)
	}

	return string(jsonData)
}
