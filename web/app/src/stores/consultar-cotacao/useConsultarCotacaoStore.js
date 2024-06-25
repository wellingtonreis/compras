import { defineStore } from 'pinia';
import request from "@/http/consultar-cotacao/request.js";

import { Notify } from 'quasar'

export const useConsultarCotacaoStore = defineStore('useConsultarCotacaoStore', {
  state: () => ({
    filtro: {
      cotacao: 0
    },
    date_between: { from: "", to: "" },
    pesquisa_avancada: false,
    opcoesCategorias: [
      "Opção 1",
      "Opção 2",
      "Opção 3",
      "Opção 4",
      "Opção 5",
    ],
    opcoesSubCategorias: [
      "Opção 1",
      "Opção 2",
      "Opção 3",
      "Opção 4",
      "Opção 5",
    ],
    opcoesSituacao: [
      "Opção 1",
      "Opção 2",
      "Opção 3",
      "Opção 4",
      "Opção 5",
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
        field: (row) => row.datahora,
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
    submitForm() {
      console.log("Filtro:", this.filtro);
      this.rows = [];
    },
    listData(){
      request.listData(this.filtro).then((res)=>{
        if(res.data?.status_code == 200){
            this.rows = res.data?.embedded
            console.log(this.rows);
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
    resetForm() {
      this.filtro = {
        id_cotacao: [],
        datainicio: "",
        datafim: "",
        categoria: null,
        subcategoria: "",
        situacao: "",
        processo_sei: "",
      };
    }
  }
});
