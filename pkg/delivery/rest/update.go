package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

	err = json.NewEncoder(w).Encode("Cliente modificado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
}

func (s serv) updateUser(w http.ResponseWriter, r *http.Request) {

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

	user, err := s.read.GetUserByID(id)
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
	err = json.Unmarshal(b, &user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = s.update.UpdateUser(user)
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

func (s serv) updateWarning(w http.ResponseWriter, r *http.Request) {
	w, r, id := validateID(w, r)

	// params := mux.Vars(r)
	// id, err := strconv.Atoi(params["id"])
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	return
	// }

	warning, err := s.read.GetWarningByID(id)
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
	err = json.Unmarshal(b, &warning)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = s.update.UpdateWarning(warning)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	err = json.NewEncoder(w).Encode("Aviso modificado com sucesso")
	if err != nil {
		log.Fatal(err)
	}

}
