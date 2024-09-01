package services

import (
	"compras/services/upload/internal/dto"
	"compras/services/upload/internal/entity"
	"compras/services/upload/internal/platform/database/mongodb"
	"compras/services/upload/internal/repository"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SequenceQuotation() (int64, error) {

	db := mongodb.ConnectToMongoDB()
	defer db.Close()

	sequence, err := db.GetNextSequenceValue("sequence_purchases_quotation")
	if err != nil {
		log.Fatal("Erro ao tentar cadastrar a sequencia de identificação:", err)
	}

	return sequence, nil
}

func CreatePurchasesQuotation(catmat *dto.Catmat) error {

	db := mongodb.ConnectToMongoDB()
	defer db.Close()

	quotation := entity.QuotationHistory{
		ID:           primitive.NewObjectID(),
		Catmat:       catmat.Catmat,
		Apresentacao: catmat.Apresentacao,
		Quantidade:   catmat.Quantidade,
		Cotacao:      catmat.Quotation,
		Situacao:     "iniciada",
	}

	repository := &repository.MongoDB{
		Client: db.Client,
	}

	err := repository.Save(&quotation)
	if err != nil {
		log.Fatal("Erro ao tentar salvar a cotação:", err)
		return err
	}

	return nil
}
