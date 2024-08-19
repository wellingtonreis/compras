package models

import (
	"context"
	"log"
	"time"

	"github.com/wellingtonreis/compras/internal/app/platform/database/mongodb"
	"go.mongodb.org/mongo-driver/bson"
)

func ListSegments() ([]Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	db, err := mongodb.NewConnectionMongoDB(ctx)
	if err != nil {
		log.Fatal("Erro ao tentar se conectar ao mongodb:", err)
		return nil, err
	}
	defer db.Close()

	filter := bson.D{{}}
	var category []Category
	err = db.GetData("classification_segment", ctx, filter, &category)
	if err != nil {
		log.Fatalf("Erro ao tentar buscar os documentos: %v", err)
		return nil, err
	}

	return category, nil
}
