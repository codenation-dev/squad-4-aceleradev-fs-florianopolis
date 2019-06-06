package controllers

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/models"
	"codenation/squad-4-aceleradev-fs-florianopolis/utils"
	"codenation/squad-4-aceleradev-fs-florianopolis/validations"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// GetUsers send the list of users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetUsers()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJSON(w, users)
}

// PostUser insert a new user on the db
func PostUser(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var u models.User
	err := json.Unmarshal(body, &u)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}

	u, err = validations.ValidateNewUser(u)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}

	_, err = models.NewUser(u)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(
		w, utils.DefaultResponse{
			"Usuário cadastrado com sucesso!",
			http.StatusCreated,
		})
}
