package rest

import (
	"fmt"
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
	"github.com/gorilla/mux"
)

func getAllCustomer(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		customers, err := reader.GetAllCustomers(params["company"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintln(w, customers)
	}
}
