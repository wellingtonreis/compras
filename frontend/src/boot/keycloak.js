import { boot } from 'quasar/wrappers'
import Keycloak from 'keycloak-js'

import { useKeycloakStore } from "@/stores/keycloak/useKeycloakStore.js";

export default boot(({ app, router }) => {
  async function createRefreshTokenTimer(keycloak) {
    setInterval(() => {
      keycloak.updateToken(60).then((refreshed) => {
        if (refreshed) {
          console.log("Token atualizado", refreshed)
        } else {
          console.warn("Token não foi atualizado, ou não é necessário")
        }
      }).catch(() => {
        console.error("Falha ao atualizar o token")
      })
    }, 120000)
  }

  return new Promise((resolve, reject) => {
    const keycloak = new Keycloak({
      url: process.env.KEYCLOAK_URL,
      realm: process.env.KEYCLOAK_REALM,
      clientId: process.env.KEYCLOAK_CLIENT_ID,
      // credentials: {
      //   secret: process.env.KEYCLOAK_SECRET_ID
      // }
    });

    keycloak
      .init({
        onLoad: "login-required",
        checkLoginIframe: false,
        enableLogging: true
      })
      .then(async (authenticated) => {
        if (authenticated) {
          await createRefreshTokenTimer(keycloak);

          app.config.globalProperties.$keycloak = keycloak;

          const keycloakStore = useKeycloakStore()
          keycloakStore.setKeycloak(keycloak)

          // Guardar as rotas
          router.beforeEach((to, from, next) => {
            if (to.matched.some(record => record.meta.requiresAuth)) {
              if (!keycloak.authenticated) {
                keycloak.login();
              } else {
                next();
              }
            } else {
              next();
            }
          });

          resolve();
        } else {
          keycloak.logout();
        }
      })
      .catch((error) => {
        reject(error);
        keycloak.logout();
      });

      app.config.globalProperties.$keycloakLogout = () => {
        keycloak.logout({ redirectUri: window.location.origin });
      };
  });
});
