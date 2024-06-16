const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') }
    ]
  },

  {
    path: '/consultar-cotacao',
    component: () => import('layouts/MainLayout.vue'),
    children: [
        {name: 'pesquisar-itens-compras', path: 'pesquisar-itens-compras', component: () => import('pages/consultar-cotacao/PesquisarItensCompras.vue')},
    ]
  },

  {
    path: '/importacao',
    component: () => import('layouts/MainLayout.vue'),
    children: [
        {name: 'importar-itens-compras', path: 'importar-itens-compras', component: () => import('pages/importacao/ImportarItensCompras.vue')},
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
