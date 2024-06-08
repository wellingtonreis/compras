package models

import (
	"compras/internal/app/platform/database/mongodb"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Data struct {
	Catalog []CatalogCode `json:"catalog"`
}

type CatalogCode struct {
	Catmat            string         `json:"catmat"`
	Apresentacao      string         `json:"apresentacao"`
	Quantidade        string         `json:"quantidade"`
	Cotacao           string         `json:"cotacao"`
	DadosAPI          []ItemPurchase `json:"dadosapi"`
	DadosConsolidados []ItemPurchase `json:"dadosconsolidados"`
}

type ItemPurchase struct {
	IDCompra                      string    `json:"idCompra"`
	IDItemCompra                  int       `json:"idItemCompra"`
	Forma                         string    `json:"forma"`
	Modalidade                    int       `json:"modalidade"`
	CriterioJulgamento            string    `json:"criterioJulgamento"`
	NumeroItemCompra              int       `json:"numeroItemCompra"`
	DescricaoItem                 string    `json:"descricaoItem"`
	CodigoItemCatalogo            int       `json:"codigoItemCatalogo"`
	NomeUnidadeMedida             string    `json:"nomeUnidadeMedida"`
	SiglaUnidadeMedida            string    `json:"siglaUnidadeMedida"`
	NomeUnidadeFornecimento       string    `json:"nomeUnidadeFornecimento"`
	SiglaUnidadeFornecimento      string    `json:"siglaUnidadeFornecimento"`
	CapacidadeUnidadeFornecimento float64   `json:"capacidadeUnidadeFornecimento"`
	Quantidade                    float64   `json:"quantidade"`
	PrecoUnitario                 float64   `json:"precoUnitario"`
	PercentualMaiorDesconto       float64   `json:"percentualMaiorDesconto"`
	NIFornecedor                  string    `json:"niFornecedor"`
	NomeFornecedor                string    `json:"nomeFornecedor"`
	Marca                         string    `json:"marca"`
	CodigoUasg                    string    `json:"codigoUasg"`
	NomeUasg                      string    `json:"nomeUasg"`
	CodigoMunicipio               int       `json:"codigoMunicipio"`
	Municipio                     string    `json:"municipio"`
	Estado                        string    `json:"estado"`
	CodigoOrgao                   int       `json:"codigoOrgao"`
	NomeOrgao                     string    `json:"nomeOrgao"`
	Poder                         string    `json:"poder"`
	Esfera                        string    `json:"esfera"`
	DataCompra                    time.Time `json:"dataCompra"`
	DataHoraAtualizacaoCompra     string    `json:"dataHoraAtualizacaoCompra"`
	DataHoraAtualizacaoItem       string    `json:"dataHoraAtualizacaoItem"`
	DataResultado                 string    `json:"dataResultado"`
	DataHoraAtualizacaoUasg       string    `json:"dataHoraAtualizacaoUasg"`
}

func FilterDocumentsByPresentation(db *mongodb.MongoDB) ([]CatalogCode, error) {
	collection := db.Client.Database("admin").Collection("purchases")
	ctx := context.Background()

	now := time.Now().UTC()
	year, month, day := now.Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, time.FixedZone("", -3*60*60))

	threeMonthsAgo := today.AddDate(0, -3, 0)
	sixMonthsAgo := today.AddDate(0, -6, 0)
	twelveMonthsAgo := today.AddDate(0, -12, 0)

	pipeline := mongo.Pipeline{
		{{"$match", bson.M{"apresentacao": bson.M{"$exists": true}}}},
		{{"$addFields", bson.D{
			{"dadosconsolidados", bson.M{"$filter": bson.M{
				"input": "$dadosapi",
				"as":    "item",
				"cond":  bson.M{"$eq": []interface{}{"$$item.nomeunidadefornecimento", "$apresentacao"}},
			}}},
		}}},
		{{"$match", bson.M{"dadosconsolidados": bson.M{"$ne": []interface{}{}}}}},
		{{"$addFields", bson.D{
			{"dadosconsolidados", bson.M{"$filter": bson.M{
				"input": "$dadosconsolidados",
				"as":    "item",
				"cond": bson.M{
					"$or": []interface{}{
						bson.M{
							"$and": []interface{}{
								bson.M{"$gte": []interface{}{"$$item.datacompra", threeMonthsAgo}},
								bson.M{"$eq": []interface{}{"$buscarSeVazio", true}},
							},
						},
						bson.M{
							"$and": []interface{}{
								bson.M{"$gte": []interface{}{"$$item.datacompra", sixMonthsAgo}},
								bson.M{"$eq": []interface{}{"$buscarSeVazio", true}},
							},
						},
						bson.M{"$gte": []interface{}{"$$item.datacompra", twelveMonthsAgo}},
					},
				},
			}}},
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, fmt.Errorf("erro ao executar a consulta de agregação: %v", err)
	}
	defer cursor.Close(ctx)

	var docs []CatalogCode
	if err = cursor.All(ctx, &docs); err != nil {
		return nil, fmt.Errorf("erro ao tentar ler todos os documentos: %v", err)
	}

	var updatedDocuments []CatalogCode

	for _, doc := range docs {
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

	// for _, doc := range docs {

	// 	if _, err := collection.DeleteOne(ctx, bson.M{"catmat": doc.Catmat, "cotacao": doc.Cotacao}); err != nil {
	// 		return nil, fmt.Errorf("erro ao tentar deletar o documento: %v", err)
	// 	}

	// 	doc.DadosAPI = nil
	// 	if _, err := collection.InsertOne(ctx, doc); err != nil {
	// 		return nil, fmt.Errorf("erro ao tentar inserir o documento: %v", err)
	// 	}

	// 	updatedDocuments = append(updatedDocuments, doc)
	// }

	return updatedDocuments, nil
}
