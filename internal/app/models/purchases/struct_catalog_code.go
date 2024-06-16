package models

import "time"

type CatalogCode struct {
	Catmat            string         `json:"catmat"`
	Apresentacao      string         `json:"apresentacao"`
	Quantidade        string         `json:"quantidade"`
	Cotacao           string         `json:"cotacao"`
	Hu                string         `json:"hu"`
	Categoria         string         `json:"categoria"`
	Subcategoria      string         `json:"subcategoria"`
	DataHora          time.Time      `json:"datahora"`
	Situacao          string         `json:"situacao"`
	ProcessoSei       string         `json:"processosei"`
	Autor             string         `json:"autor"`
	DadosAPI          []ItemPurchase `json:"dadosapi"`
	DadosConsolidados []ItemPurchase `json:"dadosconsolidados"`
}
