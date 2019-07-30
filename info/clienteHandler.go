package api

import (
	"fmt"
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/database"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"
	"github.com/gorilla/schema"
)

//Lista os funcionários
func ClienteList(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		// Handle error
	}

	filter := new(database.ClienteFilter)
	if err := schema.NewDecoder().Decode(filter, r.Form); err != nil {
		// Handle error
	}

	if filter.Offset == 0 || filter.Offset > 50 {
		filter.Offset = 50
	}
	if filter.Page == 0 {
		filter.Page = 1
	}
	if filter.SortBy == "" {
		filter.SortBy = "nome"
	}
	// Do something with filter
	fmt.Printf("%+v", filter)

	clienteList, err := database.ClienteList(filter)
	if err != nil {
		model.ErrorResponse(w, fmt.Errorf("sem dados: %v", err), http.StatusBadRequest)
		return
	}
	model.ToJSON(w, clienteList)
}
