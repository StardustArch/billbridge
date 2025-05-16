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
  import { onMounted } from "vue";
  import Chart from "chart.js/auto";
  
  onMounted(() => {
    // Gráfico de vendas mensais (linha)
    new Chart(document.getElementById("monthlySalesChart"), {
      type: "line",
      data: {
        labels: ["Jan", "Fev", "Mar", "Abr", "Mai", "Jun"],
        datasets: [{
          label: "Vendas (MZN)",
          data: [12000, 18000, 15000, 20000, 24000, 22000],
          fill: false,
          borderColor: "#1e88e5",
          tension: 0.4,
        }]
      }
    });
  
    // Gráfico de serviços vendidos (barra)
    new Chart(document.getElementById("servicesChart"), {
      type: "bar",
      data: {
        labels: ["Design", "Dev Web", "Marketing", "SEO", "Consultoria"],
        datasets: [{
          label: "Qtd. Vendida",
          data: [12, 19, 9, 14, 7],
          backgroundColor: "#10b981",
        }]
      }
    });
  
    // Gráfico de status das faturas (doughnut)
    new Chart(document.getElementById("invoiceStatusChart"), {
      type: "doughnut",
      data: {
        labels: ["Pagas", "Pendentes", "Atrasadas"],
        datasets: [{
          data: [60, 25, 15],
          backgroundColor: ["#10b981", "#f59e0b", "#ef4444"],
        }]
      }
    });
  });
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
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(300px, 1fr));
    gap: 2rem;
  }
  
  .chart-card {
    background: white;
    padding: 1.5rem;
    border-radius: 1rem;
    box-shadow: 0 0 12px rgba(0, 0, 0, 0.05);
  }
  
  .chart-card h3 {
    font-size: 1.1rem;
    margin-bottom: 1rem;
    color: #333;
  }
  </style>
  