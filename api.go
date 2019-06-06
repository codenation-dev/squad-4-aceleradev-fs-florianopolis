package main

import (
	"codenation/squad-4-aceleradev-fs-florianopolis/models"
	"codenation/squad-4-aceleradev-fs-florianopolis/routes"
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "3000"
	models.TestConnection()
	fmt.Printf("Api running on port %s\n", port)
	r := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}
