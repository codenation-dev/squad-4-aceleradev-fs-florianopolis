package rest

import (
	"fmt"
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
	"github.com/gorilla/mux"
)

func compareCustomerPublicFunc(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		// uf := r.FormValue("uf")
		// month := r.FormValue("month")
		// year := r.FormValue("year")
		// company := r.FormValue("company")

		list, err := reader.CompareCustomerPublicFunc(
			params["uf"], params["month"], params["year"], params["company"],
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("erro aocomparar tabelas (%s)", err.Error()), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, list)

	}
}

func getPublicFincByWage(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		publicFuncs, err := reader.GetPublicFuncByWage(
			params["uf"], params["year"], params["month"], params["wage"],
		)
		if err != nil {
			http.Error(w, fmt.Sprintf("erro ao buscar dados (%s)", err.Error()), http.StatusBadRequest)
			return
		}
		fmt.Fprint(w, publicFuncs[:30]) // LIMITADO PARA TESTES
	}
}
