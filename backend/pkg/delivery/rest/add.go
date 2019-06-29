package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/entity"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/importing"
)

// ImportPublicFuncFile - Download e importação de arquivo de funcionário público do estado de SP
func (s Serv) ImportPublicFuncFile(w http.ResponseWriter, r *http.Request) {
	publicFuncs, err := importing.ImportPublicFunc()
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for _, pf := range publicFuncs {
		// go func(c chan int) { // TODO: tem como fazer isso com goroutines?
		err := s.add.AddPublicFunc(pf) // TODO:melhorar velocidade dessa função
		if err != nil {
			log.Fatal(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// 	c <- 1
		// }(channel)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Arquivo de funcionários públicos importado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

// ImportCustomerFile importa o arquivo 'clientes.csv'
func (s Serv) ImportCustomerFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("importCustomerFile")
	customers, err := importing.ImportClientesCSV("clientes.csv")
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
	publicFunc := entity.PublicFunc{}
	err = json.Unmarshal(b, &publicFunc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = s.add.AddPublicFunc(publicFunc)
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
