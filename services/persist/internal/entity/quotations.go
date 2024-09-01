package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type QuotationHistory struct {
	Catmat            string             `json:"catmat"`
	Apresentacao      string             `json:"apresentacao"`
	Quantidade        string             `json:"quantidade"`
	Cotacao           int64              `json:"cotacao"`
	Hu                string             `json:"hu"`
	Categoria         primitive.ObjectID `json:"categoria"`
	Subcategoria      primitive.ObjectID `json:"subcategoria"`
	DataHora          time.Time          `json:"datahora"`
	Situacao          string             `json:"situacao"`
	ProcessoSei       string             `json:"processosei"`
	Autor             string             `json:"autor"`
	DadosAPI          []ItemPurchase     `json:"dadosapi"`
	DadosConsolidados []ItemPurchase     `json:"dadosconsolidados"`
}
