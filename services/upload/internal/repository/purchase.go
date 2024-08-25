package repository

import (
	"compras/services/upload/internal/entity"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

var schema = "admin"

type MongoDB struct {
	Client *mongo.Client
}

func (db *MongoDB) Save(data *entity.QuotationHistory) error {

	collection := db.Client.Database(schema).Collection("purchases_quotation")
	_, err := collection.InsertOne(context.Background(), data)
	if err != nil {
		return err
	}
	return nil
}
