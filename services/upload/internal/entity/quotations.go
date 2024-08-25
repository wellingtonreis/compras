package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuotationHistory struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Cotacao      int64              `json:"cotacao"`
	Hu           string             `json:"hu,omitempty"`
	Categoria    primitive.ObjectID `json:"categoria,omitempty"`
	Subcategoria primitive.ObjectID `json:"subcategoria,omitempty"`
	Situacao     string             `json:"situacao"`
	Processosei  string             `json:"processosei"`
	Autor        string             `json:"autor"`
	Datahora     time.Time          `json:"datahora,omitempty"`
}
