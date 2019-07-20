package api

import (
	"fmt"
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/database"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"
	"github.com/gorilla/schema"
)

//Lista os funcionários
func FuncionarioList(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		// Handle error
	}

	filter := new(database.FuncionarioFilter)
	if err := schema.NewDecoder().Decode(filter, r.Form); err != nil {
		// Handle error
	}

	if filter.Offset == 0 || filter.Offset > 50 {
		filter.Offset = 20
	}

	if filter.SortBy == "" {
		filter.SortBy = "nome"
	}
	// Do something with filter
	fmt.Printf("%+v", filter)

	funcionarioList, err := database.FuncionarioList(filter)
	fmt.Println("retorno", err)
	if err != nil {
		model.ErrorResponse(w, fmt.Errorf("sem dados: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println(funcionarioList)
	model.ToJSON(w, funcionarioList)
}

//Lista os funcionários
func FuncionarioGet(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		// Handle error
	}

	filter := new(database.FuncionarioFilter)
	if err := schema.NewDecoder().Decode(filter, r.Form); err != nil {
		// Handle error
	}

	if filter.ID == 0 && filter.Nome == "" {
		model.ErrorResponse(w, fmt.Errorf("Você precisa informar os dados do funcionário."), http.StatusBadRequest)
		return
	}

	// Do something with filter
	fmt.Printf("%+v", filter)

	funcionario, err := database.FuncionarioGet(filter)
	fmt.Println("retorno", err)
	if err != nil {
		model.ErrorResponse(w, fmt.Errorf("sem dados: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println(funcionario)
	model.ToJSON(w, funcionario)
}
