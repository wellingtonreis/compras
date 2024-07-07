import { defineStore } from 'pinia';
import ajaxConsultarCotacao from "@/http/ajax-consultar-cotacao/request.js";
import ajaxCategoriaEhSubcategoria from "@/http/ajax-categoria-subcategoria/request.js";

import { Notify } from 'quasar'

export const useConsultarCotacaoStore = defineStore('useConsultarCotacaoStore', {
  state: () => ({
    filtro: {
      cotacao: null,
      data_inicio: "",
      data_fim: "",
      categoria: "",
      subcategoria: "",
      situacao: "",
      processosei: "",
      autor: "",
    },
    date_between: { from: "", to: "" },
    pesquisa_avancada: false,
    categorias: [],
    opcoesCategorias: [],
    subcategorias: [],
    opcoesSubCategorias: [],
    opcoesSituacao: [
      "Iniciada",
      "Calculada",
      "Finalizada"
    ],
    columns: [
      {
        name: "acao",
        required: true,
        label: "AÇÕES",
        align: "left",
        field: (row) => row.acao,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "cotacao",
        required: true,
        label: "ID",
        align: "left",
        field: (row) => row.cotacao,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "hu",
        required: true,
        label: "HU",
        align: "left",
        field: (row) => row.hu,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "categoria",
        required: true,
        label: "CATEGORIA",
        align: "left",
        field: (row) => row.categoria,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "subcategoria",
        required: true,
        label: "SUBCATEGORIA",
        align: "left",
        field: (row) => row.subcategoria,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "data",
        required: true,
        label: "DATA",
        align: "left",
        field: (row) => {
          const date_custom = row.datahora.split("T")
          return date_custom[0].split("-").reverse().join("/")
        },
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "situacao",
        required: true,
        label: "SITUAÇÃO",
        align: "left",
        field: (row) => row.situacao,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "processo_sei",
        required: true,
        label: "PROCESSO SEI",
        align: "left",
        field: (row) => row.processosei,
        format: (val) => `${val}`,
        sortable: true,
      },
      {
        name: "autor",
        required: true,
        label: "AUTOR",
        align: "left",
        field: (row) => row.autor,
        format: (val) => `${val}`,
        sortable: true,
      },
    ],
    rows: []
  }),
  actions: {
    reiniciar() {
      this.filtro = {
        cotacao: null,
        datainicio: "",
        datafim: "",
        categoria: null,
        subcategoria: "",
        situacao: "",
        processo_sei: "",
      };
    },
    enviar() {
      this.filtro.cotacao = parseInt(this.filtro.cotacao);
      this.filtro.data_inicio = this.filtro.data_inicio.split("/").reverse().join("-");
      this.filtro.data_fim = this.filtro.data_fim.split("/").reverse().join("-");

      this.listaHitoricoCotacao();
    },
    listaHitoricoCotacao(){
      ajaxConsultarCotacao.listar(this.filtro).then((res)=>{
        if(res.data?.status_code == 200){
            this.rows = res.data?.embedded
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
      this.filtro.subcategoria = "";
      this.opcoesSubCategorias = this.subcategorias[val].map((subcategoria) => { 
        return { "value": val, "label": subcategoria.name };
      })
    },
  }
});
