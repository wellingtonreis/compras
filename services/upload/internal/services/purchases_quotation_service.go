package services

import (
	"compras/services/upload/internal/entity"
	"compras/services/upload/internal/platform/database/mongodb"
	"compras/services/upload/internal/repository"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const situation = "iniciada"

func CreatePurchasesQuotation() (int64, error) {

	db := mongodb.ConnectToMongoDB()
	defer db.Close()

	sequence, err := db.GetNextSequenceValue("sequence_purchases_quotation")
	if err != nil {
		log.Fatal("Erro ao tentar cadastrar a sequencia de identificação:", err)
	}

	data := &entity.QuotationHistory{
		ID:          primitive.NewObjectID(),
		Cotacao:     sequence,
		Hu:          "",
		Situacao:    situation,
		Processosei: "",
		Autor:       "",
	}

	repository := &repository.MongoDB{
		Client: db.Client,
	}
	err = repository.Save(data)
	if err != nil {
		log.Fatal("Erro ao tentar salvar a cotação:", err)
		return 0, err
	}
	return sequence, nil
}
