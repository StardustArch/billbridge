<template>
  <div class="invoice-list">
    <!-- Filtros e Ações -->
    <div class="actions-bar">
      <div class="filters">
        <input
          v-model="filters.search"
          type="text"
          placeholder="Buscar por cliente ou trabalho..."
          class="search-input"
        />
        <select v-model="filters.status" class="status-filter">
          <option value="">Todos os Status</option>
          <option value="pending">Pendente</option>
          <option value="confirmed">Pago</option>
          <option value="cancelled">Cancelado</option>
        </select>
      </div>
      <button
        class="new-invoice-btn"
        @click="$router.push('/dashboard/new-invoice')"
      >
        <i class="fas fa-plus"></i> Nova Fatura
      </button>
    </div>

    <!-- Tabela de Faturas -->
    <div class="table-container">
      <table v-if="!loading && filteredInvoices.length">
        <thead>
          <tr>
            <th @click="sort('clientName')">
              Cliente

            </th>
            <th @click="sort('jobName')">
              Trabalho
              
            </th>
            <th @click="sort('issueDate')">
              Data
              
            </th>
            <th @click="sort('totalAmount')">
              Valor
              
            </th>
            <th @click="sort('status')">
              Status
              
            </th>
            <th>Ações</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="invoice in filteredInvoices" :key="invoice.ID">
            <td>{{ invoice.ClientName }}</td>
            <td>{{ invoice.JobName }}</td>
            <td>{{ formatDate(invoice.IssueDate) }}</td>
            <td>{{ formatCurrency(invoice.TotalAmount) }}</td>
            <td>
              <span :class="['status-badge', invoice.Status.toLowerCase() === 'Confirmed' ? 'paid' : invoice.Status.toLowerCase() === 'pending' ? 'pending' : 'overdue']">
                {{ invoice.Status }}
              </span>
            </td>
            <td class="actions">
              <button @click="editInvoiceModal(invoice)" class="action-btn edit">
                <i class="fas fa-edit"></i>
              </button>
              <button
                @click="showDeleteConfirm(invoice)"
                class="action-btn delete"
              >
                <i class="fas fa-trash"></i>
              </button>
              <button
                @click="generatePdf(invoice)"
                class="action-btn pdf"
              >
                <i class="fas fa-file-pdf"></i>
              </button>
            </td>
          </tr>
        </tbody>
      </table>

      <div v-else-if="loading" class="loading-skeleton">
        <div v-for="n in 5" :key="n" class="skeleton-row"></div>
      </div>

      <div v-else class="no-results">
        <i class="fas fa-file-invoice"></i>
        <p>Nenhuma fatura encontrada</p>
      </div>
    </div>

    <!-- Modal de Confirmação de Deleção -->
    <div
      v-if="showDeleteModal"
      class="modal-overlay"
      @click="showDeleteModal = false"
    >
      <div class="modal" @click.stop>
        <h3>Confirmar Exclusão</h3>
        <p>Tem certeza que deseja excluir esta fatura?</p>
        <div class="modal-actions">
          <button @click="showDeleteModal = false" class="cancel-btn">
            Cancelar
          </button>
          <button @click="deleteInvoice" class="delete-btn">Excluir</button>
        </div>
      </div>
    </div>
  </div>
  <EditInvoiceModal
  v-if="showEditModal"
  :show="showEditModal"
  :invoice="editingInvoice"
  @close="showEditModal = false"
  @updated="handleInvoiceUpdated"
/>

</template>

<script>
export default {
  name: "InvoiceList",
};
</script>

<script setup>
import { ref, computed, onMounted } from "vue";

import { getMyInvoices, deleteInvoice as deleteInvoiceById, updateInvoice, downloadInvoicePDF } from "@/services/invoiceService";
import EditInvoiceModal from '@/components/EditInvoice/EditInvoiceModal.vue'
import { nextTick } from 'vue'

const showEditModal = ref(false)
const editingInvoice = ref(null)


const invoices = ref([]);
const loading = ref(true);
const showDeleteModal = ref(false);
const selectedInvoice = ref(null);
const sortConfig = ref({ field: "issueDate", ascending: false });

const editInvoiceModal = async (invoice) => {
  editingInvoice.value = { ...invoice }
  await nextTick()
  showEditModal.value = true
}
const handleInvoiceUpdated = async (updated) => {
  // Cria um objeto personalizado a partir do invoice atualizado
  const customInvoice = {
    ID: updated.ID,
    ClientName: updated.clientName,
    Country: updated.country,
    Currency: updated.currency,
    DueDate: updated.dueDate,
    IssueDate: updated.issueDate,
    JobName: updated.jobName,
    ProviderName: updated.providerName,
    Region: updated.region,
    Status: updated.status,
    TaxAmount: updated.taxAmount,
    TotalAmount: updated.totalAmount,
    // aqui pode montar só os campos que quiser mandar para o backend
  };

  try {
    // Chama o service para atualizar no backend
    const updatedFromServer = await updateInvoice(updated.ID, customInvoice);

    // Atualiza o invoice localmente na lista
    const index = invoices.value.findIndex(inv => inv.ID === updatedFromServer.ID);
    if (index !== -1) {
      invoices.value[index] = updatedFromServer;
    }
  } catch (error) {
    console.error('Erro ao atualizar invoice:', error);
  }
};

const generatePdf = (invoice) => {
  // Implementar lógica para gerar PDF da fatura
  downloadInvoicePDF(invoice.ID);
  // console.log("Gerar PDF para a fatura:", invoice);
};


const filters = ref({
  search: "",
  status: "",
});

// Formatadores
const formatCurrency = (value) => {
  return new Intl.NumberFormat("pt-BR", {
    style: "currency",
    currency: "MZN",
  }).format(value);
};

const formatDate = (dateString) => {
  return new Date(dateString).toLocaleDateString("pt-BR");
};

// Carregar faturas
const loadInvoices = async () => {
  loading.value = true;
  try {
    const data = await getMyInvoices();
    invoices.value = data;
  } catch (error) {
    console.error("Erro ao carregar faturas:", error);
  } finally {
    loading.value = false;
  }
};


// Filtrar e ordenar faturas
const filteredInvoices = computed(() => {
  let filtered = [...invoices.value];

  // Aplicar filtro de busca
  if (filters.value.search) {
    const searchTerm = filters.value.search.toLowerCase();
    filtered = filtered.filter(
      (invoice) =>
        invoice.ClientName.toLowerCase().includes(searchTerm) ||
        invoice.JobName.toLowerCase().includes(searchTerm)
    );
  }

  // Aplicar filtro de status
  if (filters.value.status) {
    filtered = filtered.filter(
      (invoice) => invoice.Status.toLowerCase() === filters.value.status
    );
  }

  // Aplicar ordenação
  filtered.sort((a, b) => {
  let aVal = a[sortConfig.value.field];
  let bVal = b[sortConfig.value.field];

  if (typeof aVal === "string") {
    aVal = aVal.toLowerCase();
  } else if (aVal === undefined || aVal === null) {
    aVal = "";
  }

  if (typeof bVal === "string") {
    bVal = bVal.toLowerCase();
  } else if (bVal === undefined || bVal === null) {
    bVal = "";
  }

  if (aVal < bVal) return sortConfig.value.ascending ? -1 : 1;
  if (aVal > bVal) return sortConfig.value.ascending ? 1 : -1;
  return 0;
});


  return filtered;
});

// Ordenação
const sort = (field) => {
  if (sortConfig.value.field === field) {
    sortConfig.value.ascending = !sortConfig.value.ascending;
  } else {
    sortConfig.value.field = field;
    sortConfig.value.ascending = true;
  }
};

// Ações de Deleção
const showDeleteConfirm = (invoice) => {
  selectedInvoice.value = invoice;
  showDeleteModal.value = true;
};

const deleteInvoice = async () => {
  if (!selectedInvoice.value) return;

  try {
    await deleteInvoiceById(selectedInvoice.value.ID);
    invoices.value = invoices.value.filter(
      (inv) => inv.ID !== selectedInvoice.value.ID
    );
    showDeleteModal.value = false;
  } catch (error) {
    console.error("Erro ao deletar fatura:", error);
  }
};

onMounted(() => {
  loadInvoices();
});
</script>

<style scoped>
.invoice-list {
  padding: 2rem;
}

.actions-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.filters {
  display: flex;
  gap: 1rem;
}

.search-input,
.status-filter {
  padding: 0.5rem 1rem;
  border: 1px solid #ddd;
  border-radius: 0.5rem;
  font-size: 0.9rem;
}

.new-invoice-btn {
  padding: 0.5rem 1rem;
  background: #1a237e;
  color: white;
  border: none;
  border-radius: 0.5rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

.table-container {
  background: white;
  border-radius: 1rem;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.08);
  overflow: hidden;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th {
  background: #f8f9fa;
  cursor: pointer;
}

th,
td {
  padding: 1rem;
  text-align: left;
  border-bottom: 1px solid #eee;
}

.actions {
  display: flex;
  gap: 0.5rem;
}

.action-btn {
  padding: 0.25rem 0.5rem;
  border: none;
  border-radius: 0.25rem;
  cursor: pointer;
}

.action-btn.edit {
  background: #4caf50;
  color: white;
}

.action-btn.delete {
  background: #f44336;
  color: white;
}

/* Status badges */
.status-badge {
  padding: 0.25rem 0.5rem;
  border-radius: 0.5rem;
  font-size: 0.85rem;
  color: white;
}

.status-badge.pending {
  background: #f59e0b;
}
.status-badge.paid {
  background: #10b981;
}
.status-badge.overdue {
  background: #ef4444;
}

/* Modal */
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal {
  background: white;
  padding: 2rem;
  border-radius: 0.5rem;
  width: 400px;
}

.modal-actions {
  display: flex;
  justify-content: flex-end;
  gap: 1rem;
  margin-top: 1.5rem;
}

/* Loading skeleton */
.loading-skeleton {
  padding: 1rem;
}

.skeleton-row {
  height: 40px;
  background: linear-gradient(-90deg, #f0f0f0 0%, #e4e4e4 50%, #f0f0f0 100%);
  background-size: 200% 100%;
  animation: shimmer 1.2s infinite;
  margin-bottom: 0.5rem;
  border-radius: 0.25rem;
}

@keyframes shimmer {
  0% {
    background-position: 200% 0;
  }
  100% {
    background-position: -200% 0;
  }
}

/* Empty state */
.no-results {
  padding: 3rem;
  text-align: center;
  color: #666;
}

.no-results i {
  font-size: 3rem;
  margin-bottom: 1rem;
}
</style>
