package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Justification struct {
	ID           primitive.ObjectID `json:"_id,omitempty"`
	Descricao    string             `json:"descricao"`
	Data         string             `json:"data"`
	Autor        string             `json:"autor"`
	Valor        *float64           `json:"valor,omitempty"`
	ValorInicial *float64           `json:"valorInicial,omitempty"`
	DeleteAt     *time.Time         `json:"deleteAt,omitempty"`
}
