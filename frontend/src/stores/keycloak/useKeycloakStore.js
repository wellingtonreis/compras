import { defineStore } from "pinia";

export const useKeycloakStore = defineStore("useKeycloakStore", {
  state: () => ({
    keycloak: null,
    username: null,
    cpf: null,
  }),
  actions: {
    setKeycloak(keycloak) {
      this.keycloak = keycloak
      this.username = keycloak.tokenParsed?.name
      this.cpf = keycloak.tokenParsed?.cpf
    },
    logout() {
      if (this.keycloak) {
        this.keycloak.logout({ redirectUri: window.location.origin })
      }
    },
  },
});
