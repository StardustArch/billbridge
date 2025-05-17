import { createRouter, createWebHistory } from 'vue-router'

// Componentes principais
import HelloWorld from '@/components/Auth.vue'
import UserDashboard from '@/components/Dashboard/Dashboard.vue'
import MainContent from '@/components/Dashboard/MainContent/MainContent.vue'
import Invoices from '@/components/Dashboard/Invoices/Invoices.vue'
import Analytics from '@/components/Dashboard/Analytics/Analytics.vue'
import NewInvoice from '@/components/Dashboard/NewInvoice/NewInvoice.vue'
// import authService from '@/services/authService.js'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: HelloWorld
  },
  {
    path: '/dashboard',
    name: 'Dashboard',
    component: UserDashboard,
    redirect: '/dashboard/overview',
    // meta: { requiresAuth: true },
    children: [
      {
        path: 'overview',
        name: 'Overview',
        component: MainContent
        // meta: { requiresAuth: true }
      },
      {
        path: 'invoices',
        name: 'Invoices',
        component: Invoices
        // meta: { requiresAuth: true }
      },
      {
        path: 'new-invoice',
        name: 'NewInvoice',
        component: NewInvoice
        // meta: { requiresAuth: true }
      },
      {
        path: 'analytics',
        name: 'Analytics',
        component: Analytics
        // meta: { requiresAuth: true }
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// Guard global para checar autenticação
// router.beforeEach((to, from, next) => {
//   const requiresAuth = to.matched.some(record => record.meta.requiresAuth)
//   const isAuthenticated = authService.isAuthenticated() 

//   if (requiresAuth && !isAuthenticated) {
//     // Se a rota requer auth e não está autenticado, redireciona para login
//     next({ name: 'Home' })
//   } else {
//     // Se está autenticado ou rota pública, segue normalmente
//     next()
//   }
// })

export default router
