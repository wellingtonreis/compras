import { defineStore } from 'pinia';
export const usePesquisarStore = defineStore('usePesquisarStore', {
  state: () => ({
    filtro: {
      id_cotacao: [],
      datainicio: "",
      datafim: "",
      categoria: null,
      subcategoria: "",
      situacao: "",
      processo_sei: "",
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
  }),
  actions: {
    submitForm() {
      console.log("Filtro:", this.filtro);
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
