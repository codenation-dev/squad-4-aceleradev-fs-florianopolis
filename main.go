package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // postgres
)

const (
	DRIVER_NAME = "postgres"
	DB_USER     = "postgres"
	DB_PASSWORD = "12345"
	DB_NAME     = "uati"
	SSLMODE     = "disable"
	HOST        = "172.17.0.2"
	PORT        = "5432"
)

type App struct {
	db     *sql.DB
	router *mux.Router
}

var app App

func (a *App) connectDB() {
	connString := fmt.Sprintf(fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		DB_USER, DB_PASSWORD, HOST, PORT, DB_NAME, SSLMODE,
	))

	db, err := sql.Open("postgres", connString)
	a.db = db
	if err != nil {
		log.Fatal(err)
	}
}

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode("API Banco Uati")
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) NewRouter() {
	r := mux.NewRouter()
	r.HandleFunc("/", getHome).Methods("GET")
	r.HandleFunc("/customers", a.getCustomers).Methods("GET")
	r.HandleFunc("/customers", a.PostCustomers).Methods("POST")
	r.HandleFunc("/customers", a.PutCustomers).Methods("PUT").Queries("id", "{id}")
	r.HandleFunc("/customers", a.DeleteCustomers).Methods("DELETE").Queries("id", "{id}")
	a.router = r
}

func main() {

	app.connectDB()
	app.NewRouter()

	apiPort := ":3000"
	http.ListenAndServe(apiPort, app.router)

}

// Customer of the bank
type Customer struct {
	ID          int     `json:"cid"`
	Name        string  `json:"name"`
	Wage        float32 `json:"wage"`
	IsPublic    int8    `json:"is_public"`
	SentWarning string  `json:"sent_warning"` //TODO: Isso pode se tornar o ID da tabela warning
}

// DeleteCustomers handles method "DELETE" to route /customer
func (a *App) DeleteCustomers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	_, err = a.db.Exec("DELETE FROM customers WHERE id=$1", id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Cadastro não encontrado")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Delete efetuado com sucesso")
}

// PutCustomers handles method "PUT" to route /customer
func (a *App) PutCustomers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	cc, err := ReadCustomers(a.db, id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "Cadastro não encontrado")
		return
	}
	c := cc[0]

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
	}

	query := "UPDATE customers SET name=$1, wage=$2, is_public=$3, sent_warning=$4 WHERE id=$5"
	_, err = a.db.Exec(query, c.Name, c.Wage, c.IsPublic, c.SentWarning, c.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "Update realizado com sucesso")
}

// PostCustomers handles the POST route to /customer
func (a *App) PostCustomers(w http.ResponseWriter, r *http.Request) {
	c := Customer{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		log.Fatal(err)
	}

	query := `INSERT INTO customers (name, wage, is_public, sent_warning) 
	VALUES ($1, $2, $3, $4)
	RETURNING id`
	_, err = a.db.Exec(query, c.Name, c.Wage, c.IsPublic, c.SentWarning)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(w, "Cliente cadastrado com sucesso")
}

func (a *App) getCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := ReadCustomers(a.db, 0)
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

// ReadCustomers read customers from the DB
// If id = 0, it will return all customers
func ReadCustomers(db *sql.DB, id int) ([]Customer, error) {
	var err error
	var query string
	var rows *sql.Rows

	if id == 0 {
		query = "SELECT * FROM customers"
		rows, err = db.Query(query)

	} else {
		query = "SELECT * FROM customers WHERE id=$1"
		rows, err = db.Query(query, id)
	}
	if err != nil {
		return nil, err
	}
	customers := []Customer{}
	for rows.Next() {
		c := Customer{}
		err = rows.Scan(&c.ID, &c.Name, &c.Wage, &c.IsPublic, &c.SentWarning)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}
	return customers, nil
}
