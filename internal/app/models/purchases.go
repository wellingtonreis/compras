package models

import (
	"compras/internal/app/platform/database/mongodb"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Data struct {
	Catalog []CatalogCode `json:"catalog"`
}

type CatalogCode struct {
	Catmat            string         `json:"catmat"`
	Apresentacao      string         `json:"apresentacao"`
	Quantidade        string         `json:"quantidade"`
	Resultado         []ItemPurchase `json:"resultado"`
	ResultadoFiltrado []ItemPurchase `json:"resultadoFiltrado"`
}

type ItemPurchase struct {
	IDCompra                      string  `json:"idCompra"`
	IDItemCompra                  int     `json:"idItemCompra"`
	Forma                         string  `json:"forma"`
	Modalidade                    int     `json:"modalidade"`
	CriterioJulgamento            string  `json:"criterioJulgamento"`
	NumeroItemCompra              int     `json:"numeroItemCompra"`
	DescricaoItem                 string  `json:"descricaoItem"`
	CodigoItemCatalogo            int     `json:"codigoItemCatalogo"`
	NomeUnidadeMedida             string  `json:"nomeUnidadeMedida"`
	SiglaUnidadeMedida            string  `json:"siglaUnidadeMedida"`
	NomeUnidadeFornecimento       string  `json:"nomeUnidadeFornecimento"`
	SiglaUnidadeFornecimento      string  `json:"siglaUnidadeFornecimento"`
	CapacidadeUnidadeFornecimento float64 `json:"capacidadeUnidadeFornecimento"`
	Quantidade                    float64 `json:"quantidade"`
	PrecoUnitario                 float64 `json:"precoUnitario"`
	PercentualMaiorDesconto       float64 `json:"percentualMaiorDesconto"`
	NIFornecedor                  string  `json:"niFornecedor"`
	NomeFornecedor                string  `json:"nomeFornecedor"`
	Marca                         string  `json:"marca"`
	CodigoUasg                    string  `json:"codigoUasg"`
	NomeUasg                      string  `json:"nomeUasg"`
	CodigoMunicipio               int     `json:"codigoMunicipio"`
	Municipio                     string  `json:"municipio"`
	Estado                        string  `json:"estado"`
	CodigoOrgao                   int     `json:"codigoOrgao"`
	NomeOrgao                     string  `json:"nomeOrgao"`
	Poder                         string  `json:"poder"`
	Esfera                        string  `json:"esfera"`
	DataCompra                    string  `json:"dataCompra"`
	DataHoraAtualizacaoCompra     string  `json:"dataHoraAtualizacaoCompra"`
	DataHoraAtualizacaoItem       string  `json:"dataHoraAtualizacaoItem"`
	DataResultado                 string  `json:"dataResultado"`
	DataHoraAtualizacaoUasg       string  `json:"dataHoraAtualizacaoUasg"`
}

func FilterDocumentsByPresentation(db *mongodb.MongoDB) ([]CatalogCode, error) {
	collection := db.Client.Database("admin").Collection("purchases")
	ctx := context.Background()

	currentTime := time.Now()

	threeMonthsAgo := currentTime.AddDate(0, -3, 0)
	sixMonthsAgo := currentTime.AddDate(0, -6, 0)
	twelveMonthsAgo := currentTime.AddDate(0, -12, 0)

	pipeline := mongo.Pipeline{
		{{"$match", bson.M{"apresentacao": bson.M{"$exists": true}}}},
		{{"$addFields", bson.M{
			"resultadoFiltrado": bson.M{"$filter": bson.M{
				"input": "$resultado",
				"as":    "item",
				"cond":  bson.M{"$eq": []interface{}{"$$item.nomeunidadefornecimento", "$apresentacao"}},
			}},
			"datacompra": bson.M{"$dateToString": bson.M{
				"format": "%Y-%m-%d",
				"date":   "$datacompra",
			}},
		}}},
		{{"$match", bson.M{"resultadoFiltrado": bson.M{"$ne": []interface{}{}}}}},
		{{"$match", bson.M{
			"datacompra": bson.M{"$gte": threeMonthsAgo.Format("2006-01-02")},
		}}},
		{{"$group", bson.M{
			"_id":   nil,
			"count": bson.M{"$sum": 1},
		}}},
		{{"$addFields", bson.M{
			"buscarSeVazio": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$eq": []interface{}{"$count", 0}},
					"then": true,
					"else": false,
				},
			},
		}}},
		{{"$match", bson.M{
			"buscarSeVazio": true,
			"datacompra":    bson.M{"$gte": sixMonthsAgo.Format("2006-01-02")},
		}}},
		{{"$group", bson.M{
			"_id":   nil,
			"count": bson.M{"$sum": 1},
		}}},
		{{"$addFields", bson.M{
			"buscarSeVazio": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$eq": []interface{}{"$count", 0}},
					"then": true,
					"else": false,
				},
			},
		}}},
		{{"$match", bson.M{
			"buscarSeVazio": true,
			"datacompra":    bson.M{"$gte": twelveMonthsAgo.Format("2006-01-02")},
		}}},
	}

	cur, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("error executing aggregate query: %v", err)
	}
	defer cur.Close(ctx)

	var updatedDocuments []CatalogCode

	var documents []CatalogCode
	if err := cur.All(ctx, &documents); err != nil {
		return nil, fmt.Errorf("error retrieving documents: %v", err)
	}

	for _, doc := range documents {

		if _, err := collection.DeleteOne(ctx, bson.M{"catmat": doc.Catmat}); err != nil {
			return nil, fmt.Errorf("error deleting document: %v", err)
		}

		doc.Resultado = doc.ResultadoFiltrado

		if _, err := collection.InsertOne(ctx, doc); err != nil {
			return nil, fmt.Errorf("error inserting document: %v", err)
		}

		updatedDocuments = append(updatedDocuments, doc)
	}

	if err := cur.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %v", err)
	}

	return updatedDocuments, nil
}
