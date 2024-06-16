package models

import "time"

type QuotationHistory struct {
	Cotacao      string    `json:"cotacao"`
	Hu           string    `json:"hu"`
	Categoria    string    `json:"categoria"`
	Subcategoria string    `json:"subcategoria"`
	Datahora     time.Time `json:"datahora"`
	Situacao     string    `json:"situacao"`
	Processosei  string    `json:"processosei"`
	Autor        string    `json:"autor"`
}
