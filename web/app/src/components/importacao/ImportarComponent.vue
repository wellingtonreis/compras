<template>
  <div>
    <q-form
      @submit="useImportarStore.submitForm()"
      @reset="useImportarStore.resetForm()"
    >
      <div class="row q-pa-md items-center justify-center bg-grey-1">
        <div class="col-12 col-md-2 text-right sm:text-left q-pa-md">
          <b>ID da Cotação:</b>
        </div>
      </div>

      <div class="row items-center justify-center">
        <div class="col-12 col-md-12">
          
        </div>
      </div>

      <div class="row items-center justify-end q-pa-md">
        <q-btn label="Filtrar" type="submit" color="primary" />
        <q-btn
          label="Limpar"
          type="reset"
          color="primary"
          outline
          class="q-ml-sm"
        />
      </div>
    </q-form>
  </div>
</template>

<script setup>
import { storeToRefs } from "pinia";
import { useImportarStore } from "stores/importacao/useImportarStore.js";

const importarStore = useImportarStore();
const {
  
} = storeToRefs(importarStore);
</script>
