package dadosabertos

import (
	"compras/internal/app/models"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type DadosAbertosComprasGov struct {
	client *http.Client
	url    string
}

func FnDadosAbertosComprasGov() *DadosAbertosComprasGov {
	return &DadosAbertosComprasGov{
		client: &http.Client{},
		url:    "https://dadosabertos.compras.gov.br/",
	}
}

func (d *DadosAbertosComprasGov) ConsultarMaterial(catmat string) ([]models.ItemPurchase, error) {

	page := 1

	if catmat == "" {
		return nil, errors.New("o catmat é um parâmetro obrigatório")
	}

	var data []models.ItemPurchase

	for {
		urlAPI := fmt.Sprintf("%smodulo-pesquisa-preco/1_consultarMaterial?pagina=%d&tamanhoPagina=500&codigoItemCatalogo=%s", d.url, page, catmat)

		resp, err := d.client.Get(urlAPI)
		if err != nil || resp.StatusCode != 200 {
			return nil, errors.New("API Dados abertos fora do ar")
		}
		defer resp.Body.Close()

		var result struct {
			Resultado        []models.ItemPurchase `json:"resultado"`
			PaginasRestantes int                   `json:"paginasRestantes"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			return nil, fmt.Errorf("erro ao decodificar a resposta JSON: %v", err)
		}

		data = append(data, result.Resultado...)

		remainingPages := result.PaginasRestantes
		if remainingPages > 0 {
			page++
		} else {
			break
		}
	}

	return data, nil
}
