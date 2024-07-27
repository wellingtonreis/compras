<template>
  <div class="q-pa-md">
    <q-card flat bordered>
      <q-card-section class="row justify-center">
        <div class="q-pa-md" style="width: 100%">
          <PesquisarComponent />
        </div>
      </q-card-section>

      <q-separator inset />

      <q-card-section>
        <TabelaComponent />
      </q-card-section>
    </q-card>
  </div>
</template>

<script>
import PesquisarComponent from "components/historico-cotacao/PesquisarComponent.vue";
import TabelaComponent from "components/historico-cotacao/TabelaComponent.vue";

export default {
  components: {
    PesquisarComponent,
    TabelaComponent
  }
};
</script>
