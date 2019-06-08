package controllers

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/models"
	"codenation/squad-4-aceleradev-fs-florianopolis/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

// GetWarnings send the list of warnings already sent
func GetWarnings(w http.ResponseWriter, r *http.Request) {
	users, err := models.GetWarnings()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJSON(w, users)
}

// GetWarningsByUser returns the list of customers from a user
func GetWarningsByUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sentTo, _ := params["sent_to"]
	warnings, err := models.GetWarningsByUser(sentTo)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJSON(w, warnings)
}

// PostWarning insert a new warning on the db
func PostWarning(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var wa models.Warning
	err := json.Unmarshal(body, &wa)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}

	_, err = models.NewWarning(wa)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(
		w, utils.DefaultResponse{
			"Aviso cadastrado com sucesso!",
			http.StatusCreated,
		})
}

//TODO: aproveitar este controller e enviar o email???
