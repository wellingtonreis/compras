<template>
  <q-layout view="lHh Lpr lFf">
    <q-header class="bg-white shadow-custom"> 
      <!-- bordered -->
      <q-toolbar class="bg-white text-primary">
        <div class="q-pa-md" style="max-width: 480px">
          <q-list>
            <q-item>
              <q-item-section top thumbnail class="q-ml-none">
                <img alt="Logo" src="~assets/sig-logo.png" width="95px">
              </q-item-section>
    
              <q-separator vertical inset="item" color="grey-5" spaced="md" />
    
              <q-item-section class="text-grey-7 text-body1">
                <q-item-label>Empresa Brasileira de Serviços Hospitalares</q-item-label>
              </q-item-section>
            </q-item>
          </q-list>
        </div>

        <q-space />

        <q-btn flat dense icon-right="arrow_drop_down">
          <q-avatar color="primary" text-color="white">W</q-avatar>

          <q-menu>
            <div class="row no-wrap q-pa-md" style="font-size: small;">
              <div class="column">
                <div class="text-h6 q-mb-md">
                  <div class="bg-blue-1 q-mt-md text-center text-caption q-mb-xs" style="padding: 1rem; font-size: 10px;">
                      <q-icon name="perm_contact_calendar" size="22px" />
                      WELLINGTON GOMES DE ANDRADE REIS <br> 
                      014.985.781-01
                  </div>
                </div>
                <q-list separator>
                    <q-item clickable v-ripple active active-class="text-blue">
                      <q-item-section class="text-primary">
                        <q-item-label>Dados Pessoais</q-item-label>
                      </q-item-section>
                    </q-item>

                    <q-item clickable v-ripple active active-class="text-blue">
                      <q-item-section class="text-primary">
                        <q-item-label>Segurança</q-item-label>
                      </q-item-section>
                    </q-item>

                    <q-item clickable v-ripple active active-class="text-blue">
                      <q-item-section class="text-primary">
                        <q-item-label>Privacidade</q-item-label>
                      </q-item-section>
                    </q-item>
                  </q-list>
              </div>

              <q-separator vertical inset class="q-mx-lg" />

              <div class="column items-center justify-between">
                <q-avatar color="primary" text-color="white" size="72px">W</q-avatar>
                <q-btn
                  outline
                  color="primary"
                  label="Sair"
                  icon-right="logout"
                  push
                  size="sm"
                  v-close-popup
                />
              </div>
            </div>
          </q-menu>
        </q-btn>
      </q-toolbar>
      <q-toolbar class="bg-white text-primary">
        <q-btn
          flat
          dense
          round
          icon="menu"
          aria-label="Menu"
          color="primary"
          @click="openLeftDrawer"
        />

        <q-toolbar-title class="text-grey-7 text-body1">
          Pesquisa de preço
        </q-toolbar-title>
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      :width="300"
      :breakpoint="500"
      overlay
      bordered
    >
      <q-list class="text-grey-7 text-body2">
        <q-toolbar>
          <q-toolbar-title class="text-body2">Pesquisa de preço</q-toolbar-title>
          <q-btn flat @click="leftDrawerOpen = false" round dense icon="close" color="primary" />
        </q-toolbar>
        <q-separator />

        <EssentialLink
          v-for="link in linksList"
          :key="link.title"
          v-bind="link"
        />
      </q-list>
    </q-drawer>

    <q-page-container class="text-grey-7">
      <router-view />
    </q-page-container>
  </q-layout>
</template>

<script setup>
import { ref } from "vue";
import EssentialLink from "components/EssentialLink.vue";

defineOptions({
  name: "MainLayout",
});

const linksList = [
  {
    title: "Histórico de cotação",
    caption: "Lista o histórico de cotações",
    icon: "search",
    link: "/historico-cotacao/listar",
  },
];

const leftDrawerOpen = ref(false);

function openLeftDrawer() {
  leftDrawerOpen.value = !leftDrawerOpen.value;
}
</script>
