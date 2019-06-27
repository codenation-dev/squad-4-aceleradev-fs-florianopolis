package rest

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/backend/pkg/importing"
	"codenation/squad-4-aceleradev-fs-florianopolis/backend/pkg/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (s serv) importPublicFuncFile(w http.ResponseWriter, r *http.Request) {
	publicFuncs, err := importing.ImportPublicFunc()
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// err = s.add.LoadPublicFuncFile()
	// if err != nil {
	// 	log.Fatal(err)
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	return
	// }

	// channel := make(chan int)
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

func (s serv) importCustomerFile(w http.ResponseWriter, r *http.Request) {
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

func (s serv) addCustomer(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	c := entity.Customer{}
	err = json.Unmarshal(b, &c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.add.AddCustomer(c)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Cliente adicionado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

func (s serv) addUser(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	u := entity.User{}
	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.add.AddUser(u)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Usuário adicionado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

func (s serv) addWarning(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	warning := entity.Warning{}
	err = json.Unmarshal(b, &warning)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.add.AddWarning(warning) // TODO: tratar este possível erro?

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Aviso adicionado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

func (s serv) addPublicFunc(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	publicFunc := entity.PublicFunc{}
	err = json.Unmarshal(b, &publicFunc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.add.AddPublicFunc(publicFunc) // TODO: tratar este possível erro?

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("funcionário público adicionado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}
