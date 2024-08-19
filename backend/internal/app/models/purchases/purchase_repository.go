package models

import (
	"context"
	"fmt"
	"time"

	"github.com/wellingtonreis/compras/internal/app/platform/database/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var schema = mongodb.GetMongoSchema()

func SavePurchaseItemDocuments(db *mongodb.MongoDB, cotacao int64) ([]CatalogCode, error) {
	collection := db.Client.Database(schema).Collection("purchases")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	now := time.Now().UTC()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.FixedZone("", -3*60*60))

	threeMonthsAgo := today.AddDate(0, -3, 0)
	sixMonthsAgo := today.AddDate(0, -6, 0)
	twelveMonthsAgo := today.AddDate(0, -12, 0)

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: bson.M{"cotacao": cotacao}}},
		{{Key: "$match", Value: bson.M{"apresentacao": bson.M{"$exists": true}}}},
		{{Key: "$addFields", Value: bson.D{
			{Key: "dadosconsolidados", Value: bson.M{"$filter": bson.M{
				"input": "$dadosapi",
				"as":    "item",
				"cond":  bson.M{"$eq": []interface{}{"$$item.nomeunidadefornecimento", "$apresentacao"}},
			}}},
		}}},
		{{Key: "$match", Value: bson.M{"dadosconsolidados": bson.M{"$ne": []interface{}{}}}}},
		{{Key: "$addFields", Value: bson.D{
			{Key: "dadosconsolidados3meses", Value: bson.M{"$filter": bson.M{
				"input": "$dadosconsolidados",
				"as":    "item",
				"cond": bson.M{
					"$and": []interface{}{
						bson.M{"$gte": []interface{}{"$$item.datacompra", threeMonthsAgo}},
					},
				},
			}}},
		}}},
		{{Key: "$addFields", Value: bson.D{
			{Key: "dadosconsolidados6meses", Value: bson.M{
				"$cond": bson.M{
					"if": bson.M{"$lt": []interface{}{bson.M{"$size": bson.M{"$ifNull": []interface{}{"$dadosconsolidados3meses", []interface{}{}}}}, 6}},
					"then": bson.M{"$filter": bson.M{
						"input": "$dadosconsolidados",
						"as":    "item",
						"cond": bson.M{
							"$and": []interface{}{
								bson.M{"$gte": []interface{}{"$$item.datacompra", sixMonthsAgo}},
							},
						},
					}},
					"else": "$dadosconsolidados3meses",
				},
			}},
		}}},
		{{Key: "$addFields", Value: bson.D{
			{Key: "dadosconsolidados12meses", Value: bson.M{
				"$cond": bson.M{
					"if": bson.M{"$lt": []interface{}{bson.M{"$size": bson.M{"$ifNull": []interface{}{"$dadosconsolidados6meses", []interface{}{}}}}, 6}},
					"then": bson.M{"$filter": bson.M{
						"input": "$dadosconsolidados",
						"as":    "item",
						"cond": bson.M{
							"$gte": []interface{}{"$$item.datacompra", twelveMonthsAgo},
						},
					}},
					"else": "$dadosconsolidados6meses",
				},
			}},
		}}},
		{{Key: "$addFields", Value: bson.D{
			{Key: "dadosconsolidados", Value: "$dadosconsolidados12meses"},
		}}},
		{{Key: "$project", Value: bson.M{
			"dadosconsolidados3meses":  0,
			"dadosconsolidados6meses":  0,
			"dadosconsolidados12meses": 0,
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar a consulta de agregação: %v", err)
	}
	defer cursor.Close(ctx)

	var updatedDocuments []CatalogCode
	for cursor.Next(ctx) {
		var doc CatalogCode
		if err := cursor.Decode(&doc); err != nil {
			return nil, fmt.Errorf("erro ao decodificar o documento: %v", err)
		}

		for i := range doc.DadosConsolidados {
			doc.DadosConsolidados[i].ID = primitive.NewObjectID()
			doc.DadosConsolidados[i].Justificativa = []Justification{}
		}

		filter := bson.M{"catmat": doc.Catmat, "cotacao": doc.Cotacao}
		update := bson.M{
			"$set": bson.M{
				"dadosapi":          nil,
				"dadosconsolidados": doc.DadosConsolidados,
			},
		}

		opts := options.Update().SetUpsert(true)
		if _, err := collection.UpdateOne(ctx, filter, update, opts); err != nil {
			return nil, fmt.Errorf("erro ao tentar atualizar o documento: %v", err)
		}

		updatedDocuments = append(updatedDocuments, doc)
	}

	return updatedDocuments, nil
}

func ListPurchaseItems(db *mongodb.MongoDB, quotation int64) ([]ItemPurchase, error) {
	collection := db.Client.Database(schema).Collection("purchases")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{
		{{"$match", bson.D{{"cotacao", quotation}, {"dadosconsolidados", bson.D{{"$ne", nil}}}}}},
		{{"$project", bson.D{{"_id", 0}, {"dadosconsolidados", 1}}}},
		{{"$unwind", "$dadosconsolidados"}},
		{{"$replaceRoot", bson.D{{"newRoot", "$dadosconsolidados"}}}},
		{{"$match", bson.D{{"deleteat", bson.D{{"$eq", nil}}}}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar a consulta de agregação: %v", err)
	}
	defer cursor.Close(ctx)

	var docs []ItemPurchase
	if err = cursor.All(ctx, &docs); err != nil {
		return nil, fmt.Errorf("erro ao tentar ler todos os documentos: %v", err)
	}

	return docs, nil

}

func UpdatePurchaseItemsJustify(db *mongodb.MongoDB, quotation int64, items *ItemPurchase) error {
	collection := db.Client.Database(schema).Collection("purchases")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(items.ID.Hex())
	if err != nil {
		return fmt.Errorf("erro ao tentar converter o ID para ObjectID: %v", err)
	}

	items.Justificativa[0].ID = primitive.NewObjectID()
	filter := bson.M{"cotacao": quotation, "dadosconsolidados.id": id}
	update := bson.M{
		"$set": bson.M{
			"dadosconsolidados.$.precounitario": items.PrecoUnitario,
		},
		"$push": bson.M{
			"dadosconsolidados.$.justificativa": items.Justificativa[0],
		},
	}

	opts := options.Update().SetUpsert(true)
	if _, err := collection.UpdateOne(ctx, filter, update, opts); err != nil {
		return fmt.Errorf("erro ao tentar atualizar o documento: %v", err)
	}

	return nil

}

func DeletePurchaseItems(db *mongodb.MongoDB, quotation int64, items *ItemPurchase) error {
	collection := db.Client.Database(schema).Collection("purchases")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(items.ID.Hex())
	if err != nil {
		return fmt.Errorf("erro ao tentar converter o ID para ObjectID: %v", err)
	}

	now := time.Now().UTC()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.FixedZone("", -3*60*60))

	items.Justificativa[0].ID = primitive.NewObjectID()
	filter := bson.M{"cotacao": quotation, "dadosconsolidados.id": id}
	update := bson.M{
		"$set": bson.M{
			"dadosconsolidados.$.deleteat": today,
		},
		"$push": bson.M{
			"dadosconsolidados.$.justificativa": items.Justificativa[0],
		},
	}

	opts := options.Update().SetUpsert(true)
	if _, err := collection.UpdateOne(ctx, filter, update, opts); err != nil {
		return fmt.Errorf("erro ao tentar atualizar o documento: %v", err)
	}

	return nil

}
