import axios from "axios";
import { getToken } from "./authService";
// import { getToken } from "./authService";


const API_URL =  "http://localhost:8180";

/**
 * @typedef {Object} Invoice
 * @property {number} [id]
 * @property {string} jobName
 * @property {string} providerName
 * @property {string} clientName
 * @property {string} country
 * @property {string} region
 * @property {string} [issueDate]
 * @property {string} dueDate
 * @property {number} totalAmount
 * @property {number} [taxAmount]
 * @property {number} [taxRate]
 * @property {string} currency
 * @property {string} [status]
 */

/**
 * Criar nova fatura
 * @param {Invoice} invoice
 * @returns {Promise<Invoice>}
 */
export const createInvoice = async (invoice) => {
  
  const res = await axios.post(`${API_URL}/invoices`, invoice);
  return res.data;
};

/**
 * Obter uma fatura por ID
 * @param {number} id
 * @returns {Promise<Invoice>}
 */
export const getInvoice = async (id) => {
  const res = await axios.get(`${API_URL}/invoices/${id}`);
  return res.data;
};

/**
 * Atualizar uma fatura
 * @param {number} id
 * @param {Invoice} invoice
 * @returns {Promise<Invoice>}
 */
export const updateInvoice = async (id, invoice) => {
  const res = await axios.put(`${API_URL}/invoices/${id}`, invoice);
  return res.data;
};

/**
 * Deletar uma fatura
 * @param {number} id
 * @returns {Promise<void>}
 */
export const deleteInvoice = async (id) => {
  await axios.delete(`${API_URL}/invoices/${id}`);
};

/**
 * Buscar faturas do provider logado
 * @returns {Promise<Invoice[]>}
 */
export const getMyInvoices = async () => {
  const res = await axios.get(`${API_URL}/invoices/me`, {
    headers: {
      Authorization: `Bearer ${getToken()}`
    }
  });
  return res.data;
};

/**
 * Gerar e baixar o PDF da fatura
 * @param {number} id
 * @returns {Promise<void>}
 */
export const downloadInvoicePDF = async (id) => {
  const res = await axios.get(`${API_URL}/invoices/${id}/pdf`, {
    responseType: "blob",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`
    }
  });

  const blob = new Blob([res.data], { type: "application/pdf" });
  const link = document.createElement("a");
  link.href = window.URL.createObjectURL(blob);
  link.download = `invoice_${id}.pdf`;
  document.body.appendChild(link);
  link.click();
  document.body.removeChild(link);
};

/**
 * @typedef {Object} TaxRule
 * @property {string} country
 * @property {string} region
 * @property {string} currency
 * @property {number} rate
 */

/**
 * Buscar todas as regras de taxa ativas
 * @returns {Promise<TaxRule[]>}
 */
export const getAllTaxRules = async () => {
  const res = await axios.get(`${API_URL}/taxes`);
  return res.data;
};
