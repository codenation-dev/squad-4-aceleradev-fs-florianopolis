package rest

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
)

func readAllPublicFunc(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		params := mux.Vars(r)

		// uf := r.FormValue("uf")
		// year := r.FormValue("year")
		// month := r.FormValue("month")

		publicFuncs, err := reader.GetAllPublicFunc(params["uf"], params["year"], params["month"])
		if err != nil {
			http.Error(w, fmt.Sprintf("erro ao ler todos os dados (%s)", err.Error()), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, publicFuncs[:10]) // LIMITADO A 10 PARA TESTES
	}

}
