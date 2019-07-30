package rest

import (
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/adding"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
)
 
func getPublicFunc(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		publicFuncs, err := reader.GetPublicFunc(r.Form)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}
		respondWithJSON(w, http.StatusOK, publicFuncs)
	}
}

func statsPublicFunc(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		publicStats, err := reader.StatsPublicFunc(r.Form)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, publicStats)
	}
}

func distrPublicFunc(reader reading.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		distPublicStats, err := reader.DistPublicFunc(r.Form)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		respondWithJSON(w, http.StatusOK, distPublicStats)
	}
}

func importPublicFunc(adder adding.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}

		err = adder.ImportPublicFunc(r.Form)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, err)
			return
		}
		respondWithJSON(w, http.StatusOK, "importação realizada com sucesso")
	}
}
