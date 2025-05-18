<template>
    <div class="analytics-dashboard">
      <h2 class="page-title">Análises de Vendas</h2>
  
      <div class="charts-grid">
        <!-- Gráfico de Vendas Mensais -->
        <div class="chart-card">
          <h3>Vendas Mensais</h3>
          <canvas id="monthlySalesChart"></canvas>
        </div>
  
        <!-- Gráfico de Serviços Vendidos -->
        <div class="chart-card">
          <h3>Serviços Vendidos</h3>
          <canvas id="servicesChart"></canvas>
        </div>
  
        <!-- Gráfico de Status das Faturas -->
        <div class="chart-card">
          <h3>Status das Faturas</h3>
          <canvas id="invoiceStatusChart"></canvas>
        </div>
      </div>
    </div>
  </template>
  <script>
export default {
    name: "AnalyticsDashboard"
  };
</script>
<script setup>
import { onMounted } from 'vue'
import Chart from 'chart.js/auto'
import { getMyInvoices } from '@/services/invoiceService'

onMounted(async () => {
  const invoices = await getMyInvoices()

  // 1. Vendas Mensais
  const monthlyTotals = Array(12).fill(0)
  invoices.forEach(inv => {
    const month = new Date(inv.IssueDate).getMonth() // 0-11
    monthlyTotals[month] += inv.TotalAmount
  })
  const monthLabels = ['Jan', 'Fev', 'Mar', 'Abr', 'Mai', 'Jun', 'Jul', 'Ago', 'Set', 'Out', 'Nov', 'Dez']
  new Chart(document.getElementById("monthlySalesChart"), {
    type: "line",
    data: {
      labels: monthLabels,
      datasets: [{
        label: "Vendas (MZN)",
        data: monthlyTotals,
        fill: false,
        borderColor: "#1e88e5",
        tension: 0.4,
      }]
    }
  })

  // 2. Serviços Vendidos
  const serviceCount = {}
invoices.forEach(inv => {
  if (inv.Status === "Confirmed") {
    serviceCount[inv.JobName] = (serviceCount[inv.JobName] || 0) + 1
  }
})
  const serviceLabels = Object.keys(serviceCount)
  const serviceValues = Object.values(serviceCount)
  new Chart(document.getElementById("servicesChart"), {
    type: "bar",
    data: {
      labels: serviceLabels,
      datasets: [{
        label: "Qtd. Vendida",
        data: serviceValues,
        backgroundColor: "#10b981",
      }]
    }
  })

  // 3. Status das Faturas
  const statusCount = {}
  invoices.forEach(inv => {
    statusCount[inv.Status] = (statusCount[inv.Status] || 0) + 1
  })
  const statusLabels = Object.keys(statusCount)
  const statusValues = Object.values(statusCount)
  new Chart(document.getElementById("invoiceStatusChart"), {
    type: "doughnut",
    data: {
      labels: statusLabels,
      datasets: [{
        data: statusValues,
        backgroundColor: ["#10b981", "#f59e0b", "#ef4444", "#3b82f6"],
      }]
    }
  })
})
</script>

  
  <style scoped>
  .analytics-dashboard {
    padding: 2rem;
  }
  
  .page-title {
    font-size: 1.8rem;
    margin-bottom: 2rem;
    font-weight: bold;
    color: #1a237e;
  }
  
  .charts-grid {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}
  

  .chart-card {
    border-radius: 1rem;
    box-shadow: 0 0 12px rgba(0, 0, 0, 0.05);
    background: white;
  width: 70%;
  max-width: 900px;
  margin: 0 auto;
  margin-bottom: 7rem;
  /* height: 300px; */
}

  .chart-card h3 {
    font-size: 1.1rem;
    margin-bottom: 1rem;
    color: #333;
  }
  </style>
  