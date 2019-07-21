package rest

import (
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/adding"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
)

func getCustomer(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		customers, err := reader.GetCustomer(r.Form)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}
		respondWithJSON(w, http.StatusOK, customers)
	}
}

func importCustomer(adder adding.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		err := adder.ImportCustomer()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}
		respondWithJSON(w, http.StatusOK, "clientes importados com sucesso")
	}
}
