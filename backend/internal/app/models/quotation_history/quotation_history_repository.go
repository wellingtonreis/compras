package models

import (
	"context"
	"fmt"
	"time"

	"github.com/wellingtonreis/compras/pkg/date_custom"

	"github.com/wellingtonreis/compras/internal/app/platform/database/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var schema = mongodb.GetMongoSchema()

func UpdateQuotationHistorySegment(db *mongodb.MongoDB, quotation int64, items *QuotationHistory) error {
	collection := db.Client.Database(schema).Collection("purchases")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"cotacao": quotation}
	update := bson.M{
		"$set": bson.M{
			"categoria":    items.Categoria,
			"subcategoria": items.Subcategoria,
			"processosei":  items.Processosei,
		},
	}

	opts := options.Update().SetUpsert(true)
	if _, err := collection.UpdateOne(ctx, filter, update, opts); err != nil {
		return fmt.Errorf("erro ao tentar atualizar o documento: %v", err)
	}

	return nil
}

func SearchQuotationHistory(db *mongodb.MongoDB, filter *FilterQuotationHistory) ([]QuotationHistory, error) {
	collection := db.Client.Database(schema).Collection("purchases")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pipeline := mongo.Pipeline{}

	if filter.Cotacao > 0 {
		pipeline = append(pipeline, bson.D{{"$match", bson.D{{"cotacao", filter.Cotacao}}}})
	}

	if filter.Hu != "" {
		pipeline = append(pipeline, bson.D{{"$match", bson.D{{"hu", filter.Hu}}}})
	}

	if filter.Categoria != primitive.NilObjectID {
		pipeline = append(pipeline, bson.D{{"$match", bson.D{{"categoria", filter.Categoria}}}})
	}

	if filter.Subcategoria != primitive.NilObjectID {
		pipeline = append(pipeline, bson.D{{"$match", bson.D{{"subcategoria", filter.Subcategoria}}}})
	}

	if filter.Situacao != "" {
		pipeline = append(pipeline, bson.D{{"$match", bson.D{{"situacao", filter.Situacao}}}})
	}

	if filter.Autor != "" {
		pipeline = append(pipeline, bson.D{{"$match", bson.D{{"autor", filter.Autor}}}})
	}

	if filter.Processosei != "" {
		pipeline = append(pipeline, bson.D{{"$match", bson.D{{"processosei", filter.Processosei}}}})
	}

	if filter.DataInicio != "" && filter.DataFim != "" {
		date_start := date_custom.ConvertDateStringToISO8601(filter.DataInicio)
		date_end := date_custom.ConvertDateStringToISO8601(filter.DataFim)
		pipeline = append(pipeline, bson.D{{"$match", bson.D{
			{"datahora", bson.D{
				{"$gte", date_start},
				{"$lte", date_end},
			}},
		}}})
	}

	pipeline = append(pipeline, bson.D{
		{"$group", bson.D{
			{"_id", "$cotacao"},
			{"cotacao", bson.D{{"$first", "$cotacao"}}},
			{"datahora", bson.D{{"$first", "$datahora"}}},
			{"situacao", bson.D{{"$first", "$situacao"}}},
			{"processosei", bson.D{{"$first", "$processosei"}}},
			{"autor", bson.D{{"$first", "$autor"}}},
		}},
	})

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar a consulta de agregação: %v", err)
	}
	defer cursor.Close(ctx)

	var docs []QuotationHistory
	if err = cursor.All(ctx, &docs); err != nil {
		return nil, fmt.Errorf("erro ao tentar ler todos os documentos de histórico de cotação: %v", err)
	}

	return docs, nil
}
