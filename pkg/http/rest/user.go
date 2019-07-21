package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/adding"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/deleting"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/updating"
	"github.com/gorilla/mux"
)

func getUser(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		u := entity.User{}
		u, err := reader.GetUser(params["email"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}
		respondWithJSON(w, http.StatusOK, u)
	}
}

func addUser(adder adding.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		assertError := func(err error) {
			if err != nil {
				http.Error(w, fmt.Sprintf("erro ao adicionar usuário (%s)", err.Error()), http.StatusBadRequest)
				return
			}
		}

		b, err := ioutil.ReadAll(r.Body)
		assertError(err)

		newUser := entity.User{}
		err = json.Unmarshal(b, &newUser)
		assertError(err)

		err = adder.AddUser(newUser)
		assertError(err)

		respondWithJSON(w, http.StatusOK, "usuário adicionado com sucesso")
	}
}

func deleteUser(deleter deleting.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		err := deleter.DeleteUser(params["email"])
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}
		respondWithJSON(w, http.StatusOK, "usuário deletado com sucesso")
	}
}

func updateUser(updater updating.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user := entity.User{}
		b, err := ioutil.ReadAll(r.Body)

		if err != nil {
			respondWithError(w, http.StatusInternalServerError, errors.New("erro no ioutil"))
			return
		}
		err = json.Unmarshal(b, &user)
		if err != nil {
			respondWithError(w, http.StatusUnprocessableEntity, entity.ErrUnmarshal)
			return
		}
		err = updater.ChangePassword(user)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}
		respondWithJSON(w, http.StatusOK, "usuário modificado com sucesso")
	}
}
