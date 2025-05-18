<template>
  <div class="dashboard-layout">
    <!-- Sidebar -->
    <aside class="sidebar">
      <div class="sidebar-header">
        <h1 class="logo">BillBridge</h1>
      </div>

      <nav class="sidebar-nav">
        <router-link 
          v-for="item in menuItems" 
          :key="item.path"
          :to="item.path"
          class="nav-item"
          :class="{ active: currentRoute === item.path }"
        >
          <i :class="item.icon"></i>
          <span>{{ item.name }}</span>
        </router-link>
      </nav>

      <div class="sidebar-footer">
        <div class="user-info">
          <span class="username">{{ username }}</span>
        </div>
        <button @click="handleLogout" class="logout-btn">
          <i class="fas fa-sign-out-alt"></i>
          Sair
        </button>
      </div>
    </aside>

    <!-- Main Content -->
    <div class="main-content">
      <router-view></router-view>
    </div>
  </div>
</template>
<script>
export default {
  name: 'DashboardLayout'
}
</script>
<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import authService from '@/services/authService'


const params = new URLSearchParams(window.location.search);
const token = params.get('token');
const userJson = params.get('user');
const username = ref('')
if (token && userJson) {
  try {

    authService.saveAuthData(token, userJson);

 
    window.history.replaceState({}, document.title, window.location.pathname);
  } catch (e) {
    console.error('Erro ao parsear user JSON:', e);
  }
}

onMounted(() => {
  username.value = JSON.parse(authService.getUsername()).username;
  router.push('/dashboard/overview');
})

const router = useRouter()
const route = useRoute()




const currentRoute = computed(() => route.path)

const menuItems = [
  { 
    name: 'Dashboard', 
    path: '/dashboard/overview', 
    icon: 'fas fa-chart-line' 
  },
  { 
    name: 'Faturas', 
    path: '/dashboard/invoices', 
    icon: 'fas fa-file-invoice' 
  },
  { 
    name: 'Nova Fatura', 
    path: '/dashboard/new-invoice', 
    icon: 'fas fa-plus' 
  },
  { 
    name: 'Analytics', 
    path: '/dashboard/analytics', 
    icon: 'fas fa-chart-pie' 
  }
]

const handleLogout = () => {
  authService.clearAuthData();
  router.push('/')
}
</script>

<style scoped>
.dashboard-layout {
  display: flex;
  min-height: 100vh;
}

.sidebar {
  width: 260px;
  background: var(--primary-color, #1f2937);
  color: white;
  display: flex;
  flex-direction: column;
  position: fixed;
  left: 0;
  top: 0;
  height: 100vh;
  box-shadow: 2px 0 5px rgba(0, 0, 0, 0.5);
  transition: background 0.3s;  
  background: var(--sidebar-background, #1f2937); 
}

.sidebar-header {
  padding: 1.5rem;
  text-align: center;
  background: rgba(255, 255, 255, 0.1);
}

.logo {
  font-size: 1.5rem;
  font-weight: bold;
  margin: 0;
}

.sidebar-nav {
  flex: 1;
  padding: 1rem 0;
}

.nav-item {
  display: flex;
  align-items: center;
  padding: 0.75rem 1.25rem;
  color: white;
  text-decoration: none;
  transition: background 0.3s;
  border-radius: 0.5rem;
  gap: 0.75rem;
}

.nav-item i {
  margin-right: 0.5rem;
}

.nav-item:hover {
  background: #374151;
}

.active {
  background: #2563eb;
}

.sidebar-footer {
  padding: 1rem;
  background: rgba(255, 255, 255, 0.05);
  text-align: center;
}

.user-info {
  margin-bottom: 0.5rem;
}

.username {
  font-weight: 600;
  font-size: larger;
}

.logout-btn {
  background: #ef4444;
  border: none;
  color: white;
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  cursor: pointer;
  font-size: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 0.5rem;
  transition: background 0.3s;
  margin-left: auto;
  margin-right: auto;
  width: 100%;
} 

.main-content {
  flex: 1;
  margin-left: 250px;
  padding: 2rem;
  background-color: var(--gray-100);
}
</style>