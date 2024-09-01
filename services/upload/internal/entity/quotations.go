package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuotationHistory struct {
	ID                primitive.ObjectID `json:"_id,omitempty"`
	Catmat            string             `json:"catmat,omitempty"`
	Apresentacao      string             `json:"apresentacao,omitempty"`
	Quantidade        string             `json:"quantidade,omitempty"`
	Cotacao           int64              `json:"cotacao"`
	Hu                string             `json:"hu,omitempty"`
	Categoria         primitive.ObjectID `json:"categoria,omitempty"`
	Subcategoria      primitive.ObjectID `json:"subcategoria,omitempty"`
	DataHora          time.Time          `json:"datahora,omitempty"`
	Situacao          string             `json:"situacao"`
	ProcessoSei       string             `json:"processosei,omitempty"`
	Autor             string             `json:"autor,omitempty"`
	DadosAPI          []ItemPurchase     `json:"dadosapi,omitempty"`
	DadosConsolidados []ItemPurchase     `json:"dadosconsolidados,omitempty"`
}
