package routers

import (
	"github.com/FernandoCagale/serverless-go/src/handlers"
	"github.com/gorilla/mux"
)

func authRoutes(r *mux.Router) {
	r.HandleFunc("/public/login", handlers.Login).Methods("POST")
}
