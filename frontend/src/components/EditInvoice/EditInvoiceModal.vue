  <template>
      <div class="modal-overlay" @click.self="close">
        <div class="modal-content">
          <h2 class="modal-title">Editar Fatura</h2>
          <form @submit.prevent="submit" class="form">
            <!-- <pre>{{ form }}</pre> -->
            <!-- Nome do Trabalho -->
            <div class="form-group">
              <label for="jobName">Nome do Trabalho</label>
              <input id="jobName" v-model="form.jobName" required />
            </div>
    
            <!-- Cliente -->
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
    
            <!-- Valor Total -->
            <div class="form-group">
              <label for="totalAmount">Valor Total</label>
              <input id="totalAmount" v-model="form.totalAmount" type="number" step="0.01" required />
            </div>
    
            <!-- Valor da Taxa -->
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
              <button type="submit" class="save-btn">Atualizar</button>
              <button type="button" class="cancel-btn" @click="close">Cancelar</button>
            </div>
          </form>
        </div>
      </div>

    </template>
    
    <script setup>
    import { reactive, ref, watch, onMounted, defineProps, defineEmits, nextTick } from 'vue'
    import { updateInvoice, getAllTaxRules } from '@/services/invoiceService'
    
    const props = defineProps({
      invoice: Object,
      show: Boolean,
    })
    console.log('Props recebidas:', props.invoice)

    const emit = defineEmits(['close', 'updated'])
    const taxRules = ref([])
    
    const form = reactive({
    ID: null,
    jobName: '',
    clientName: '',
    providerName: '',
    country: '',
    region: '',
    issueDate: '',
    dueDate: '',
    totalAmount: 0,
    taxAmount: 0,
    taxRate: 0,
    currency: '',
    status: ''
  })


    
    watch(() => props.invoice, async (newVal) => {
      console.log('Nova invoice recebida no modal:', newVal)

    if (newVal) {

      Object.assign(form, { ...newVal,   issueDate: newVal.issueDate?.substring(0, 10),
        dueDate: newVal.dueDate?.substring(0, 10) })
        console.log('Form populado:', JSON.parse(JSON.stringify(form)))
        await nextTick()
      handleCountryChange() // importante se país estiver predefinido
      populateForm(newVal)
    }
  })
  function populateForm(invoice) {
  Object.assign(form, {
    ID: invoice.ID,
    jobName: invoice.JobName,
    clientName: invoice.ClientName,
    providerName: invoice.ProviderName,
    country: invoice.Country,
    region: invoice.Region,
    issueDate: invoice.IssueDate?.substring(0, 10),
    dueDate: invoice.DueDate?.substring(0, 10),
    totalAmount: invoice.TotalAmount,
    taxAmount: invoice.TaxAmount,
    taxRate: invoice.TaxRate,
    currency: invoice.Currency,
    status: invoice.Status
  })
}

    
    watch(() => form.totalAmount, () => {
      if (form.totalAmount && form.taxRate) {
        form.taxAmount = (parseFloat(form.totalAmount) * parseFloat(form.taxRate)) / 100
      }
    })

    watch(() => props.show, (visible) => {
    if (visible) {
      console.log('Modal aberto')
      console.log('Invoice recebida:', props.invoice)
    }
  })

    
  onMounted(async () => {
    try {
      taxRules.value = await getAllTaxRules()

      // Se já tiver invoice, chame aqui também
      if (props.invoice) {
        Object.assign(form, {
          ...props.invoice,
          issueDate: props.invoice.issueDate?.substring(0, 10),
          dueDate: props.invoice.dueDate?.substring(0, 10),
        })
        handleCountryChange()
        populateForm(props.invoice)
      }
    } catch (error) {
      console.error('Erro ao carregar regras fiscais:', error)
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
      const updated = {
        ...form,
        taxRate: parseFloat(form.taxRate),
        totalAmount: parseFloat(form.totalAmount),
        taxAmount: parseFloat(form.taxAmount),
        issueDate: new Date(form.issueDate).toISOString(),
        dueDate: new Date(form.dueDate).toISOString(),
      }
    
      updateInvoice(updated.ID, updated)
        .then(() => {
          emit('updated', updated)
          close()
        })
        .catch((err) => {
          console.error('Erro ao atualizar fatura:', err)
        })
    }
    
    function close() {
      emit('close')
    }
    </script>
    
    <style scoped>
    /* Reutiliza os mesmos estilos do modal de criação */
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
    textarea,
    select {
      padding: 0.5rem 0.75rem;
      border-radius: 0.5rem;
      border: 1px solid #cbd5e1;
      font-size: 1rem;
      background-color: #f9fafb;
      transition: border-color 0.2s;
    }
    
    input:focus,
    textarea:focus,
    select:focus {
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
    }
    
    .cancel-btn:hover {
      background-color: #f3f4f6;
    }
    </style>
    