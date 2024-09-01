package services

import (
	"compras/services/persist/internal/dto"
	"compras/services/persist/internal/entity"
	"compras/services/persist/internal/platform/database/mongodb"
	"compras/services/persist/internal/repository"
	"log"
)

func Save(ItemPurchaseMessage *dto.ItemPurchaseMessage) ([]entity.QuotationHistory, error) {
	db := mongodb.ConnectToMongoDB()
	defer db.Close()
	quotationHistory, err := repository.SavePurchaseItemDocuments(db, ItemPurchaseMessage)
	if err != nil {
		log.Fatalf("Erro ao tentar cadastrar os documentos: %v", err)
		return nil, err
	}

	return quotationHistory, nil
}
