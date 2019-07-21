package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"
	"github.com/gorilla/mux"
)

func (s Serv) updateCustomer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		msg := fmt.Errorf("dados inválidos (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	customer, err := s.read.GetCustomerByID(id)
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &customer)
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	err = s.update.UpdateCustomer(customer)
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, "cliente modificado com sucesso")

}

func (s Serv) updateUser(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	user, err := s.read.GetUserByID(id)
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &user)
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	err = s.update.UpdateUser(user)
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, "usuário modificado com sucesso")
}

func (s Serv) updateWarning(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	warning, err := s.read.GetWarningByID(id)
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(b, &warning)
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	err = s.update.UpdateWarning(warning)
	if err != nil {
		msg := fmt.Errorf("erro na solicitação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, "aviso modificado com sucesso")
}
