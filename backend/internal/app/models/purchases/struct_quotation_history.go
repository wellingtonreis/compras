package models

import (
	"time"
)

type FilterQuotationHistory struct {
	Cotacao      int64  `json:"cotacao"`
	Hu           string `json:"hu"`
	Categoria    string `json:"categoria"`
	Subcategoria string `json:"subcategoria"`
	Situacao     string `json:"situacao"`
	Processosei  string `json:"processosei"`
	Autor        string `json:"autor"`
	DataInicio   string `json:"data_inicio"`
	DataFim      string `json:"data_fim"`
}

type QuotationHistory struct {
	Cotacao      int64     `json:"cotacao"`
	Hu           string    `json:"hu"`
	Categoria    string    `json:"categoria"`
	Subcategoria string    `json:"subcategoria"`
	Datahora     time.Time `json:"datahora"`
	Situacao     string    `json:"situacao"`
	Processosei  string    `json:"processosei"`
	Autor        string    `json:"autor"`
}
