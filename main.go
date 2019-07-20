// Projeto final do AceleraDev Full Stack Go + React Presencial
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/database"

	"github.com/rs/cors"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/deleting"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/delivery/rest"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/reading"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/storage/postgres"
	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/updating"

	"github.com/codenation-dev/squad-4-aceleradev-fs-florianopolis/backend/pkg/adding"
	_ "github.com/lib/pq" // postgres
)

func main() {

	// set services
	var adder adding.Service
	var reader reading.Service
	var deleter deleting.Service
	var updater updating.Service

	// If have more than one storage types, make the case/switch here
	database.DBCon = postgres.Connect()
	s, err := postgres.NewStorage(database.DBCon)
	if err != nil {
		log.Fatalf("could not set new storage: %v", err)
	}

	adder = adding.NewService(s)
	reader = reading.NewService(s)
	deleter = deleting.NewService(s)
	updater = updating.NewService(s)

	// set uo HTTP server
	router := rest.Handler(
		adder,
		reader,
		deleter,
		updater,
	)

	fmt.Println("Server running ou port 3000")
	r := cors.Default().Handler(router)
	log.Fatal(http.ListenAndServe(":3000", r))

}
