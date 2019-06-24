package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getHome(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode("API Banco Uati")
	if err != nil {
		log.Fatal(err)
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

func (s serv) getAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := s.read.GetAllUsers()
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(fmt.Sprintf("Erro lendo o banco de dados: %v", err))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		w.Header().Set("Content-type", "application/json")
		b, err := json.Marshal(users)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
}

func (s serv) getCustomerByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	//TODO: validar pattern para o modelo da codenation
	customers, err := s.read.GetCustomerByName(params["name"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Houve um problema na procura deste cliente: %v", err)
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

func (s serv) getUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	//TODO: validar pattern para o modelo da codenation
	users, err := s.read.GetUserByEmail(params["email"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Houve um problema na procura deste usuario: %v", err)
		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		b, err := json.Marshal(users)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
}

func (s serv) getCustomerByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
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

func (s serv) getUserByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	user, err := s.read.GetUserByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Houve um problema na procura deste usuário: %v", err)
		err := json.NewEncoder(w).Encode(msg)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		b, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
}
