<template>
  <div class="modal-overlay" @click.self="close">
    <div class="modal-content">
      <h2 class="modal-title">Nova Fatura</h2>
      <form @submit.prevent="submit" class="form">
        <!-- Job Name -->
        <div class="form-group">
          <label for="jobName">Nome do Trabalho</label>
          <input id="jobName" v-model="form.jobName" required />
        </div>

        <!-- Client Name -->
        <div class="form-group">
          <label for="clientName">Cliente</label>
          <input id="clientName" v-model="form.clientName" required />
        </div>

        <!-- País -->
        <div class="form-group">
          <label for="country">País</label>
          <select id="country" v-model="form.country" @change="handleCountryChange" required>
            <option value="" disabled>Selecione um país</option>
            <option v-for="rule in taxRules" :key="rule.Country" :value="rule.Country">
              {{ rule.Country }}
            </option>
          </select>
        </div>

        <!-- Região -->
        <div class="form-group">
          <label for="region">Região</label>
          <input id="region" v-model="form.region" disabled />
        </div>

        <!-- Moeda -->
        <div class="form-group">
          <label for="currency">Moeda</label>
          <input id="currency" v-model="form.currency" disabled />
        </div>

        <!-- Taxa (%) -->
        <div class="form-group">
          <label for="taxRate">Taxa (%)</label>
          <input id="taxRate" v-model="form.taxRate" disabled />
        </div>

        <!-- Montante Total -->
        <div class="form-group">
          <label for="totalAmount">Valor Total</label>
          <input id="totalAmount" v-model="form.totalAmount" type="number" step="0.01" required />
        </div>

        <!-- Valor da Taxa (calculado) -->
        <div class="form-group">
          <label for="taxAmount">Valor da Taxa</label>
          <input id="taxAmount" v-model="form.taxAmount" type="number" step="0.01" disabled />
        </div>

        <!-- Data de Emissão -->
        <div class="form-group">
          <label for="issueDate">Data de Emissão</label>
          <input id="issueDate" v-model="form.issueDate" type="date" required />
        </div>

        <!-- Data de Vencimento -->
        <div class="form-group">
          <label for="dueDate">Data de Vencimento</label>
          <input id="dueDate" v-model="form.dueDate" type="date" required />
        </div>

        <!-- Status -->
        <div class="form-group">
          <label for="status">Status</label>
          <select id="status" v-model="form.status">
            <option value="Pending">Pendente</option>
            <option value="Confirmed">Pago</option>
            <option value="Cancelled">Cancelado</option>
          </select>
        </div>

        <!-- Botões -->
        <div class="form-actions">
          <button type="submit" class="save-btn">Salvar</button>
          <button type="button" class="cancel-btn" @click="close">Cancelar</button>
        </div>
      </form>
    </div>
  </div>
</template>


<script setup>
import { reactive, ref, onMounted, watch } from 'vue'
import { useRouter } from 'vue-router'
import { getAllTaxRules, createInvoice } from '@/services/invoiceService'
import { getUsername } from '@/services/authService'

const router = useRouter()
const taxRules = ref([])

const form = reactive({
  jobName: '',
  providerName: '',
  clientName: '',
  country: '',
  region: '',
  currency: '',
  taxRate: '',
  taxAmount: 0,
  totalAmount: '',
  issueDate: new Date().toISOString().slice(0, 10),
  dueDate: '',
  status: 'PENDING',
})

onMounted(async () => {
  try {
    const user = JSON.parse(getUsername())
    form.providerName = user?.username || ''
    taxRules.value = (await getAllTaxRules()).map(rule => ({
      ...rule,
      currency: getCurrencyForCountry(rule.Country)
    }))
  } catch (error) {
    console.error('Erro ao carregar dados iniciais:', error)
  }
})

watch(() => form.totalAmount, () => {
  if (form.totalAmount && form.taxRate) {
    form.taxAmount = (parseFloat(form.totalAmount) * parseFloat(form.taxRate)) / 100
  }
})

function handleCountryChange() {
  const rule = taxRules.value.find(r => r.Country === form.country)
  if (rule) {
    form.region = rule.Region
    form.taxRate = rule.Rate
    form.currency = getCurrencyForCountry(rule.Country)
    if (form.totalAmount) {
      form.taxAmount = (parseFloat(form.totalAmount) * rule.Rate) / 100
    }
  }
}

function getCurrencyForCountry(code) {
  const currencyMap = {
    MZ: 'MZN', ZA: 'ZAR', ZW: 'ZWL', ZM: 'ZMW', MW: 'MWK', TZ: 'TZS', SZ: 'SZL',
    DE: 'EUR', FR: 'EUR', GB: 'GBP', ES: 'EUR', IT: 'EUR', PT: 'EUR', NL: 'EUR',
    US: 'USD', CA: 'CAD', BR: 'BRL', MX: 'MXN', CO: 'COP'
  }
  return currencyMap[code] || ''
}

function submit() {
  const invoice = {
    jobName: form.jobName,
    providerName: form.providerName,
    clientName: form.clientName,
    country: form.country,
    region: form.region,
    currency: form.currency,
    taxRate: parseFloat(form.taxRate),
    taxAmount: form.taxAmount,
    totalAmount: parseFloat(form.totalAmount),
    issueDate: toISOStringDateOnly(form.issueDate),
    dueDate: toISOStringDateOnly(form.dueDate),
    status: form.status
  }
  // console.log('Invoice pronta para envio:', invoice)
  createInvoice(invoice)
    .then(() => {
      console.log('Fatura criada com sucesso!')
      router.push('/dashboard/invoices')
    })
    .catch(error => {
      console.error('Erro ao criar fatura:', error)
    })
}
function toISOStringDateOnly(dateStr) {
  return new Date(dateStr).toISOString(); // Resultado: "2024-05-17T00:00:00.000Z"
}


function close() {
  router.push('/dashboard/invoices')
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
  display: grid;
  grid-template-columns: 1fr;
  gap: 1rem;
}

@media (min-width: 640px) {
  .form {
    grid-template-columns: 1fr 1fr;
  }

  .form-actions {
    grid-column: span 2;
    justify-content: flex-end;
  }
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
  grid-column: span 2;
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
