package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ItemPurchase struct {
	ID                            primitive.ObjectID `json:"_id,omitempty"`
	IDCompra                      string             `json:"idCompra"`
	IDItemCompra                  int                `json:"idItemCompra"`
	Forma                         string             `json:"forma"`
	Modalidade                    int                `json:"modalidade"`
	CriterioJulgamento            string             `json:"criterioJulgamento"`
	NumeroItemCompra              int                `json:"numeroItemCompra"`
	DescricaoItem                 string             `json:"descricaoItem"`
	CodigoItemCatalogo            int                `json:"codigoItemCatalogo"`
	NomeUnidadeMedida             string             `json:"nomeUnidadeMedida"`
	SiglaUnidadeMedida            string             `json:"siglaUnidadeMedida"`
	NomeUnidadeFornecimento       string             `json:"nomeUnidadeFornecimento"`
	SiglaUnidadeFornecimento      string             `json:"siglaUnidadeFornecimento"`
	CapacidadeUnidadeFornecimento float64            `json:"capacidadeUnidadeFornecimento"`
	Quantidade                    float64            `json:"quantidade"`
	PrecoUnitario                 float64            `json:"precoUnitario"`
	PercentualMaiorDesconto       float64            `json:"percentualMaiorDesconto"`
	NIFornecedor                  string             `json:"niFornecedor"`
	NomeFornecedor                string             `json:"nomeFornecedor"`
	Marca                         string             `json:"marca"`
	CodigoUasg                    string             `json:"codigoUasg"`
	NomeUasg                      string             `json:"nomeUasg"`
	CodigoMunicipio               int                `json:"codigoMunicipio"`
	Municipio                     string             `json:"municipio"`
	Estado                        string             `json:"estado"`
	CodigoOrgao                   int                `json:"codigoOrgao"`
	NomeOrgao                     string             `json:"nomeOrgao"`
	Poder                         string             `json:"poder"`
	Esfera                        string             `json:"esfera"`
	DataCompra                    time.Time          `json:"dataCompra"`
	DataHoraAtualizacaoCompra     string             `json:"dataHoraAtualizacaoCompra"`
	DataHoraAtualizacaoItem       string             `json:"dataHoraAtualizacaoItem"`
	DataResultado                 string             `json:"dataResultado"`
	DataHoraAtualizacaoUasg       string             `json:"dataHoraAtualizacaoUasg"`
	Justificativa                 []Justification    `json:"justificativa,omitempty"`
	DeleteAt                      *time.Time         `json:"deleteAt,omitempty"`
}
