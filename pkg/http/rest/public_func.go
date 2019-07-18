package rest

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
)

func readAllPublicFunc(tpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.ExecuteTemplate(w, "list_public_func.html", nil)
	}
}

func processReadAllPublicFunc(tpl *template.Template, reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uf := r.FormValue("uf")
		year := r.FormValue("year")
		month := r.FormValue("month")

		publicFuncs, err := reader.GetAllPublicFunc(uf, year, month)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, publicFuncs)
	}

}
