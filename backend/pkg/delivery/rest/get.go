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

// All
func getPayload(w http.ResponseWriter, r *http.Request, payload interface{}, err error) (http.ResponseWriter, *http.Request) {

	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(fmt.Sprintf("Erro lendo o banco de dados: %v", err))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		w.Header().Set("Content-type", "application/json")
		b, err := json.Marshal(payload)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
	return w, r
}

func (s Serv) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := s.read.GetAllCustomers()
	w, r = getPayload(w, r, customers, err)
	// if err != nil {
	// 	w.Header().Set("Content-type", "application/json")
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	err := json.NewEncoder(w).Encode(fmt.Sprintf("Erro lendo o banco de dados: %v", err))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// } else {
	// 	w.Header().Set("Content-type", "application/json")
	// 	b, err := json.Marshal(customers)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	w.Write(b)
	// }
}

func (s Serv) getAllUsers(w http.ResponseWriter, r *http.Request) {
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

func (s Serv) getAllWarnings(w http.ResponseWriter, r *http.Request) {
	warnings, err := s.read.GetAllWarnings()
	if err != nil {
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(w).Encode(fmt.Sprintf("Erro lendo o banco de dados: %v", err))
		if err != nil {
			log.Fatal(err)
		}
	} else {
		w.Header().Set("Content-type", "application/json")
		b, err := json.Marshal(warnings)
		if err != nil {
			log.Fatal(err)
		}
		w.Write(b)
	}
}

// ByName
func (s Serv) getCustomerByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	//TODO: validar pattern para o modelo da codenation
	customers, err := s.read.GetCustomerByName(params["pattern"])
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

func (s Serv) getUserByEmail(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	//TODO: validar pattern para o modelo da codenation
	user, err := s.read.GetUserByEmail(params["pattern"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Houve um problema na procura deste usuario: %v", err)
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

func (s *Serv) getWarningByCustomer(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	payload, err := s.read.GetWarningByCustomer(params["pattern"])
	w, r = getPayload(w, r, payload, err)
}

func (s *Serv) getWarningByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	payload, err := s.read.GetWarningByUser(params["pattern"])
	w, r = getPayload(w, r, payload, err)
}

func (s *Serv) getPublicByWage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	pattern, err := strconv.ParseFloat(params["pattern"], 32)
	if err != nil {
		w, r = getPayload(w, r, nil, err)
		return
	}
	payload, err := s.read.GetPublicByWage(float32(pattern))
	w, r = getPayload(w, r, payload, err)
}

//ByID
func (s Serv) getCustomerByID(w http.ResponseWriter, r *http.Request) {
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

func (s Serv) getUserByID(w http.ResponseWriter, r *http.Request) {
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

func (s Serv) getWarningByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal(err)
	}
	user, err := s.read.GetWarningByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Houve um problema na procura desta mensagem: %v", err)
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

func (s Serv) getPublicByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w, r, id := validateID(w, r)
	// params := mux.Vars(r)
	// id, err := strconv.Atoi(params["id"])
	// if err != nil {
	// 	log.Fatal(err)
	// }

	user, err := s.read.GetPublicByID(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Houve um problema na procura deste funcionário público: %v", err)
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
