// Composables
import { createRouter, createWebHashHistory, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    children: [
      {
        path: '',
        name: 'AddLantanaPage',
        component: () => import('@/views/add_lantana_page.vue')
      },
      {
        path: '/lantana_log_viewer',
        name: 'LantanaLogViewerPage',
        component: () => import('@/views/lantana_log_viewer_page.vue'),
      }
    ],
  },
]

const router = createRouter({
  history: createWebHashHistory(),
  // history: createWebHistory(process.env.BASE_URL),
  routes,
})

export default router
