<template>
  <div>
    <q-form
      @submit="historicoCotacaoStore.enviar()"
      @reset="historicoCotacaoStore.reiniciar()"
    >
      <div class="row q-pa-md items-center justify-center bg-grey-1">
        <div class="col-12 col-md-2 text-right sm:text-left q-pa-md">
          <b>ID da Cotação:</b>
        </div>
        <div class="col-12 col-md-4">
          <q-select
            label="Digite o número da cotação e pressione Enter."
            outlined
            v-model="filtro.cotacao"
            use-input
            use-chips
            hide-dropdown-icon
            input-debounce="0"
            @new-value="createValue"
          />
        </div>
        <div class="col-12 col-md-2 q-pa-md">
          <q-toggle
            v-model="pesquisa_avancada"
            label="Pesquisa avançada"
            class="q-mb-md"
          />
        </div>
      </div>

      <div class="row items-center justify-center">
        <div class="col-12 col-md-12">
          <q-expansion-item
            v-model="pesquisa_avancada"
            icon="search"
            label="Pesquisa avançada"
            caption="Clique aqui para exibir os filtros da pesquisa avançada"
            dense
          >
            <div class="row q-pa-md items-start justify-left">
              <div class="col-12 col-md-2 text-right sm:text-left q-mr-md">
                <q-btn flat color="primary" icon="event" size="md">
                  Selecione aqui o período
                  <q-popup-proxy
                    cover
                    transition-show="scale"
                    transition-hide="scale"
                  >
                    <q-date
                      v-model="date_between"
                      range
                      @range-end="
                        (range) => {
                          const fromDay =
                            range.from.day < 10
                              ? `0${range.from.day}`
                              : range.from.day;
                          const toDay =
                            range.to.day < 10
                              ? `0${range.to.day}`
                              : range.to.day;

                          const fromMonth =
                            range.from.month < 10
                              ? `0${range.from.month}`
                              : range.from.month;
                          const toMonth =
                            range.to.month < 10
                              ? `0${range.to.month}`
                              : range.to.month;

                          filtro.data_inicio = `${fromDay}/${fromMonth}/${range.from.year}`;
                          filtro.data_fim = `${toDay}/${toMonth}/${range.to.year}`;
                        }
                      "
                    >
                      <q-btn label="Fecha" color="primary" flat v-close-popup />
                    </q-date>
                  </q-popup-proxy>
                </q-btn>
              </div>
              <div class="col-12 col-md-8 q-pr-md">
                <div class="row q-col-gutter-sm">
                  <div class="col-12 col-md-6">
                    <q-input
                      outlined
                      dense
                      v-model="filtro.data_inicio"
                      mask="##/##/####"
                      hint="Data de Início"
                      disable
                    >
                      <template v-slot:append>
                        <q-icon name="event" class="cursor-pointer">
                          <q-popup-proxy
                            cover
                            transition-show="scale"
                            transition-hide="scale"
                          >
                            <q-date
                              v-model="filtro.data_inicio"
                              mask="DD/MM/YYYY"
                            >
                              <div class="row items-center justify-end">
                                <q-btn
                                  v-close-popup
                                  label="Fechar"
                                  color="primary"
                                  flat
                                />
                              </div>
                            </q-date>
                          </q-popup-proxy>
                        </q-icon>
                      </template>
                    </q-input>
                  </div>

                  <div class="col-12 col-md-6">
                    <q-input
                      outlined
                      dense
                      v-model="filtro.data_fim"
                      mask="##/##/####"
                      hint="Data Fim"
                      disable
                    >
                      <template v-slot:append>
                        <q-icon name="event" class="cursor-pointer">
                          <q-popup-proxy
                            cover
                            transition-show="scale"
                            transition-hide="scale"
                          >
                            <q-date v-model="filtro.data_fim" mask="DD/MM/YYYY">
                              <div class="row items-center justify-end">
                                <q-btn
                                  v-close-popup
                                  label="Fechar"
                                  color="primary"
                                  flat
                                />
                              </div>
                            </q-date>
                          </q-popup-proxy>
                        </q-icon>
                      </template>
                    </q-input>
                  </div>
                </div>
              </div>
            </div>

            <div class="row q-pa-md items-center justify-left">
              <div class="col-12 col-md-2 text-right sm:text-left q-mr-md">
                <b>Categoria:</b>
              </div>
              <div class="col-12 col-md-8">
                <q-select
                  outlined
                  v-model="filtro.categoria"
                  use-input
                  hide-selected
                  fill-input
                  input-debounce="0"
                  :options="opcoesCategorias"
                  @filter="historicoCotacaoStore.filtraCategoria"
                  @update:model-value="historicoCotacaoStore.selecionaSubcategoria(filtro.categoria.value)"
                  hint="Lista suspensa de categorias"
                  dense
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
                  v-model="filtro.subcategoria"
                  use-input
                  hide-selected
                  fill-input
                  input-debounce="0"
                  :options="opcoesSubCategorias"
                  @filter="historicoCotacaoStore.filtraSubcategoria"
                  hint="Lista suspensa de subcategorias"
                  dense
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
                <b>Situação:</b>
              </div>
              <div class="col-12 col-md-8">
                <q-select
                  outlined
                  v-model="filtro.situacao"
                  use-input
                  hide-selected
                  fill-input
                  input-debounce="0"
                  :options="opcoesSituacao"
                  @filter="filtraSituacao"
                  hint="Lista suspensa de situação"
                  dense
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
                <b>Processo SEI:</b>
              </div>
              <div class="col-12 col-md-8">
                <q-input outlined v-model="text" dense />
              </div>
            </div>
          </q-expansion-item>
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
import { ref, toRaw } from "vue";
import { storeToRefs } from "pinia";
import { useHistoricoCotacaoStore } from "stores/historico-cotacao/useHistoricoCotacaoStore.js";

const historicoCotacaoStore = useHistoricoCotacaoStore();
const text = ref(null);
const {
  filtro,
  date_between,
  pesquisa_avancada,
  opcoesCategorias,
  opcoesSubCategorias,
  opcoesSituacao,
} = storeToRefs(historicoCotacaoStore);

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

const situacaoFiltro = toRaw(opcoesSituacao.value);
const filtraSituacao = (val, update) => {
  update(() => {
    const needle = val.toLowerCase();
    const filtragem = situacaoFiltro?.filter(
      (v) => v.toLowerCase().indexOf(needle) > -1
    );

    if (filtragem?.length > 0) {
      opcoesSituacao.value = filtragem;
    }
  });
};

const createValue = (val, done) => {
  done(val, "add-unique");
};
</script>