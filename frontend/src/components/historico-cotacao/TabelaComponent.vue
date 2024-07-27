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
          <q-btn color="primary" label="Adicionar cotação" icon="add" @click="updateValues" />
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
              :to="{ name: 'itens-de-compras-listar', params: { cotacao: props.row.cotacao }}"
            />
            <q-btn color="primary" round flat size="sm" icon="delete" />
            <q-btn color="primary" round flat size="sm" icon="history" />
            <q-btn color="primary" round flat size="sm" icon="picture_as_pdf" />
          </div>
        </q-td>
      </template>
    </q-table>

    <q-dialog v-model="dialog" 
      persistent
      backdrop-filter="blur(4px) saturate(150%)" 
      :maximized="true"
      transition-show="slide-up"
      transition-hide="slide-down"
      >
      <q-card flat bordered>
        <q-bar>
          <q-space />
          <q-btn dense flat icon="close" v-close-popup>
            <q-tooltip class="bg-white text-primary">Fecha</q-tooltip>
          </q-btn>
        </q-bar>

        <q-card-section>
          <div class="text-h6">Criar nova cotação</div>
          <div class="text-subtitle2">Cotação</div>
        </q-card-section>

        <q-tabs v-model="tab" class="text-teal">
          <q-tab label="Manual" name="manual" />
          <q-tab label="Automatico" name="automatico" />
        </q-tabs>

        <q-separator />

        <q-tab-panels v-model="tab" animated>
          <q-tab-panel name="manual">
            <div class="column items-center">
              <div class="col">
                <q-btn color="primary" label="Crie uma cotação de forma manual" icon="add" style="width: 80vw" />
              </div>
            </div>
          </q-tab-panel>

          <q-tab-panel name="automatico">
            <div class="column justify-center">
              <div class="col">
                <q-uploader 
                flat
                bordered
                field-name="file" 
                url="http://localhost:3000/upload" 
                max-file-size="10240"
                accept=".csv, .xls, .xlsx"
                style="width: 100%"
                @finish="historicoCotacaoStore.uploadFinalizado" />
              </div>
            </div>
          </q-tab-panel>
        </q-tab-panels>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup>
import { ref } from 'vue';
import { onMounted } from 'vue';
import { storeToRefs } from "pinia";
import { useHistoricoCotacaoStore } from "@/stores/historico-cotacao/useHistoricoCotacaoStore.js";

const historicoCotacaoStore = useHistoricoCotacaoStore();
const {
    columns,
    rows
} = storeToRefs(historicoCotacaoStore);

const dialog = ref(false)
const tab = ref('manual')

function updateValues() {
  dialog.value = true;
}

onMounted(()=>{
  historicoCotacaoStore.listaHitoricoCotacao();
  historicoCotacaoStore.listaSuspensaCategoriaEhSubcategoria();
})
</script>