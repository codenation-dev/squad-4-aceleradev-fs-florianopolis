package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/http/rest"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/adding"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/deleting"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/reading"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/service/updating"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/pkg/storage/postgres"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := setup()

	apiPort := ":3000"
	fmt.Printf("API running on port%s\n", apiPort)

	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "content-type", "Origin", "Accept"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	//allowCred := handlers.AllowCredentials()
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	if err := http.ListenAndServe(apiPort, handlers.CORS(originsOk, headersOk, methodsOk)(router)); err != nil {
		log.Fatal(err)
	}
}

func setup() *mux.Router {
	dbName := "uati"

	repo := postgres.NewStorage(dbName)

	adder := adding.NewService(repo)
	reader := reading.NewService(repo)
	updater := updating.NewService(repo)
	deleter := deleting.NewService(repo)

	router := rest.NewRouter(adder, reader, updater, deleter)

	return router
}
