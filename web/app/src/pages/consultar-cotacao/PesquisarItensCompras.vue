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
              <q-btn
                color="primary"
                label="Adicionar cotação"
                icon="add"
                to="#"
              />
            </div>
          </template>

          <template v-slot:header="props">
            <q-tr :props="props">
              <q-th
                v-for="col in props.cols"
                :key="col.name"
                :props="props"
              >
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
                />
                <q-btn color="primary" round flat size="sm" icon="delete" />
                <q-btn color="primary" round flat size="sm" icon="history" />
                <q-btn color="primary" round flat size="sm" icon="picture_as_pdf" />
              </div>
            </q-td>
          </template>
        </q-table>
      </q-card-section>
    </q-card>
  </div>
</template>

<script>
import PesquisarComponent from "components/consultar-cotacao/PesquisarComponent.vue";

const columns = [
  {
    name: "acao",
    required: true,
    label: "AÇÕES",
    align: "left",
    field: (row) => row.acao,
    format: (val) => `${val}`,
    sortable: true,
  },
  {
    name: "id",
    required: true,
    label: "ID",
    align: "left",
    field: (row) => row.id,
    format: (val) => `${val}`,
    sortable: true,
  },
  {
    name: "hu",
    required: true,
    label: "HU",
    align: "left",
    field: (row) => row.hu,
    format: (val) => `${val}`,
    sortable: true,
  },
  {
    name: "categoria",
    required: true,
    label: "CATEGORIA",
    align: "left",
    field: (row) => row.categoria,
    format: (val) => `${val}`,
    sortable: true,
  },
  {
    name: "subcategoria",
    required: true,
    label: "SUBCATEGORIA",
    align: "left",
    field: (row) => row.subcategoria,
    format: (val) => `${val}`,
    sortable: true,
  },
  {
    name: "data",
    required: true,
    label: "DATA",
    align: "left",
    field: (row) => row.data,
    format: (val) => `${val}`,
    sortable: true,
  },
  {
    name: "situacao",
    required: true,
    label: "SITUAÇÃO",
    align: "left",
    field: (row) => row.situacao,
    format: (val) => `${val}`,
    sortable: true,
  },
  {
    name: "processo_sei",
    required: true,
    label: "PROCESSO SEI",
    align: "left",
    field: (row) => row.processo_sei,
    format: (val) => `${val}`,
    sortable: true,
  },
  {
    name: "autor",
    required: true,
    label: "AUTOR",
    align: "left",
    field: (row) => row.autor,
    format: (val) => `${val}`,
    sortable: true,
  },
];

const rows = [
  {
    id: 1,
    hu: "Hospital 1",
    categoria: "Medicamento 1",
    subcategoria: "Remédio 1",
    data: "11/05/2024 18:00",
    situacao: "Iniciada",
    processo_sei: "99999.999999/9999-99",
    autor: "Colaborador 1",
  },
  {
    id: 2,
    hu: "Hospital 2",
    categoria: "Medicamento 2",
    subcategoria: "Remédio 2",
    data: "11/05/2024 19:00",
    situacao: "Aguardando",
    processo_sei: "99999.999999/9999-98",
    autor: "Colaborador 2",
  },
  {
    id: 3,
    hu: "Hospital 3",
    categoria: "Medicamento 3",
    subcategoria: "Remédio 3",
    data: "11/05/2024 20:00",
    situacao: "Concluída",
    processo_sei: "99999.999999/9999-97",
    autor: "Colaborador 3",
  },
  {
    id: 4,
    hu: "Hospital 4",
    categoria: "Medicamento 4",
    subcategoria: "Remédio 4",
    data: "11/05/2024 21:00",
    situacao: "Cancelada",
    processo_sei: "99999.999999/9999-96",
    autor: "Colaborador 4",
  },
  {
    id: 5,
    hu: "Hospital 5",
    categoria: "Medicamento 5",
    subcategoria: "Remédio 5",
    data: "11/05/2024 22:00",
    situacao: "Em análise",
    processo_sei: "99999.999999/9999-95",
    autor: "Colaborador 5",
  }
];

export default {
  components: {
    PesquisarComponent,
  },
  setup() {
    return {
      columns,
      rows,
    };
  },
};
</script>
