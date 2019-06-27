// @title Captação de Clientes Banco Uati
// @version 0.1
// @description Projeto final para o AceleraDev FullStack presencial CodeNation
// @termsOfService http://terms.of.service

// @contact.name API Support
// @contact.url http://www.contact.url
// @contact.email support@email.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /

// @securityDefinitions.basic TODO: not implemented yet

// @securityDefinitions.apikey ApiKeyAuth TODO: not implemented yet
// @in header TODO: not implemented yet
// @name Authorization TODO: not implemented yet

// @securitydefinitions.oauth2.application OAuth2Application TODO: not implemented yet
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit TODO: not implemented yet
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password TODO: not implemented yet
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode TODO: not implemented yet
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// Access-Control-Allow-*
package main

import (
	"fmt"
	"log"
	"net/http"

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
