package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"

	"github.com/gorilla/mux"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	model.ToJSON(w, struct {
		Message string `json:"message"`
	}{
		Message: "Banco Uati",
	})
}

func (s Serv) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := s.read.GetAllCustomers()
	if err != nil {
		model.ErrorResponse(w, fmt.Errorf("lista de clientes não retornada: %v", err), http.StatusBadRequest)
		return
	}
	model.ToJSON(w, customers)
}

func (s Serv) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.read.GetAllUsers()
	if err != nil {
		msg := fmt.Errorf("Erro lendo o banco de dados: %v", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, users)
}

func (s Serv) getAllWarnings(w http.ResponseWriter, r *http.Request) {
	warnings, err := s.read.GetAllWarnings()
	if err != nil {
		msg := fmt.Errorf("Erro lendo o banco de dados: %v", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, warnings)
}

// ByName
func (s Serv) getCustomerByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	customers, err := s.read.GetCustomerByName(params["pattern"])
	if err != nil {
		msg := fmt.Errorf("Houve um problema na procura deste cliente: %v", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, customers)
}

func (s Serv) getUserByEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user, err := s.read.GetUserByEmail(params["pattern"])
	if err != nil {
		msg := fmt.Errorf("Houve um problema na procura deste usuario: %v", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, user)
}

func (s *Serv) getWarningByCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	warnings, err := s.read.GetWarningByCustomer(params["pattern"])
	if err != nil {
		model.ErrorResponse(w, fmt.Errorf("sem dados: %v", err), http.StatusBadRequest)
		return
	}
	model.ToJSON(w, warnings)
}

func (s *Serv) getWarningByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	warning, err := s.read.GetWarningByUser(params["pattern"])
	if err != nil {
		model.ErrorResponse(w, fmt.Errorf("sem dados: %v", err), http.StatusBadRequest)
		return
	}
	model.ToJSON(w, warning)
}

func (s *Serv) getPublicByWage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pattern, err := strconv.ParseFloat(params["pattern"], 32)
	if err != nil {
		model.ErrorResponse(w, fmt.Errorf("dados inválidos: %v", err), http.StatusBadRequest)
		return
	}
	publicFuncs, err := s.read.GetPublicByWage(float32(pattern))
	if err != nil {
		model.ErrorResponse(w, fmt.Errorf("sem dados: %v", err), http.StatusBadRequest)
		return
	}
	model.ToJSON(w, publicFuncs)
}

//ByID
func (s Serv) getCustomerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		msg := fmt.Errorf("dados inválidos (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	c, err := s.read.GetCustomerByID(id)
	if err != nil {
		msg := fmt.Errorf("Houve um problema na procura deste cliente: %v", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, c)
}

func (s Serv) getUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		msg := fmt.Errorf("dados inválidos (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	user, err := s.read.GetUserByID(id)
	if err != nil {
		msg := fmt.Errorf("Houve um problema na procura deste usuário: %v", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, user)
}

func (s Serv) getWarningByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		msg := fmt.Errorf("dados inválidos (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	user, err := s.read.GetWarningByID(id)
	if err != nil {
		msg := fmt.Errorf("Houve um problema na procura desta mensagem: %v", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, user)
}

func (s Serv) getPublicByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		msg := fmt.Errorf("dados inválidos (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	user, err := s.read.GetPublicByID(id)
	if err != nil {
		msg := fmt.Errorf("Houve um problema na procura deste funcionário público: %v", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, user)
}
