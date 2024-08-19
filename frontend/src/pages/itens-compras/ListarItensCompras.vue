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
          <q-btn @click="nextStep()" color="primary" :label="step === 3 ? 'Calcular cotação' : 'Avançar'" />
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
import { storeToRefs } from "pinia";

import { useItensComprasStore } from "@/stores/itens-compras/useItensComprasStore.js";
const itensComprasStore = useItensComprasStore();
const {
    cotacao,
    categoria,
    subcategoria,
    processosei,
    validaCategoria,
    validaSubcategoria,
    validaProcessosei
} = storeToRefs(itensComprasStore);

export default {
  components: {
    TabelaComponent,
    CadastrarComponent
  },
  setup () {
    return {
      step: ref(1),
      cotacao: cotacao,
      categoria: categoria,
      subcategoria: subcategoria,
      processosei: processosei,
      validaCategoria: validaCategoria,
      validaSubcategoria: validaSubcategoria,
      validaProcessosei: validaProcessosei
    }
  },
  created () {
    const route = useRoute()
    this.cotacao = ref(route.params.cotacao);
  },
  methods:{
    notificaUsuario(msg){
      this.$q.notify({
        color: 'negative',
        position: 'bottom',
        message: msg,
        icon: 'report_problem',
        actions: [
          { icon: 'close', color: 'white', round: true, handler: () => { /* ... */ } }
        ]
      });
    },
    nextStep(){
      const proximoPasso = this.step === 2;
      if(proximoPasso && this.categoria.value == null){
        this.notificaUsuario('Selecione uma categoria');
        this.validaCategoria = false;
        return;
      }
      this.validaCategoria = true;

      if(proximoPasso && this.subcategoria.value == null){
        this.notificaUsuario('Selecione uma Subcategoria');
        this.validaSubcategoria = false;
        return;
      }
      this.validaSubcategoria = true;
      
      if(proximoPasso && this.processosei == null){
        this.notificaUsuario('Adicione um número do processo SEI');
        this.validaProcessosei = false;
        return;
      }

      const numeroProcesso = this.processosei?.replace(/[^0-9]/g, '') ?? 0;
      if(proximoPasso && numeroProcesso.length < 17){
        this.notificaUsuario('Adicione um número do processo SEI');
        this.validaProcessosei = false;
        return;
      }
      this.validaProcessosei = true;

      if(proximoPasso){
        itensComprasStore.salvarClassificacaoSegmento(this.step, this.cotacao)
      }
      this.$refs.stepper.next();
    },
  }
};
</script>
