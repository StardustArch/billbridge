<template>
    <div class="dashboard-content">
      <!-- Cards de Estatísticas -->
      <div class="stats-container">
        <div
          class="stat-card"
          v-for="(stat, index) in animatedStats"
          :key="index"
          :class="{ 'skeleton': loading }"
        >
          <h3 v-if="!loading">{{ stat.label }}</h3>
          <div class="stat-value" v-if="!loading">{{ stat.display }}</div>
        </div>
      </div>
  
      <!-- Tabela de Faturas Recentes -->
      <div class="recent-invoices">
        <h2>Faturas Recentes</h2>
  
        <div v-if="loading" class="table-skeleton">
          <div v-for="n in 5" :key="n" class="skeleton-row"></div>
        </div>
  
        <table v-else class="fade-in">
          <thead>
            <tr>
              <th>Cliente</th>
              <th>Trabalho</th>
              <th>Data</th>
              <th>Valor</th>
              <th>Status</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="invoice in recentInvoices" :key="invoice.ID">
              <td>{{ invoice.ClientName }}</td>
              <td>{{ invoice.JobName }}</td>
              <td>{{ formatDate(invoice.IssueDate) }}</td>
              <td>{{ formatCurrency(invoice.TotalAmount) }}</td>
              <td>
                <span :class="['status-badge', invoice.Status.toLowerCase() === 'Confirmed' ? 'paid' : invoice.Status.toLowerCase() === 'pending' ? 'pending' : 'overdue']">
                  {{ invoice.Status }}
                </span>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  </template>
  <script setup>
  import { ref, onMounted, watch } from 'vue'
  import { getMyInvoices } from '@/services/invoiceService'
  
  // Refs principais
  const totalInvoices = ref(0)
  const totalAmount = ref(0)
  const pendingInvoices = ref(0)
  const recentInvoices = ref([])
  const loading = ref(true)
  
  const animatedStats = ref([
    { label: 'Total de Faturas', value: 0, display: 0 },
    { label: 'Valor Total', value: 0, display: 0 },
    { label: 'Pendentes', value: 0, display: 0 }
  ])
  
  const animateValue = (stat, target, formatFn = (v) => v, duration = 1000) => {
    const start = 0
    const startTime = performance.now()
  
    const step = (now) => {
      const progress = Math.min((now - startTime) / duration, 1)
      const current = start + (target - start) * progress
      stat.display = formatFn(current)
      if (progress < 1) requestAnimationFrame(step)
    }
  
    requestAnimationFrame(step)
  }
  
  const formatCurrency = (value) => {
    return new Intl.NumberFormat('pt-BR', {
      style: 'currency',
      currency: 'MZN',
    }).format(value)
  }
  
  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleDateString('pt-BR')
  }
  
  onMounted(async () => {
    loading.value = true
    try {
      const invoices = await getMyInvoices()
  
      totalInvoices.value = invoices.length
      totalAmount.value = invoices.reduce((sum, inv) => sum + inv.TotalAmount, 0)
      pendingInvoices.value = invoices.filter(inv => inv.Status === 'Pending').length
      recentInvoices.value = invoices.slice(0, 5)
  

    } catch (error) {
      console.error('Erro ao carregar faturas:', error)
    } finally {
      loading.value = false
    }
  })
  
  watch(loading, (val) => {
    if (!val) {
      animatedStats.value[0].value = totalInvoices.value
      animatedStats.value[1].value = totalAmount.value
      animatedStats.value[2].value = pendingInvoices.value
  
      animateValue(animatedStats.value[0], totalInvoices.value, (v) => Math.round(v))
      animateValue(animatedStats.value[1], totalAmount.value, formatCurrency)
      animateValue(animatedStats.value[2], pendingInvoices.value, (v) => Math.round(v))
    }
  })
  </script>
  
  
  <style scoped>
  .dashboard-content {
    padding: 2rem;
    display: flex;
    flex-direction: column;
    gap: 2rem;
  }
  
  /* Cards */
  .stats-container {
    display: flex;
    flex-wrap: wrap;
    gap: 1.5rem;
  }
  
  .stat-card {
    flex: 1 1 200px;
    padding: 1.5rem;
    border-radius: 1rem;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.08);
    background-color: #fff;
    transition: all 0.3s ease;
  }
  
  .stat-card h3 {
    font-size: 1rem;
    color: #555;
    margin-bottom: 0.5rem;
  }
  
  .stat-value {
    font-size: 1.5rem;
    font-weight: bold;
    color: #333;
  }
  
  /* Tabela */
  .recent-invoices {
    background: #fff;
    padding: 1.5rem;
    border-radius: 1rem;
    box-shadow: 0 4px 10px rgba(0, 0, 0, 0.08);
  }
  
  .recent-invoices h2 {
    margin-bottom: 1rem;
    font-size: 1.2rem;
    color: #333;
  }
  
  table {
    width: 100%;
    border-collapse: collapse;
    font-size: 0.95rem;
  }
  
  th,
  td {
    padding: 0.75rem 1rem;
    text-align: left;
    border-bottom: 1px solid #eee;
  }
  
  .status-badge {
    padding: 0.25rem 0.5rem;
    border-radius: 0.5rem;
    font-size: 0.85rem;
    color: #fff;
    text-transform: capitalize;
  }
  
  .status-badge.pending {
    background-color: #f59e0b;
  }
  
  .status-badge.paid {
    background-color: #10b981;
  }
  
  .status-badge.overdue {
    background-color: #ef4444;
  }
  
  /* Skeleton loading */
  .skeleton {
    background: linear-gradient(-90deg, #f0f0f0 0%, #e4e4e4 50%, #f0f0f0 100%);
    background-size: 200% 100%;
    animation: shimmer 1.2s infinite;
    border-radius: 1rem;
    height: 80px;
  }
  
  .table-skeleton {
    display: flex;
    flex-direction: column;
    gap: 1rem;
  }
  
  .skeleton-row {
    height: 24px;
    width: 100%;
    border-radius: 4px;
    background: linear-gradient(-90deg, #f0f0f0 0%, #e4e4e4 50%, #f0f0f0 100%);
    background-size: 200% 100%;
    animation: shimmer 1.2s infinite;
  }
  
  /* Fade-in da tabela */
  .fade-in {
    animation: fadeIn 0.6s ease-in-out;
  }
  
  @keyframes shimmer {
    0% {
      background-position: 200% 0;
    }
    100% {
      background-position: -200% 0;
    }
  }
  
  @keyframes fadeIn {
    0% {
      opacity: 0;
      transform: translateY(8px);
    }
    100% {
      opacity: 1;
      transform: translateY(0);
    }
  }
  </style>
  