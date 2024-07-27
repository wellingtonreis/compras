package mongodb

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

func ConnectToMongoDB() *MongoDB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := NewConnectionMongoDB(ctx)
	if err != nil {
		log.Fatal("Erro ao tentar se conectar ao mongodb:", err)
	}
	return db
}

func NewConnectionMongoDB(ctx context.Context) (*MongoDB, error) {
	mongoURI := GetMongoDBURI()
	if mongoURI == "" {
		log.Fatal("MONGODB_URI não especificada na variável de ambiente")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	clientOptions.SetConnectTimeout(10 * time.Second)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	return &MongoDB{Client: client}, nil
}

func (db *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return db.Client.Disconnect(ctx)
}
