package controllers

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/models"
	"codenation/squad-4-aceleradev-fs-florianopolis/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetCustomers returns the list of customers
func GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := models.GetCustomers()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJSON(w, customers)
}

// GetCustomersPublicFuncs returns the list of customers that are public employees
func GetCustomersPublicFuncs(w http.ResponseWriter, r *http.Request) {
	customers, err := models.GetCustomersPublicFuncs()
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJSON(w, customers)
}

// GetVIPCustomers returns the list of customers with good earnings
func GetVIPCustomers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	wage, _ := strconv.ParseFloat(params["wage"], 32)
	customers, err := models.GetVIPCustomers(float32(wage))
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJSON(w, customers)
}

// GetCustomerByName returns one customer
func GetCustomerByName(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name, _ := params["name"]
	customers, err := models.GetCustomersByName(name)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	utils.ToJSON(w, customers)
}

// PostCustomer insert new customer on the db
func PostCustomer(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	var c models.Customer
	err := json.Unmarshal(body, &c)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}

	_, err = models.NewCustomer(c)
	if err != nil {
		utils.ErrorResponse(w, err, http.StatusUnprocessableEntity)
		return
	}
	utils.ToJSON(
		w, utils.DefaultResponse{
			"Cliente cadastrado com sucesso!",
			http.StatusCreated,
		})
}
