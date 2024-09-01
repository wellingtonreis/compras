package services

import (
	"compras/services/govbr/internal/entity"
	dadosabertos "compras/services/govbr/pkg/dados_abertos"
	"fmt"
	"time"
)

type SearchDataCatmat struct {
	Catmat string
	Retry  int
}

func (cm *SearchDataCatmat) Search() ([]entity.ItemPurchase, error) {
	var result []entity.ItemPurchase
	var err error

	api := dadosabertos.FnDadosAbertosComprasGov()
	for i := 1; i <= cm.Retry; i++ {
		result, err = api.ConsultarMaterial(cm.Catmat)
		if err == nil {
			return result, nil
		}
		time.Sleep(5 * time.Second)
	}
	return nil, fmt.Errorf("erro ao tentar consultar os itens de compras")
}

type ChannelDataItemPurchase struct {
	Channel chan ItemPurchaseMessage
	Result  []entity.ItemPurchase
}

type ItemPurchaseMessage struct {
	Items     []entity.ItemPurchase `json:"items"`
	Quotation int64                 `json:"quotation"`
}

func (c *ChannelDataItemPurchase) ChannelDataItemPurchase(sequence int64) {

	items := make([]entity.ItemPurchase, 0)
	for _, item := range c.Result {
		itemPopulated := entity.ItemPurchase{
			IDCompra:                      item.IDCompra,
			IDItemCompra:                  item.IDItemCompra,
			Forma:                         item.Forma,
			Modalidade:                    item.Modalidade,
			CriterioJulgamento:            item.CriterioJulgamento,
			NumeroItemCompra:              item.NumeroItemCompra,
			DescricaoItem:                 item.DescricaoItem,
			CodigoItemCatalogo:            item.CodigoItemCatalogo,
			NomeUnidadeMedida:             item.NomeUnidadeMedida,
			SiglaUnidadeMedida:            item.SiglaUnidadeMedida,
			NomeUnidadeFornecimento:       item.NomeUnidadeFornecimento,
			SiglaUnidadeFornecimento:      item.SiglaUnidadeFornecimento,
			CapacidadeUnidadeFornecimento: item.CapacidadeUnidadeFornecimento,
			Quantidade:                    item.Quantidade,
			PrecoUnitario:                 item.PrecoUnitario,
			PercentualMaiorDesconto:       item.PercentualMaiorDesconto,
			NIFornecedor:                  item.NIFornecedor,
			NomeFornecedor:                item.NomeFornecedor,
			Marca:                         item.Marca,
			CodigoUasg:                    item.CodigoUasg,
			NomeUasg:                      item.NomeUasg,
			CodigoMunicipio:               item.CodigoMunicipio,
			Municipio:                     item.Municipio,
			Estado:                        item.Estado,
			CodigoOrgao:                   item.CodigoOrgao,
			NomeOrgao:                     item.NomeOrgao,
			Poder:                         item.Poder,
			Esfera:                        item.Esfera,
			DataCompra:                    item.DataCompra,
			DataHoraAtualizacaoCompra:     item.DataHoraAtualizacaoCompra,
			DataHoraAtualizacaoItem:       item.DataHoraAtualizacaoItem,
			DataResultado:                 item.DataResultado,
			DataHoraAtualizacaoUasg:       item.DataHoraAtualizacaoUasg,
		}

		items = append(items, itemPopulated)
	}

	itemPurchaseMessage := ItemPurchaseMessage{
		Items:     items,
		Quotation: sequence,
	}
	c.Channel <- itemPurchaseMessage
}
