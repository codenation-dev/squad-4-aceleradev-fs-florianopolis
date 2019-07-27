package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/emailserv"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/adding"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/deleting"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/updating"

	"github.com/gorilla/mux"
)

// NewRouter implements handlers to all routes
func NewRouter(adder adding.Service, reader reading.Service, updater updating.Service, deleter deleting.Service) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/login", login(reader)).Methods(http.MethodPost, http.MethodOptions)
	router.Handle("/", getIndex()).Methods(http.MethodGet)

	router.Handle("/user/{email}", getUser(reader)).Methods(http.MethodGet)
	router.Handle("/user/{email}", deleteUser(deleter)).Methods(http.MethodDelete)
	router.Handle("/user", addUser(adder)).Methods(http.MethodPost)
	router.Handle("/user", updateUser(updater)).Methods(http.MethodPut)

	router.Handle("/public_func", getPublicFunc(reader)).Methods(http.MethodGet, http.MethodOptions)
	router.Handle("/public_func/import", importPublicFunc(adder)).Methods(http.MethodGet)
	router.Handle("/public_func/stats", statsPublicFunc(reader)).Methods(http.MethodGet)

	router.Handle("/customer", getCustomer(reader)).Methods(http.MethodGet)
	router.Handle("/customer/import", importCustomer(adder)).Methods(http.MethodGet)
 
	router.Handle("/email_to", sendEmail(reader)).Methods(http.MethodPost)
	// router.Handle("/fetch/data/compare_customer_x_public_func/{company}/{uf}/{year}/{month}", compareCustomerPublicFunc(reader)).Methods(http.MethodGet)
	// router.Handle("/fetch/data/public_func_above_wage/{uf}/{year}/{month}/{wage}", getPublicFincByWage(reader)).Methods(http.MethodGet)

	router.Handle("/query", handleQuery(reader)).Methods(http.MethodGet)

	router.Use(authorize)
	return router
}

func handleQuery(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// err := r.ParseForm()
		// if err != nil {
		// 	respondWithError(w, http.StatusBadRequest, err)
		// 	return
		// }
		q := r.FormValue("q")
		if q == "" {
			respondWithError(w, http.StatusBadRequest, errors.New("precisa indicar um parametro 'q'"))
			return
		}
		offset := r.FormValue("Offset")
		n, err := strconv.Atoi(offset)
		if err != nil {
			offset = ""
		}
		if offset == "" || n == 0 || n > 50 {
			offset = "50"
		}

		page := r.FormValue("Page")
		if page == "" {
			page = "1"
		}

		resp, err := reader.Query(q, offset, page)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		respondWithJSON(w, http.StatusOK, resp)
	}
}

func sendEmail(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		email := entity.Email{}
		err = json.Unmarshal(b, &email)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		err = emailserv.Send(email)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, "email enviado com sucesso")
	}
}

func getIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}

func respondWithError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	fmt.Fprint(w, err.Error())
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.WriteHeader(code)
	w.Write(response)
}
