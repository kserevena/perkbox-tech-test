package main

import (
	"database/sql"

	"github.com/rychanfox/couponService/services"

	"github.com/gorilla/mux"
)

func initRoutes(router *mux.Router, db *sql.DB) *mux.Router {
	services.NewIndex(router, db)
	services.NewCreate(router, db)
	return router
}
