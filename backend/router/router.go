package router

import (
	"billbridge/controller"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/invoices/generate/{jobId}", controller.GenerateInvoice).Methods("POST")
	r.HandleFunc("/invoices/{id}", controller.GetInvoice).Methods("GET")
	r.HandleFunc("/invoices/user/{userId}", controller.ListInvoicesByUser).Methods("GET")
	return r
}
