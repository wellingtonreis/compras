const routes = [
  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      { path: '', component: () => import('pages/IndexPage.vue') },
      { name: 'historico-cotacao-listar', path: 'historico-cotacao/listar', component: () => import('pages/historico-cotacao/HistoricoCotacao.vue')},
      { name: 'itens-de-compras-listar', path: 'itens-de-compras/listar/:cotacao', component: () => import('pages/itens-compras/ListarItensCompras.vue')},
    ]
  },

  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
