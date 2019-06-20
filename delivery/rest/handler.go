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
		panic(err)
	}
}

func (s serv) updateCustomer(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}

	customer, err := s.read.GetCustomerByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			panic(err)
		}
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &customer)
	if err != nil {
		panic(err)
	}
	// updateCustomer := entity.Customer{
	// 	Name: customer.Name, customer.Wage, customer.IsPublic, customer.SentWarning,
	// }

	err = s.update.UpdateCustomer(customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			panic(err)
		}
		return
	} else {
		err = json.NewEncoder(w).Encode("Usuário modificado com sucesso")
		if err != nil {
			panic(err)
		}
	}
}

func (s serv) deleteCustomerByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	err = s.del.DeleteCustomerByID(id)

	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			panic(err)
		}
	} else {
		err = json.NewEncoder(w).Encode("Usuário deletado com sucesso")
		if err != nil {
			panic(err)
		}
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
		panic(err)
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
			panic(err)
		}
	} else {
		b, err := json.Marshal(customers)
		if err != nil {
			panic(err)
		}
		w.Write(b)
	}
}

func (s serv) getCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		panic(err)
	}
	// c := reading.Customer{ID: id}
	c, err := s.read.GetCustomerByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Houve um problema na procura deste cliente: %v", err)
		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			panic(err)
		}
	} else {
		b, err := json.Marshal(c)
		if err != nil {
			panic(err)
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
			panic(err)
		}
	} else {
		w.Header().Set("Content-type", "application/json")
		b, err := json.Marshal(customers)
		if err != nil {
			panic(err)
		}
		w.Write(b)
	}
}

// // DeleteCustomers handles method "DELETE" to route /customer
// func (a *App) Delete(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	id, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		panic(err)
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
// 		panic(err)
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
// 		panic(err)
// 	}
// 	err = json.Unmarshal(b, &c)
// 	if err != nil {
// 		panic(err)
// 	}

// 	query := "UPDATE customers SET name=$1, wage=$2, is_public=$3, sent_warning=$4 WHERE id=$5"
// 	_, err = a.db.Exec(query, c.Name, c.Wage, c.IsPublic, c.SentWarning, c.ID)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Fprintln(w, "Update realizado com sucesso")
// }

// // PostCustomers handles the POST route to /customer
// func (a *App) PostCustomers(w http.ResponseWriter, r *http.Request) {
// 	c := Customer{}
// 	b, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = json.Unmarshal(b, &c)
// 	if err != nil {
// 		panic(err)
// 	}

// 	query := `INSERT INTO customers (name, wage, is_public, sent_warning)
// 	VALUES ($1, $2, $3, $4)
// 	RETURNING id`
// 	_, err = a.db.Exec(query, c.Name, c.Wage, c.IsPublic, c.SentWarning)

// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Fprintln(w, "Cliente cadastrado com sucesso")
// }
