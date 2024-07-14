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
      <template v-slot:header="props">
        <q-tr :props="props">
          <q-th v-for="col in props.cols" :key="col.name" :props="props">
            {{ col.label }}
          </q-th>
        </q-tr>
      </template>

      <template v-slot:body-cell-idCompra="props">
        <q-td :props="props">
          <span v-html="itensComprasStore.formatIdCompra(props.row)"></span>
        </q-td>
      </template>

      <template v-slot:body-cell-precoUnitario="props">
        <q-td :props="props">
          <div class="row items-center">
            <div class="col">
              <q-chip clickable icon="edit">{{ props.row.precoUnitario }}</q-chip>
              <q-popup-edit 
                v-model.number="props.row.precoUnitario" 
                buttons
                label-set="CONFIRMAR"
                :validate="itensComprasStore.validarPrecoUnitario"
                v-slot="scope">

                <q-input 
                :maxlength="15"
                v-model.number="scope.value" 
                :error="erroPrecoUnitario"
                :error-message="erroMensagemPrecoUnitario"
                hint="Somente ponto para casas decimais. Exemplo: 10.00"
                dense 
                autofocus 
                counter />
              </q-popup-edit>
            </div>
          </div>
        </q-td>
      </template>

      <template v-slot:body-cell-justificativa="props">
        <q-td :props="props">
          <div class="q-pa-md q-gutter-sm">
            <q-btn color="primary" round flat size="md" icon="history" />
          </div>
        </q-td>
      </template>

    </q-table>
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router';
import { onMounted } from 'vue';
import { storeToRefs } from "pinia";
import { useItensComprasStore } from "@/stores/itens-compras/useItensComprasStore.js";

const itensComprasStore = useItensComprasStore();
const {
    columns,
    rows,
    erroPrecoUnitario,
    erroMensagemPrecoUnitario
} = storeToRefs(itensComprasStore);

onMounted(()=>{
  const route = useRoute()
  const cotacao = route.params.cotacao;
  itensComprasStore.listaItensCompras(cotacao);
})
</script>