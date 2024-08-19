package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FilterQuotationHistory struct {
	QuotationHistory
	DataInicio string `json:"data_inicio"`
	DataFim    string `json:"data_fim"`
}

type QuotationHistory struct {
	Cotacao      int64              `json:"cotacao"`
	Hu           string             `json:"hu,omitempty"`
	Categoria    primitive.ObjectID `json:"categoria,omitempty"`
	Subcategoria primitive.ObjectID `json:"subcategoria,omitempty"`
	Situacao     string             `json:"situacao"`
	Processosei  string             `json:"processosei"`
	Autor        string             `json:"autor"`
	Datahora     time.Time          `json:"datahora,omitempty"`
}
