package rest

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/adding"
	"codenation/squad-4-aceleradev-fs-florianopolis/deleting"
	"codenation/squad-4-aceleradev-fs-florianopolis/reading"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type serv struct {
	add  adding.Service
	read reading.Service
	// u updating.Service
	del deleting.Service
}

// Handler handle the API routes
func Handler(
	add adding.Service,
	read reading.Service,
	del deleting.Service,
	// u updating.Service,
) http.Handler {
	s := serv{add: add, read: read, del: del}
	// u: u,

	router := mux.NewRouter()
	router.HandleFunc("/", getHome).Methods("GET")
	router.HandleFunc("/customer/all", s.getAllCustomers).Methods("GET")
	router.HandleFunc("/customer", s.getCustomer).Methods("GET").Queries("id", "{id}")
	router.HandleFunc("/customer", s.getCustomerByName).Methods("GET").Queries("name", "{pattern}")
	router.HandleFunc("/customer", s.addCustomer).Methods("POST")
	router.HandleFunc("/customer", s.deleteCustomerByID).Methods("DELETE").Queries("id", "{id}")
	// r.HandleFunc("/customers", a.PutCustomers).Methods("PUT").Queries("id", "{id}")

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

func (s serv) deleteCustomerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	err = s.del.DeleteCustomerByID(id)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		err = json.NewEncoder(w).Encode("Usuário deletado com sucesso")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (s serv) addCustomer(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	c := adding.Customer{}
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
	c := reading.Customer{ID: id}
	c, err = s.read.GetCustomerByID(c)
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
		err := json.NewEncoder(w).Encode(fmt.Sprintf("Sorry, something bad happened: %v", err))
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

// // DeleteCustomers handles method "DELETE" to route /customer
// func (a *App) Delete(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	_, err = a.db.Exec("DELETE FROM customers WHERE id=$1", id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintln(w, "Cadastro não encontrado")
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintln(w, "Delete efetuado com sucesso")
// }

// // PutCustomers handles method "PUT" to route /customer
// func (a *App) PutCustomers(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	cc, err := ReadCustomers(a.db, id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintln(w, "Cadastro não encontrado")
// 		return
// 	}
// 	c := cc[0]

// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = json.Unmarshal(b, &c)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	query := "UPDATE customers SET name=$1, wage=$2, is_public=$3, sent_warning=$4 WHERE id=$5"
// 	_, err = a.db.Exec(query, c.Name, c.Wage, c.IsPublic, c.SentWarning, c.ID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Fprintln(w, "Update realizado com sucesso")
// }

// // PostCustomers handles the POST route to /customer
// func (a *App) PostCustomers(w http.ResponseWriter, r *http.Request) {
// 	c := Customer{}
// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	err = json.Unmarshal(b, &c)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	query := `INSERT INTO customers (name, wage, is_public, sent_warning)
// 	VALUES ($1, $2, $3, $4)
// 	RETURNING id`
// 	_, err = a.db.Exec(query, c.Name, c.Wage, c.IsPublic, c.SentWarning)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Fprintln(w, "Cliente cadastrado com sucesso")
// }
