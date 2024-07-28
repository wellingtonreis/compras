<template>
  <div class="q-pa-md">
    <div class="row q-ma-md justify-center">
      <q-chip size="xl" icon="calculate">
        Cotação nº {{ cotacao }}
      </q-chip>
    </div>

    <q-stepper
      v-model="step"
      ref="stepper"
      color="primary"
      animated
    >
      <q-step
        :name="1"
        title="Editar Cotação"
        icon="edit"
        :done="step > 1"
        style="min-height: 200px;"
      >
        <TabelaComponent />
      </q-step>

      <q-step
        :name="2"
        title="Cadastrar segmentação e processo SEI"
        icon="add"
        :done="step > 2"
        style="min-height: 200px;"
      >
        <CadastrarComponent />
      </q-step>

      <q-step
        :name="3"
        title="Detalhamento da cotação"
        icon="assignment"
        style="min-height: 200px;"
      >
        Conteúdo
      </q-step>

      <template v-slot:navigation>
        <q-stepper-navigation>
          <q-btn @click="$refs.stepper.next()" color="primary" :label="step === 3 ? 'Calcular cotação' : 'Avançar'" />
          <q-btn v-if="step > 1" flat color="primary" @click="$refs.stepper.previous()" label="Voltar" class="q-ml-sm" />
        </q-stepper-navigation>
      </template>

      <template v-slot:message>
        <q-banner v-if="step === 1" class="bg-grey-2 text-primary q-px-lg">
          Atualize as informações e valores da cotação
        </q-banner>
        <q-banner v-else-if="step === 2" class="bg-grey-2 text-primary q-px-lg">
          Selecione o segmento da cotação e adicione processo SEI
        </q-banner>
        <q-banner v-else class="bg-grey-2 text-primary q-px-lg">
          Visualização das amostras da pesquisa de preço desta cotação
        </q-banner>
      </template>
    </q-stepper>
  </div>
</template>

<script>
import TabelaComponent from "components/itens-compras/TabelaComponent.vue";
import CadastrarComponent from "components/itens-compras/CadastrarComponent.vue";

import { useRoute } from 'vue-router';
import { ref } from 'vue'

export default {
  components: {
    TabelaComponent,
    CadastrarComponent
  },
  setup () {
    return {
      step: ref(1),
      cotacao: ref(0)
    }
  },
  mounted () {
    const route = useRoute()
    this.cotacao = ref(route.params.cotacao);
  }
};
</script>
