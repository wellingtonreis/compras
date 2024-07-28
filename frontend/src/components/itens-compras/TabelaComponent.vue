<template>
  <div>
    <q-table
      style="height: 500px"
      :rows="rows"
      :columns="columns"
      row-key="name"
      separator="cell"
      virtual-scroll
      v-model:pagination="pagination"
      :rows-per-page-options="[0]"
      no-data-label="Não existe nenhum registro."
      no-results-label="Nenhum registro encontrado."
      loading-label="Carregando"
      class="table-custom-header-table"
    >
      <template v-slot:top>
        <q-btn color="primary" icon="add" label="Adicionar item de compra" />
      </template>

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
                @save="(valor, valorInicial) => { itensComprasStore.atualizarPrecoUnitario(cotacao, props.row, valor, valorInicial) }"
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
            <q-btn color="primary" round flat size="md" icon="history" @click="itensComprasStore.exibirHistorico(props.row)" />
          </div>
        </q-td>
      </template>

      <template v-slot:body-cell-remover="props">
        <q-td :props="props">
          <div class="q-pa-md q-gutter-sm">
            <q-btn color="red" round flat size="md" icon="delete" @click="itensComprasStore.removerItemDeCompra(props.rowIndex, cotacao, props.row)" />
          </div>
        </q-td>
      </template>

    </q-table>

    <q-dialog v-model="fixed" full-width backdrop-filter="blur(4px)">
      <q-card style="width: 100%;">
        <q-card-section>
          <div class="text-h6">Histórico de justificativa</div>
          <div class="text-h6 text-grey-6">
            {{ catmat }} - {{ descricao }} - {{ apresentacao }}
          </div>
        </q-card-section>

        <q-separator />

        <q-card-section>
          <div class="q-px-lg q-py-md">
            <q-scroll-area style="height: 260px;">
              <q-timeline layout="comfortable" side="right" color="secondary">  
                <q-timeline-entry v-for="j in justificativas" :key="j">
                  <template v-slot:title>
                    {{ j.descricao }}
                    <div class="text-overline">
                      {{ j.autor }} 
                      alterou o valor inicial <q-badge outline color="red" :label="itensComprasStore.formatarMoeda(j.valorInicial)" /> para <q-badge outline color="green" :label="itensComprasStore.formatarMoeda(j.valor)" />
                    </div>
                  </template>
                  <template v-slot:subtitle>
                    <div class="text-caption">
                    {{ new Date(j.data + 'T00:00:00Z').toLocaleDateString('pt-BR', {
                      timeZone: 'UTC',
                      weekday: 'long', 
                      year: 'numeric',
                      month: 'long',
                      day: 'numeric'
                    }) }}
                    </div>
                  </template>
                </q-timeline-entry>
              </q-timeline>
            </q-scroll-area>
          </div>
        </q-card-section>

        <q-separator />

        <q-card-actions align="right">
          <q-btn flat label="Ok" color="primary" v-close-popup />
        </q-card-actions>
      </q-card>
    </q-dialog>
  </div>
</template>

<script setup>
import { useRoute } from 'vue-router';
import { ref, onMounted } from 'vue';
import { storeToRefs } from "pinia";
import { useKeycloakStore } from "@/stores/keycloak/useKeycloakStore.js";
import { useItensComprasStore } from "@/stores/itens-compras/useItensComprasStore.js";

const keycloakStore = useKeycloakStore();
const {
    username
} = storeToRefs(keycloakStore);

const itensComprasStore = useItensComprasStore();
const {
    autor,
    columns,
    rows,
    erroPrecoUnitario,
    erroMensagemPrecoUnitario,
    pagination,
    fixed,
    catmat,
    descricao,
    apresentacao,
    justificativas
} = storeToRefs(itensComprasStore);

const route = useRoute()
const cotacao = ref(route.params.cotacao);

onMounted(()=>{
  autor.value = ref(username)
  itensComprasStore.listaItensCompras(cotacao.value.toString());
})
</script>