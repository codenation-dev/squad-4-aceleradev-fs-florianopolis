package models

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq" // postgres
)

//App carries the db and the router to dependency injection
type App struct {
	db     *sql.DB
	router *mux.Router
}
