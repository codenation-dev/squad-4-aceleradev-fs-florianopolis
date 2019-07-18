package rest

import (
	"html/template"
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/entity"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/adding"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
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

func addUser(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "new_user.html", nil)
	}
}

func addUserProcess(adder adding.Service, tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// u := entity.User{}
		// b, _ := ioutil.ReadAll(r.Body)
		// err := json.Unmarshal(b, &u)
		// if err != nil {
		// 	respondWithError(w, http.StatusUnprocessableEntity, entity.ErrUnmarshal)
		// 	return
		// }
		user := entity.User{}
		// _ = r.ParseForm()
		// fmt.Println(r.Form)

		user.Email = r.FormValue("email")
		user.Password = r.FormValue("pass")
		err := adder.AddUser(user)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, err)
			return
		}

		tpl.ExecuteTemplate(w, "login.html", nil)
	}

}
