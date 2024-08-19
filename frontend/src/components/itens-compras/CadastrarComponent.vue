<template>
    <div>
      <q-form @reset="itensComprasStore.reiniciar()">
        <div class="row items-center justify-center">
          <div class="col-12 col-md-12">
            <div class="row q-pa-md items-center justify-left">
                <div class="col-12 col-md-2 text-right sm:text-left q-mr-md">
                    <b>Categoria:</b>
                </div>
                <div class="col-12 col-md-8">
                    <q-select
                        outlined
                        v-model="categoria"
                        use-input
                        hide-selected
                        fill-input
                        input-debounce="0"
                        :options="opcoesCategorias"
                        @filter="itensComprasStore.filtraCategoria"
                        @update:model-value="itensComprasStore.selecionaSubcategoria(categoria.value)"
                        hint="Lista suspensa de categorias"
                        dense
                        bottom-slots
                        error-message="Por favor, selecione uma categoria"
                        :error="!validaCategoria"
                    >
                        <template v-slot:no-option>
                            <q-item>
                                <q-item-section class="text-grey">
                                    No results
                                </q-item-section>
                            </q-item>
                        </template>
                    </q-select>
                </div>
            </div>

            <div class="row q-pa-md items-center justify-left">
                <div class="col-12 col-md-2 text-right sm:text-left q-mr-md">
                    <b>Subcategoria:</b>
                </div>
                <div class="col-12 col-md-8">
                    <q-select
                        outlined
                        v-model="subcategoria"
                        use-input
                        hide-selected
                        fill-input
                        input-debounce="0"
                        :options="opcoesSubCategorias"
                        @filter="itensComprasStore.filtraSubcategoria"
                        hint="Lista suspensa de subcategorias"
                        dense
                        bottom-slots
                        error-message="Por favor, selecione uma subcategoria"
                        :error="!validaSubcategoria"
                    >
                        <template v-slot:no-option>
                            <q-item>
                                <q-item-section class="text-grey">
                                    No results
                                </q-item-section>
                            </q-item>
                        </template>
                    </q-select>
                </div>
            </div>

            <div class="row q-pa-md items-center justify-left">
                <div class="col-12 col-md-2 text-right sm:text-left q-mr-md">
                    <b>Número do processo SEI:</b>
                </div>
                <div class="col-12 col-md-8">
                    <q-input 
                    outlined 
                    v-model="processosei" 
                    @keyup="itensComprasStore.preencheProcessoSei()"
                    dense 
                    mask="#####.######/####-##"
                    error-message="Por favor, adicione um número do processo SEI"
                    :error="!validaProcessosei" />
                </div>
            </div>

            <div class="row q-pa-md items-center justify-left">
                <div class="col-12 col-md-2 text-right sm:text-left q-mr-md">
                    <b>Link processo SEI:</b>
                </div>
                <div class="col-12 col-md-8">
                  <q-input outlined v-model="urlprotocolosei" dense standout bottom-slots
                    :disable="true" prefix="https://sei.ebserh.gov.br/sei/controlador.php?acao=" mask="#################">
                    <template v-slot:hint>
                      Ex.: https://sei.ebserh.gov.br/sei/controlador.php?acao=protocolo
                    </template>
                  </q-input>
                </div>
            </div>
          </div>
        </div>
  
        <div class="row items-center justify-end q-pa-md">
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

  import { toRaw, onMounted } from "vue";
  import { storeToRefs } from "pinia";
  import { useItensComprasStore } from "@/stores/itens-compras/useItensComprasStore.js";
  
  const itensComprasStore = useItensComprasStore();
  const {
    categoria,
    subcategoria,
    processosei,
    urlprotocolosei,
    opcoesCategorias,
    opcoesSubCategorias,
    validaCategoria,
    validaSubcategoria,
    validaProcessosei
  } = storeToRefs(itensComprasStore);
  
  const categoriaFiltro = toRaw(opcoesCategorias.value);
  const filtraCategoria = (val, update) => {
    update(() => {
      const needle = val.toLowerCase();
      const filtragem = categoriaFiltro?.filter(
        (v) => v.toLowerCase().indexOf(needle) > -1
      );
  
      if (filtragem?.length > 0) {
        opcoesCategorias.value = filtragem;
      }
    });
  };
  
  const subCategoriaFiltro = toRaw(opcoesSubCategorias.value);
  const filtraSubCategoria = (val, update) => {
    update(() => {
      const needle = val.toLowerCase();
      const filtragem = subCategoriaFiltro?.filter(
        (v) => v.toLowerCase().indexOf(needle) > -1
      );
  
      if (filtragem?.length > 0) {
        opcoesSubCategorias.value = filtragem;
      }
    });
  };

  onMounted(() => {
    itensComprasStore.listaSuspensaCategoriaEhSubcategoria();
  });
  </script>