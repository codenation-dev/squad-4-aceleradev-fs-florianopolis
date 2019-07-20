package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/emailserv"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/importing"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/model"
)

func (s *Serv) sendEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["userEmail"]
	user, err := s.read.GetUserByEmail(email)
	if err != nil {
		msg := fmt.Errorf("usuário nao encontrado (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	publicFuncs := []model.PublicFunc{}
	err = json.NewDecoder(r.Body).Decode(&publicFuncs)
	if err != nil {
		msg := fmt.Errorf("corpo do email inválido (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	err = emailserv.Send(publicFuncs, user)
	if err != nil {
		msg := fmt.Errorf("erro no envio do email (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, "Email enviado com sucesso")
}

// ImportPublicFuncFile - Download e importação de arquivo de funcionário público do estado de SP
func (s Serv) ImportPublicFuncFile(w http.ResponseWriter, r *http.Request) {
	salarioList, err := importing.ImportPublicFunc()
	if err != nil {
		msg := fmt.Errorf("erro na importação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	err = s.add.AddPublicFunc(salarioList...)
	if err != nil {
		msg := fmt.Errorf("erro ao adicionar funcionário público (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, "Arquivo de funcionários públicos importado com sucesso")
}

// ImportCustomerFile importa o arquivo 'clientes.csv'
func (s Serv) ImportCustomerFile(w http.ResponseWriter, r *http.Request) {
	clienteList, err := importing.ImportClientesCSV("backend/cmd/data/clientes.csv")
	if err != nil {
		msg := fmt.Errorf("erro na importação (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}

	for _, cliente := range clienteList {
		err := s.add.AddCustomer(cliente)
		if err != nil {
			msg := fmt.Errorf("erro na importação (%v)", err)
			model.ErrorResponse(w, msg, http.StatusBadRequest)
			return
		}
	}
	model.ToJSON(w, "Arquivo de clientes importado com sucesso")
}

// AddCustomer adiciona clientes ao banco de dados
func (s Serv) AddCustomer(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	c := model.Cliente{}
	err = json.Unmarshal(b, &c)
	if err != nil {
		msg := fmt.Errorf("erro ao adicionar cliente (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	err = s.add.AddCustomer(c)
	if err != nil {
		msg := fmt.Errorf("erro ao adicionar cliente (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, "cliente adicionado com sucesso")

}

// AddUser implements the route to add an user to the table
func (s Serv) addUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	u := model.User{}
	err = json.Unmarshal(b, &u)
	if err != nil {
		msg := fmt.Errorf("erro ao adicionar usuário (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	fmt.Println(string(b))
	fmt.Println(u)
	err = s.add.AddUser(u)
	if err != nil {
		msg := fmt.Errorf("erro ao adicionar usuário (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, "usuário adicionado com sucesso")

}

// AddWarning handles the route to add a new warning to the db
func (s Serv) AddWarning(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	warning := model.Warning{}
	err = json.Unmarshal(b, &warning)
	if err != nil {
		msg := fmt.Errorf("erro ao adicionar aviso (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	err = s.add.AddWarning(warning)
	if err != nil {
		msg := fmt.Errorf("erro ao adicionar aviso (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, "aviso adicionado com sucesso")

}

func (s Serv) addPublicFunc(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	salarioList := []model.Funcionario{}
	err = json.Unmarshal(b, &salarioList)
	if err != nil {
		msg := fmt.Errorf("erro ao adicionar funcionário público (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	err = s.add.AddPublicFunc(salarioList...)
	if err != nil {
		msg := fmt.Errorf("erro ao adicionar funcionário público (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	model.ToJSON(w, "funcionário público adicionado com sucesso")
}

// Login handles the authorization to the user
func (s *Serv) Login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := fmt.Errorf("erro ao fazer login (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	var user model.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		msg := fmt.Errorf("erro ao fazer login (%v)", err)
		model.ErrorResponse(w, msg, http.StatusBadRequest)
		return
	}
	password := user.Pass
	user, err = s.read.GetUserByEmail(user.Email)
	if err != nil {
		msg := fmt.Errorf("login não autorizado")
		model.ErrorResponse(w, msg, http.StatusUnauthorized)
		return
	}
	err = model.IsPassword(user.Pass, password)
	if err != nil {
		msg := fmt.Errorf("login não autorizado")
		model.ErrorResponse(w, msg, http.StatusUnauthorized)
		return
	}
	model.ToJSON(w, "login efetuado com sucesso")

}
