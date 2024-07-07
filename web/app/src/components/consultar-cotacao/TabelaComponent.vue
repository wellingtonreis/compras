<template>
  <div>
    <q-table
      :rows="rows"
      :columns="columns"
      row-key="name"
      separator="cell"
      no-data-label="Não existe nenhum registro."
      no-results-label="Nenhum registro encontrado."
      loading-label="Carregando"
      class="table-custom-header-table"
    >
      <template v-slot:top>
        <div class="col" align="right">
          <q-btn color="primary" label="Adicionar cotação" icon="add" to="#" />
        </div>
      </template>

      <template v-slot:header="props">
        <q-tr :props="props">
          <q-th v-for="col in props.cols" :key="col.name" :props="props">
            {{ col.label }}
          </q-th>
        </q-tr>
      </template>

      <template v-slot:body-cell-acao="props">
        <q-td :props="props">
          <div class="q-pa-md q-gutter-sm">
            <q-btn color="primary" round flat size="sm" icon="visibility" />
            <q-btn
              color="primary"
              round
              flat
              size="sm"
              icon="edit"
              title="Editar"
            />
            <q-btn color="primary" round flat size="sm" icon="delete" />
            <q-btn color="primary" round flat size="sm" icon="history" />
            <q-btn color="primary" round flat size="sm" icon="picture_as_pdf" />
          </div>
        </q-td>
      </template>
    </q-table>
  </div>
</template>

<script setup>
import { onMounted } from 'vue';
import { storeToRefs } from "pinia";
import { useConsultarCotacaoStore } from "@/stores/consultar-cotacao/useConsultarCotacaoStore.js";

const consultarCotacaoStore = useConsultarCotacaoStore();
const {
    columns,
    rows
} = storeToRefs(consultarCotacaoStore);

onMounted(()=>{
  consultarCotacaoStore.listaHitoricoCotacao();
  consultarCotacaoStore.listaSuspensaCategoriaEhSubcategoria();
})
</script>