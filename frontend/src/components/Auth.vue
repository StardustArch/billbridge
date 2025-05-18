<template>
  <div class="redirect-page">
    <div class="redirect-card">
      <!-- <img src="../assets/billbridge_logo.png" alt="Logo BillBridge" class="logo" /> -->

      <h1>Bem-vindo ao BillBridge</h1>

      <p>
        Para começar a usar o BillBridge, faça login pela sua conta do <strong>WorkBridge</strong>.
      </p>

      <p class="small-info">
        Esse acesso integrado garante uma experiência unificada, segura e simples para gerenciar suas faturas e análises.
      </p>

      <button @click="redirectToWorkBridge" class="redirect-button">
        Acessar com WorkBridge
      </button>
    </div>
  </div>
</template>
<script>
export default {
  name: 'RedirectPage',
}
</script>

<script setup>
function redirectToWorkBridge() {
  window.open('http://localhost:4200', '_blank') // ou '_self' para mesma aba
}
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  box-sizing: border-box;

}
.redirect-page {
  min-height: 100dvh;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 2rem;
  /* min-height: 100vh;   */
  overflow: hidden;
  scroll-behavior: unset;
}

.redirect-card {
  background: #ffffff;
  border-radius: 1.5rem;
  box-shadow: 0 8px 24px rgba(0, 0, 0, 0.08);
  padding: 3rem;
  max-width: 500px;
  width: 100%;
  text-align: center;
  animation: fadeIn 0.5s ease;
}

.logo {
  width: 200px;
  margin-bottom: 1.5rem;
}

h1 {
  font-size: 1.75rem;
  color: #111827;
  margin-bottom: 1rem;
}

p {
  font-size: 1rem;
  color: #4b5563;
  margin-bottom: 1rem;
}

.small-info {
  font-size: 0.875rem;
  color: #6b7280;
  margin-bottom: 2rem;
}

.redirect-button {
  background-color: #2563eb;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  font-size: 1rem;
  border-radius: 0.75rem;
  cursor: pointer;
  transition: background 0.3s ease;
}

.redirect-button:hover {
  background-color: #1e40af;
}

@keyframes fadeIn {
  from {
    opacity: 0;
    transform: translateY(12px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}
</style>
