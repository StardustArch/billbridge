<template>
  <div class="modal-overlay" @click.self="close">
    <div class="modal-content">
      <h2 class="modal-title">Nova Fatura</h2>
      <form @submit.prevent="submit" class="form">
        <div class="form-group">
          <label for="title">Título</label>
          <input id="title" v-model="form.title" required />
        </div>

        <div class="form-group">
          <label for="amount">Valor</label>
          <input id="amount" v-model="form.amount" type="number" step="0.01" required />
        </div>

        <div class="form-group">
          <label for="date">Data</label>
          <input id="date" v-model="form.date" type="date" required />
        </div>

        <div class="form-group">
          <label for="description">Descrição</label>
          <textarea id="description" v-model="form.description" rows="3" />
        </div>

        <div class="form-actions">
          <button type="submit" class="save-btn">
            <i class="fas fa-check"></i> Salvar
          </button>
          <button type="button" @click="close" class="cancel-btn">
            <i class="fas fa-times"></i> Cancelar
          </button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const form = reactive({
  title: '',
  amount: '',
  date: '',
  description: '',
})

function submit() {
  console.log('Fatura manual submetida:', form)
  router.push('/dashboard/invoices')
}

function close() {
  router.back()
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(15, 23, 42, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal-content {
  background-color: white;
  padding: 2rem;
  border-radius: 0.75rem;
  width: 100%;
  max-width: 500px;
  box-shadow: 0 10px 25px rgba(0, 0, 0, 0.15);
  font-family: 'Inter', sans-serif;
}

.modal-title {
  font-size: 1.25rem;
  font-weight: 600;
  margin-bottom: 1.5rem;
  color: #1f2937;
}

.form {
  display: flex;
  flex-direction: column;
  gap: 1rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

input,
textarea {
  padding: 0.5rem 0.75rem;
  border-radius: 0.5rem;
  border: 1px solid #cbd5e1;
  font-size: 1rem;
  background-color: #f9fafb;
  transition: border-color 0.2s;
}

input:focus,
textarea:focus {
  border-color: #2563eb;
  outline: none;
}

.form-actions {
  display: flex;
  justify-content: flex-end;
  gap: 0.75rem;
  margin-top: 1rem;
}

.save-btn {
  background-color: #2563eb;
  color: white;
  border: none;
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  font-weight: 500;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.save-btn:hover {
  background-color: #1e40af;
}

.cancel-btn {
  background: transparent;
  border: 1px solid #cbd5e1;
  color: #1f2937;
  padding: 0.5rem 1rem;
  border-radius: 0.5rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.cancel-btn:hover {
  background-color: #f3f4f6;
}
</style>
