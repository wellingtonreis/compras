package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Subcategory struct {
	Name string `bson:"name" json:"name"`
}

type Category struct {
	Name          string        `bson:"name" json:"name"`
	Subcategories []Subcategory `bson:"subcategories" json:"subcategories"`
}

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Conectado ao MongoDB!")

	collection := client.Database("admin").Collection("classification_segment")

	fileCsv, err := os.Open("payload.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fileCsv.Close()

	readerCsv := csv.NewReader(fileCsv)
	readerCsv.Comma = '|'

	rows, err := readerCsv.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	categoryMap := make(map[string]*Category)

	for _, row := range rows[1:] {
		idCategoria := row[3]
		nomeCategoria := row[4]
		nomeSubcategoria := row[1]

		if _, exists := categoryMap[idCategoria]; !exists {
			categoryMap[idCategoria] = &Category{
				Name:          nomeCategoria,
				Subcategories: []Subcategory{},
			}
		}

		categoryMap[idCategoria].Subcategories = append(categoryMap[idCategoria].Subcategories, Subcategory{
			Name: nomeSubcategoria,
		})
	}

	for _, cat := range categoryMap {
		_, err := collection.InsertOne(context.TODO(), cat)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Categoria inserida com sucesso: %+v\n", cat)
	}

	fmt.Println("Inserção concluída!")
}
