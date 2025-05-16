package router

import (
	"billbridge/controllers"
	"billbridge/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	// Inicializa o controller de faturas
	invoiceController := controllers.NewInvoiceController()


	// Rotas de faturas
	r.HandleFunc("/invoices", invoiceController.CreateInvoice).Methods("POST")
	r.HandleFunc("/invoices", invoiceController.GetAllInvoices).Methods("GET")
	r.Handle("/invoices/me", middleware.JWTMiddleware()(http.HandlerFunc(invoiceController.GetMyInvoices))).Methods("GET")
	r.HandleFunc("/invoices/{id}", invoiceController.GetInvoice).Methods("GET")
	r.HandleFunc("/invoices/{id}", invoiceController.UpdateInvoice).Methods("PUT")
	r.HandleFunc("/invoices/{id}", invoiceController.DeleteInvoice).Methods("DELETE")
	r.HandleFunc("/invoices/{id}/pdf", invoiceController.GetInvoicePDF).Methods("GET")


	// r.HandleFunc("/invoices/{id}/email", invoiceController.EmailInvoice).Methods("POST")

	return r
}
