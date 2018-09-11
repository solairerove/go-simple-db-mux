package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gorilla/mux"
)

// App ... tbd
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// Initialize ... tbd
func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
}

// Run ... tbd
func (a *App) Run(addr string) {}
