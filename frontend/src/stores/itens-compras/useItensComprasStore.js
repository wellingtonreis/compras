import { defineStore } from 'pinia';
import ajaxItensCompras from "@/http/ajax-itens-compras/request.js";
import ajaxCategoriaEhSubcategoria from "@/http/ajax-categoria-subcategoria/request.js";
import ajaxClassificacaoSegmento from "@/http/ajax-historico-cotacao/request.js";

import { Dialog, Notify } from 'quasar'

export const useItensComprasStore = defineStore('useItensComprasStore', {
  state: () => ({
    cotacao: 0,
    autor: "Anônimo",
    categorias: [],
    opcoesCategorias: [],
    subcategorias: [],
    opcoesSubCategorias: [],
    categoria: "",
    subcategoria: "",
    processosei: "",
    urlprotocolosei: "",
    validaCategoria: true,
    validaSubcategoria: true,
    validaProcessosei: true,
    columns: [
      {
        name: "catmat",
        required: true,
        label: "CATMAT",
        align: "center",
        field: (row) => row.codigoItemCatalogo,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "descricao",
        required: true,
        label: "DESCRIÇÃO",
        align: "center",
        field: (row) => row.descricaoItem,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "apresentacao",
        required: true,
        label: "APRESENTAÇÃO",
        align: "center",
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
        align: "center",
        field: (row) => row,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "dataCompra",
        required: true,
        label: "DATA COMPRA",
        align: "center",
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
        align: "center",
        field: (row) => row.quantidade,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "precoUnitario",
        required: true,
        label: "PREÇO UNITÁRIO",
        align: "center",
        sortable: true,
      },
      {
        name: "justificativa",
        required: false,
        label: "JUSTIFICATIVA",
        align: "center",
        sortable: false,
      },
      {
        name: "remover",
        required: false,
        label: "REMOVER",
        align: "center",
        sortable: false,
      }
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
    removerItemDeCompra (index, cotacao, dados) {
      const vm = this;
      Dialog.create({
        title: 'Justificativa',
        message: 'Descreva o motivo da exclusão deste item? (Minimum 20 characters)',
        prompt: {
          model: '',
          isValid: val => val.length > 20,
          type: 'text'
        },
        cancel: true,
        persistent: true
      })
      .onCancel(() => {
        Notify.create({
          message: 'Operação cancelada!',
          color: 'negative',
          position: 'top-right',
          timeout: 3500,
          textColor: 'white'
        })
      })
      .onOk(justificativa => {
        
        dados.deleteat = new Date().toISOString();
        dados.justificativa = [{
          descricao: justificativa,
          data: new Date().toISOString().slice(0, 10),
          autor: vm.autor,
          valor: null,
          valorInicial: null,
          deleteat: new Date().toISOString(),
        }];

        ajaxItensCompras.remover(cotacao, dados).then((res)=>{
          if(res.data?.status_code == 200){
            Notify.create({
              message: `Item excluído!`,
              color: 'positive',
              position: 'top-right',
              timeout: 3500,
              textColor: 'white'
            })
            vm.rows = [ ...vm.rows.slice(0, index), ...vm.rows.slice(index + 1) ]
            setTimeout(() => {
              location.reload();
            }, 3000)
          }
        }).catch(error => {
          console.log(error);
          Notify.create({
            message: 'Erro ao tentar excluir o item!',
            color: 'negative',
            position: 'top-right',
            timeout: 3500,
            textColor: 'white'
          })
        })
      })
    },
    atualizarPrecoUnitario(cotacao, dados, valor, valorInicial){
      const vm = this;
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
        dados.deleteat = null;
        dados.precoUnitario = valor;
        dados.justificativa = [{
          descricao: justificativa,
          data: new Date().toISOString().slice(0, 10),
          autor: vm.autor,
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
    listaSuspensaCategoriaEhSubcategoria(){
      ajaxCategoriaEhSubcategoria.listar().then((res)=>{
        if(res.data?.status_code == 200){
          this.categorias = res.data?.embedded.map((categoria) => {
            this.subcategorias[categoria._id] = categoria.subcategories;
            return { "value": categoria._id, "label": categoria.name };
          })
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
    filtraCategoria (val, update) {
      this.opcoesCategorias = this.categorias

      update(() => {
        const needle = val.toLowerCase()
        this.opcoesCategorias = this.categorias.filter(v => v.label.toLowerCase().indexOf(needle) > -1)
      })
    },
    filtraSubcategoria (val, update) {
      update(() => {
        const needle = val.toLowerCase()
        this.opcoesSubCategorias = this.opcoesSubCategorias.filter(v => v.label.toLowerCase().indexOf(needle) > -1)
      })
    },
    selecionaSubcategoria (val) {
      this.subcategoria = "";
      this.opcoesSubCategorias = this.subcategorias[val].map((subcategoria) => { 
        return { "value": val, "label": subcategoria.name };
      })
    },
    reiniciar(){
      this.categoria = "";
      this.subcategoria = "";
      this.processosei = "";
      this.urlprotocolosei = "";
    },
    salvarClassificacaoSegmento(step, cotacao){

      const historicoCotacao = {
        "cotacao": cotacao,
        "situacao": 1,
        "categoria": this.categoria.value,
        "subcategoria": this.subcategoria.value,
        "processosei": this.processosei,
      }
      ajaxClassificacaoSegmento.atualizar(historicoCotacao, cotacao).then((res)=>{
        if(res.data?.status_code == 200){
          Notify.create({
            message: `Classificação salva com sucesso!`,
            color: 'positive',
            position: 'top-right',
            timeout: 3500,
            textColor: 'white'
          })
          this.reiniciar()
          step()
        }
      }
      ).catch(error => {
        Notify.create({
          message: 'Erro ao tentar salvar classificação!',
          color: 'negative',
          position: 'top-right',
          timeout: 3500,
          textColor: 'white'
        })
      })
    },
    preencheProcessoSei(){
      this.urlprotocolosei = this.processosei.replace(/[^0-9]/g, '');
    }
  }
});
