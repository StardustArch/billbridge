package controllers

import (
	"billbridge/models"
	"billbridge/services"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type InvoiceController struct {
	service *services.InvoiceService
}

// NewInvoiceController cria uma nova instância do controller de faturas
func NewInvoiceController() *InvoiceController {
	return &InvoiceController{
		service: services.NewInvoiceService(),
	}
}

// CreateInvoice cria uma nova fatura
func (c *InvoiceController) CreateInvoice(w http.ResponseWriter, r *http.Request) {
	var invoice models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := c.service.Create(&invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(invoice)
}

// GetInvoice retorna uma fatura específica
func (c *InvoiceController) GetInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}

	invoice, err := c.service.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Invoice not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

// // GetAllInvoices retorna todas as faturas
// func (c *InvoiceController) GetAllInvoices(w http.ResponseWriter, r *http.Request) {
// 	invoices, err := c.service.GetAll()
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(invoices)
// }

// UpdateInvoice atualiza uma fatura existente
func (c *InvoiceController) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}

	var invoice models.Invoice
	if err := json.NewDecoder(r.Body).Decode(&invoice); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	invoice.ID = uint(id)
	if err := c.service.Update(&invoice); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoice)
}

// DeleteInvoice remove uma fatura
func (c *InvoiceController) DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}

	if err := c.service.Delete(uint(id)); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}


// GetInvoicePDF gera e retorna o PDF da fatura
func (c *InvoiceController) GetInvoicePDF(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		http.Error(w, "Invalid invoice ID", http.StatusBadRequest)
		return
	}

	invoice, err := c.service.GetByID(uint(id))
	if err != nil {
		http.Error(w, "Invoice not found", http.StatusNotFound)
		return
	}

	pdfBytes, err := c.service.GeneratePDFBytes(invoice)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to generate PDF: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/pdf")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=invoice_%d.pdf", invoice.ID))
	w.WriteHeader(http.StatusOK)
	w.Write(pdfBytes)
}

// GetMyInvoices retorna todas as faturas do provider logado
func (c *InvoiceController) GetMyInvoices(w http.ResponseWriter, r *http.Request) {
	providerName := r.Context().Value("username")
	if providerName == nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	invoices, err := c.service.GetByProviderName(providerName.(string))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}
