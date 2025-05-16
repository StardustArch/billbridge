import { createRouter, createWebHistory } from 'vue-router'

// Componentes principais
import HelloWorld from '@/components/HelloWorld.vue'
import UserDashboard from '@/components/Dashboard/Dashboard.vue'
import MainContent from '@/components/Dashboard/MainContent/MainContent.vue'
import Invoices from '@/components/Dashboard/Invoices/Invoices.vue'
import Analytics from '@/components/Dashboard/Analytics/Analytics.vue'

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
    children: [
      {
        path: '',
        redirect: { name: 'Overview' }
      },
      {
        path: 'overview',
        name: 'Overview',
        component: MainContent
      },
      {
        path: 'invoices',
        name: 'Invoices',
        component: Invoices
      },
      {
        path: 'analytics',
        name: 'Analytics',
        component: Analytics
      }
    ]
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
