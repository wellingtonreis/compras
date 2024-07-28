<template>
  <div class="q-pa-md">
    <q-card flat bordered>

      <q-card-section class="q-pt-none">
        Bem-vindo {{ username }}, você está logado!
      </q-card-section>
    </q-card>
    
  </div>
</template>

<script>
export default {
  computed: {
    username() {
      return this.$keycloak.tokenParsed?.name;
    }
  }
}
</script>
