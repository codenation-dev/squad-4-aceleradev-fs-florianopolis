package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func deleteIt(w http.ResponseWriter, r *http.Request, err error, ra int64) (http.ResponseWriter, *http.Request) {
	if err != nil || ra == 0 {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
		return w, r
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode("deletado com sucesso")
	if err != nil {
		log.Fatal(err)
	}
	return w, r
}

func validateID(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request, int) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		err = json.NewEncoder(w).Encode(fmt.Sprintf("Erro na solicitação: %v", err))
		if err != nil {
			log.Fatal(err)
		}
		return w, r, 0
	}
	return w, r, id
}

// Se funcionar, passar as que faltam para usar o deleteIt
func (s Serv) deleteCustomerByID(w http.ResponseWriter, r *http.Request) {
	w, r, id := validateID(w, r)
	ra, err := s.del.DeleteCustomerByID(id)
	w, r = deleteIt(w, r, err, ra)
}

func (s Serv) deleteWarningByID(w http.ResponseWriter, r *http.Request) {
	w, r, id := validateID(w, r)
	ra, err := s.del.DeleteWarningByID(id)
	w, r = deleteIt(w, r, err, ra)
}

func (s Serv) deletePublicByID(w http.ResponseWriter, r *http.Request) {
	w, r, id := validateID(w, r)
	ra, err := s.del.DeletePublicByID(id)
	w, r = deleteIt(w, r, err, ra)
}

func (s Serv) deleteUserByID(w http.ResponseWriter, r *http.Request) {
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
	ra, err := s.del.DeleteUserByID(id)
	if err != nil || ra == 0 {
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
