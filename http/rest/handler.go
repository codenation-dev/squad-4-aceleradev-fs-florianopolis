package rest

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/adding"
	"codenation/squad-4-aceleradev-fs-florianopolis/reading"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type serv struct {
	a adding.Service
	r reading.Service
	// u updating.Service
	// d delleting.Service
}

// Handler handle the API routes
func Handler(
	a adding.Service,
	r reading.Service,
	// u updating.Service,
	// d deleting.Service,
) http.Handler {
	s := serv{a: a, r: r}
	// u: u,
	// d: d,

	r := mux.NewRouter()
	r.HandleFunc("/", getHome).Methods("GET")
	r.Handlefunc("/customer/all", s.getAllCustomer).Methods("GET")
	r.HandleFunc("/customers", s.addCustomer).Methods("POST")
	// r.HandleFunc("/customers", a.PutCustomers).Methods("PUT").Queries("id", "{id}")
	// r.HandleFunc("/customers", a.DeleteCustomers).Methods("DELETE").Queries("id", "{id}")

	return r
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, "API Banco Uati")
	// err := json.NewEncoder(w).Encode("API Banco Uati")
	if err != nil {
		log.Fatal(err)
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
	s.a.AddCustomer(c)

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Usuário adicionado com sucesso")
}

func (s serv) getCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := s.r.GetAllCustomer()
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		err := json.NewEncoder(w).Encode("Sorry, something bad happened.")
		if err != nil {
			log.Fatal(err)
		}
	}
	w.Header().Set("Content-type", "application/json")

	b, err := json.Marshal(customers)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(b)
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
