import { defineStore } from 'pinia';
import ajaxItensCompras from "@/http/ajax-itens-compras/request.js";

import { Dialog, Notify } from 'quasar'
import { data } from 'autoprefixer';

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
    pagination: {
      rowsPerPage: 0
    },
    fixed: false,
    catmat: '',
    descricao: '',
    apresentacao: '',
    justificativas: []
  }),
  actions: {
    atualizarPrecoUnitario(cotacao, dados, valor, valorInicial){
      Dialog.create({
        title: 'Justificativa',
        message: 'Descreva o motivo da alteração? (Minimum 20 characters)',
        prompt: {
          model: '',
          isValid: val => val.length > 20,
          type: 'text'
        },
        cancel: true,
        persistent: true
      })
      .onCancel(() => {
        dados.precoUnitario = valorInicial;
        Notify.create({
          message: 'Operação cancelada!',
          color: 'negative',
          position: 'top-right',
          timeout: 3500,
          textColor: 'white'
        })
      })
      .onOk(justificativa => {
        dados.precoUnitario = valor;
        dados.justificativa = [{
          descricao: justificativa,
          data: new Date().toISOString().slice(0, 10),
          autor: 'Usuário',
          valor: valor,
          valorInicial: valorInicial
        }];
        ajaxItensCompras.atualizar(cotacao, dados).then((res)=>{
          if(res.data?.status_code == 200){
            Notify.create({
              message: `Preço unitário de ${valorInicial} para ${valor} atualizado com sucesso!`,
              color: 'positive',
              position: 'top-right',
              timeout: 3500,
              textColor: 'white'
            })

            setTimeout(() => {
              location.reload();
            }, 3000)
              
          }
        }).catch(error => {
          Notify.create({
            message: 'Erro ao atualizar preço unitário!',
            color: 'negative',
            position: 'top-right',
            timeout: 3500,
            textColor: 'white'
          })
        })
      })

    },
    exibirHistorico(item){
      if(item?.justificativa?.length == 0){
        Notify.create({
          message: 'Não há histórico de alterações!',
          color: 'negative',
          position: 'top-right',
          timeout: 3500,
          textColor: 'white'
        })
        return;
      }

      this.fixed = true;
      this.catmat = item?.codigoItemCatalogo;
      this.descricao = item?.descricaoItem;

      const nomeUnidadeFornecimento = item?.nomeUnidadeFornecimento || ""
      const capacidadeUnidadeFornecimento = item?.capacidadeUnidadeFornecimento || ""
      const siglaUnidadeMedida = item?.siglaUnidadeMedida || ""
      const apresentacao = `${nomeUnidadeFornecimento} ${capacidadeUnidadeFornecimento} ${siglaUnidadeMedida}`

      this.apresentacao = apresentacao
      this.justificativas = item?.justificativa.sort((a, b) => new Date(b.data) - new Date(a.data));
    },
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
    formatarMoeda(valor) {
      return valor.toLocaleString('pt-BR', {
        style: 'currency',
        currency: 'BRL',
      });
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
