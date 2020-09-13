package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/kserevena/perkbox-tech-test/services"
)

func initRoutes(router *mux.Router, db *sql.DB) *mux.Router {
	services.NewIndex(router, db)
	services.NewCreate(router, db)
	return router
}
