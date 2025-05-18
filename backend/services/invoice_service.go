package services

import (
	"billbridge/database"
	"billbridge/models"
	"errors"
	"time"
)

type InvoiceService struct{}

// NewInvoiceService cria uma nova instância do serviço de faturas
func NewInvoiceService() *InvoiceService {
	return &InvoiceService{}
}

// Create cria uma nova fatura
func (s *InvoiceService) Create(invoice *models.Invoice) error {
	if invoice.IssueDate.IsZero() {
		invoice.IssueDate = time.Now()
	}

	if invoice.Status == "" {
		invoice.Status = "Pending"
	}

	// Aplicar a regra de imposto
	if err := s.ApplyTaxRule(invoice); err != nil {
		return err
	}

	// Salvar o invoice no banco de dados
	result := database.DB.Create(invoice)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// GetByID retorna uma fatura pelo ID
func (s *InvoiceService) GetByID(id uint) (*models.Invoice, error) {
	var invoice models.Invoice
	result := database.DB.First(&invoice, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &invoice, nil
}

// GetAll retorna todas as faturas
func (s *InvoiceService) GetAll() ([]models.Invoice, error) {
	var invoices []models.Invoice
	result := database.DB.Find(&invoices)
	if result.Error != nil {
		return nil, result.Error
	}
	return invoices, nil
}

// Update atualiza uma fatura existente
func (s *InvoiceService) Update(invoice *models.Invoice) error {
	if invoice.ID == 0 {
		return errors.New("invoice ID is required for update")
	}

	result := database.DB.Save(invoice)
	return result.Error
}

// Delete remove uma fatura
func (s *InvoiceService) Delete(id uint) error {
	result := database.DB.Delete(&models.Invoice{}, id)
	return result.Error
}

// ApplyTaxRule aplica a regra de imposto a uma fatura
func (s *InvoiceService) ApplyTaxRule(invoice *models.Invoice) error {
	var taxRule models.TaxRule
	result := database.DB.Where("country = ? AND region = ? AND active = ?", invoice.Country, invoice.Region, true).First(&taxRule)
	if result.Error != nil {
		return errors.New("no active tax rule found for the specified country and region")
	}

	// Aplicar a taxa de imposto ao invoice
	invoice.TaxRate = taxRule.Rate
	invoice.TaxAmount = invoice.TotalAmount * (taxRule.Rate/100)
	return nil
}
func (s *InvoiceService) GeneratePDFBytes(invoice *models.Invoice) ([]byte, error) {
    pdfService := NewPDFService()
    return pdfService.GeneratePDFBytes(invoice)
}

func (s *InvoiceService) GetByProviderName(name string) ([]models.Invoice, error) {
	var invoices []models.Invoice
	if err := database.DB.Where("provider_name = ?", name).Find(&invoices).Error; err != nil {
		return nil, err
	}
	return invoices, nil
}

func (s *InvoiceService) GetTaxRules() ([]models.TaxRule, error) {
	var taxRules []models.TaxRule
	result := database.DB.Where("active = ?", true).Find(&taxRules)
	return taxRules, result.Error
}
