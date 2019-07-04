package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"
	"github.com/gorilla/mux"
)

func (s Serv) deleteCustomerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		msg := fmt.Errorf("dados inválidos (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	ra, err := s.del.DeleteCustomerByID(id)
	if err != nil || ra == 0 {
		msg := fmt.Errorf("erro: nada foi deletado (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, fmt.Sprintf("Deletado %v cliente", ra))
}

func (s Serv) deleteWarningByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		msg := fmt.Errorf("dados inválidos (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	ra, err := s.del.DeleteWarningByID(id)
	if err != nil || ra == 0 {
		msg := fmt.Errorf("erro: nada foi deletado (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, fmt.Sprintf("Deletado %v aviso", ra))
}

func (s Serv) deletePublicByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		msg := fmt.Errorf("dados inválidos (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	ra, err := s.del.DeletePublicByID(id)
	if err != nil || ra == 0 {
		msg := fmt.Errorf("erro: nada foi deletado (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, fmt.Sprintf("Deletado %v aviso", ra))
}

func (s Serv) deleteUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		msg := fmt.Errorf("dados inválidos (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	ra, err := s.del.DeleteUserByID(id)
	if err != nil || ra == 0 {
		msg := fmt.Errorf("erro: nada foi deletado (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, fmt.Sprintf("Deletado %v usuário com sucesso", ra))
}
