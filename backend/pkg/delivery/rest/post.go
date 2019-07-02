package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/emailserv"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/utils"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/entity"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/importing"
)

func (s *Serv) sendEmail(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	email := params["userEmail"]
	fmt.Println(params)
	user, err := s.read.GetUserByEmail(email)
	if err != nil {
		err := json.NewEncoder(w).Encode("Usuário não encontrado")
		if err != nil {
			log.Fatalf("Error in json.NewEncoder: %v", err)
		}
		// w.WriteHeader(http.StatusNotFound)
		return
	}

	publicFuncs := []entity.PublicFunc{}
	err = json.NewDecoder(r.Body).Decode(&publicFuncs)
	if err != nil {
		err := json.NewEncoder(w).Encode("Corpo do email inválido")
		if err != nil {
			log.Fatalf("Error in json.NewDecoder: %v", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = emailserv.Send(publicFuncs, user)
	if err != nil {
		err := json.NewEncoder(w).Encode("Erro no envio do email")
		if err != nil {
			log.Fatalf("Error in json.NewDecoder: %v", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode("Email enviado com sucesso")
	if err != nil {
		log.Fatalf("Error in json.NewDecoder: %v", err)
	}
	// w.WriteHeader(http.StatusBadRequest)
	return
}

// ImportPublicFuncFile - Download e importação de arquivo de funcionário público do estado de SP
func (s Serv) ImportPublicFuncFile(w http.ResponseWriter, r *http.Request) {
	publicFuncs, err := importing.ImportPublicFunc()
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = s.add.AddPublicFunc(publicFuncs...)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Arquivo de funcionários públicos importado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

// ImportCustomerFile importa o arquivo 'clientes.csv'
func (s Serv) ImportCustomerFile(w http.ResponseWriter, r *http.Request) {
	customers, err := importing.ImportClientesCSV("backend/cmd/data/clientes.csv")
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, customer := range customers {
		err := s.add.AddCustomer(customer)
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Arquivo de clientes importado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

// AddCustomer adiciona clientes ao banco de dados
func (s Serv) AddCustomer(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	c := entity.Customer{}
	err = json.Unmarshal(b, &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.add.AddCustomer(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Cliente adicionado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

// AddUser godoc
// @Summary Adiciona um usuário ao BD
// @Description Adiciona um usuário ao BD
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entity.User
// @Failure 400 {object} URL não encontrada
// @Failure 404 {object} Erro na solicitação
// @Router /user [POST]
func (s Serv) addUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	u := entity.User{}
	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// bPass, err := utils.Bcrypt(u.Pass)
	// if err != nil {
	// 	panic(err)
	// }
	// u.Pass = string(bPass)

	err = s.add.AddUser(u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Usuário adicionado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

// AddWarning handles the route to add a new warning to the db
func (s Serv) AddWarning(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	warning := entity.Warning{}
	err = json.Unmarshal(b, &warning)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.add.AddWarning(warning)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Aviso adicionado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

func (s Serv) addPublicFunc(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	publicFuncs := []entity.PublicFunc{}
	err = json.Unmarshal(b, &publicFuncs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.add.AddPublicFunc(publicFuncs...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("funcionário público adicionado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

// Login handles the authorization to the user
func (s *Serv) Login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	var user entity.User
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	password := user.Pass
	user, err = s.read.GetUserByEmail(user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	err = utils.IsPassword(user.Pass, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Fatal(err)
	}
}
