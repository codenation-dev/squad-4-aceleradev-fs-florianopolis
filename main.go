package main

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/adding"
	"codenation/squad-4-aceleradev-fs-florianopolis/deleting"
	"codenation/squad-4-aceleradev-fs-florianopolis/delivery/rest"
	"codenation/squad-4-aceleradev-fs-florianopolis/reading"
	"codenation/squad-4-aceleradev-fs-florianopolis/storage/postgres"
	"codenation/squad-4-aceleradev-fs-florianopolis/updating"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // postgres
)

func main() {

	// set services
	var adder adding.Service
	var reader reading.Service
	var deleter deleting.Service
	var updater updating.Service

	// If have more than one storage types, make the case/switch here
	s, err := postgres.NewStorage(postgres.Connect())
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
	log.Fatal(http.ListenAndServe(":3000", router))

}
