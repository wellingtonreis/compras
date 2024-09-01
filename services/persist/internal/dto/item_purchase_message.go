package dto

import "compras/services/persist/internal/entity"

type ItemPurchaseMessage struct {
	Items     []entity.ItemPurchase `json:"items"`
	Quotation int64                 `json:"quotation"`
}
