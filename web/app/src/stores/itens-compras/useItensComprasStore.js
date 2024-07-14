import { defineStore } from 'pinia';
import ajaxItensCompras from "@/http/ajax-itens-compras/request.js";

import { Notify } from 'quasar'

export const useItensComprasStore = defineStore('useItensComprasStore', {
  state: () => ({
    columns: [
      {
        name: "catmat",
        required: true,
        label: "CATMAT",
        align: "left",
        field: (row) => row.codigoItemCatalogo,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "descricao",
        required: true,
        label: "DESCRIÇÃO",
        align: "left",
        field: (row) => row.descricaoItem,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "apresentacao",
        required: true,
        label: "APRESENTAÇÃO",
        align: "left",
        field: (row) => {
          const nomeUnidadeFornecimento = row.nomeUnidadeFornecimento || ""
          const capacidadeUnidadeFornecimento = row.capacidadeUnidadeFornecimento || ""
          const siglaUnidadeMedida = row.siglaUnidadeMedida || ""
          return `${nomeUnidadeFornecimento} ${capacidadeUnidadeFornecimento} ${siglaUnidadeMedida}`
        },
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "idCompra",
        required: true,
        label: "ID COMPRA",
        align: "left",
        field: (row) => row,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "dataCompra",
        required: true,
        label: "DATA COMPRA",
        align: "left",
        field: (row) => row.dataCompra,
        format: (val) => {
          const date = new Date(val);
          const month = date.toLocaleString('default', { month: 'long' });
          const year = date.getFullYear();
            return `${month.charAt(0).toUpperCase() + month.slice(1)} ${year}`;
        },
        sortable: true,
      },
      {
        name: "quantidade",
        required: true,
        label: "QUANTIDADE COMPRADA DE ITEM",
        align: "left",
        field: (row) => row.quantidade,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "precoUnitario",
        required: true,
        label: "PREÇO UNITÁRIO",
        align: "left",
        sortable: true,
      },
      {
        name: "justificativa",
        required: true,
        label: "JUSTIFICATIVA",
        align: "left",
        sortable: true,
      },
    ],
    rows: [],
    erroPrecoUnitario: false,
    erroMensagemPrecoUnitario: "",
  }),
  actions: {
    validarPrecoUnitario (val) {
      const regexDecimal = /^\d+(\.\d+)?$/;
      if (!regexDecimal.test(val)) {
        this.erroPrecoUnitario = true;
        this.erroMensagemPrecoUnitario = 'O valor não é um decimal válido. Exemplo: 10.00 ou 10.';
        return false;
      }
      const regexVirgula = /^\d+(\,\d+)?$/;
      if (regexVirgula.test(val)) {
        this.erroPrecoUnitario = true;
        this.erroMensagemPrecoUnitario = 'Utilize ponto como separador decimal.';
        return false;
      }
      this.erroPrecoUnitario = false
      this.erroMensagemPrecoUnitario = ''
      return true
    },
    formatIdCompra(val) {
      const numeroItemCompra = val?.numeroItemCompra.toString() || ""
      const idCompra = val?.idCompra || ""

      const numeroIdentificadoItemCompra = idCompra + "" + numeroItemCompra.padStart(5, '0')
      if(numeroIdentificadoItemCompra?.length > 0){
        const cincoUltimosDigitos = numeroIdentificadoItemCompra.slice(-5);
        const digitosRestantes = numeroIdentificadoItemCompra.slice(0, -5);
        const dezesseteDigitos = digitosRestantes.padStart(17, '0');

        const url = `https://cnetmobile.estaleiro.serpro.gov.br/comprasnet-web/public/compras/acompanhamento-compra/item/${cincoUltimosDigitos}?compra=${dezesseteDigitos}`;
        
        return `<a href="${url}" target="_blank">${numeroIdentificadoItemCompra}</a>`;
      } else {
        return 'Outros. Verificar/anexar no processo';
      }
    },
    listaItensCompras(cotacao){
      ajaxItensCompras.listar(cotacao).then((res)=>{
        if(res.data?.status_code == 200){
            this.rows = res.data?.embedded ?? [];
            this.reiniciar()
        }
      }).catch(error => {
        Notify.setDefaults({
          position: 'top-right',
          timeout: 2500,
          textColor: 'white',
          actions: [{ icon: 'close', color: 'white' }]
        })
      })
    },
  }
});
