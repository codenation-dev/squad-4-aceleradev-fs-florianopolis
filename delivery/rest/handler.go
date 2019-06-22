package rest

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/adding"
	"codenation/squad-4-aceleradev-fs-florianopolis/deleting"
	"codenation/squad-4-aceleradev-fs-florianopolis/entity"
	"codenation/squad-4-aceleradev-fs-florianopolis/reading"
	"codenation/squad-4-aceleradev-fs-florianopolis/updating"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type serv struct {
	add    adding.Service
	read   reading.Service
	del    deleting.Service
	update updating.Service
}

// Handler handle the API routes
func Handler(
	add adding.Service,
	read reading.Service,
	del deleting.Service,
	update updating.Service,
) http.Handler {
	s := serv{add: add, read: read, del: del, update: update}

	router := mux.NewRouter()
	router.HandleFunc("/", getHome).Methods("GET")
	router.HandleFunc("/customer/all", s.getAllCustomers).Methods("GET")
	router.HandleFunc("/customer", s.getCustomer).Methods("GET").Queries("id", "{id}")
	router.HandleFunc("/customer", s.getCustomerByName).Methods("GET").Queries("name", "{pattern}")
	router.HandleFunc("/customer", s.addCustomer).Methods("POST")
	router.HandleFunc("/customer", s.deleteCustomerByID).Methods("DELETE").Queries("id", "{id}")
	router.HandleFunc("/customer", s.updateCustomer).Methods("PUT").Queries("id", "{id}")

	return router
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("API Banco Uati")
	if err != nil {
		log.Fatal(err)
	}
}

func (s serv) updateCustomer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	customer, err := s.read.GetCustomerByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	err = json.Unmarshal(b, &customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = s.update.UpdateCustomer(customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = json.NewEncoder(w).Encode("Usuário modificado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

func (s serv) deleteCustomerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	err = s.del.DeleteCustomerByID(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("Usuário deletado com sucesso")
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
	err = json.NewEncoder(w).Encode("Usuário adicionado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

func (s serv) getCustomerByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	//TODO: validar pattern para o modelo da codenation
	customers, err := s.read.GetCustomerByName(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Houve um problema na procura deste nome: %v", err)
		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		b, err := json.Marshal(customers)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
}

func (s serv) getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	// c := reading.Customer{ID: id}
	c, err := s.read.GetCustomerByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Houve um problema na procura deste cliente: %v", err)
		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		b, err := json.Marshal(c)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
}

func (s serv) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := s.read.GetAllCustomers()
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(fmt.Sprintf("Erro lendo o banco de dados: %v", err))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		w.Header().Set("Content-type", "application/json")
		b, err := json.Marshal(customers)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
}
