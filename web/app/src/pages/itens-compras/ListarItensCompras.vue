<template>
  <div class="q-pa-md">
    <q-card flat bordered>
      <q-card-section>
        <TabelaComponent />
      </q-card-section>
    </q-card>
  </div>
</template>

<script>
import TabelaComponent from "components/itens-compras/TabelaComponent.vue";

export default {
  components: {
    TabelaComponent
  }
};
</script>
