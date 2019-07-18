package rest

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gbletsch/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
)

func compareCustomerPublicFunc(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "compare_customer_public_func.html", nil)
	}
}

func processCompareCustomerPublicFunc(tpl *template.Template, reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uf := r.FormValue("uf")
		month := r.FormValue("month")
		year := r.FormValue("year")
		company := r.FormValue("company")

		list, err := reader.CompareCustomerPublicFunc(uf, month, year, company)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprint(w, list)

	}
}
