package rest

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gbletsch/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
)

func getAllCustomer(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "list_customer.html", nil)
	}
}

func processGetAllCustomer(tpl *template.Template, reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		company := r.FormValue("company")
		customers, err := reader.GetAllCustomers(company)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, customers)
	}
}
