package routers

import (
	"github.com/FernandoCagale/serverless-go/src/handlers"
	"github.com/gorilla/mux"
)

func taskRoutes(r *mux.Router) {
	r.HandleFunc("/api/task/{id}", handlers.DeleteById).Methods("DELETE")
	r.HandleFunc("/api/task/{id}", handlers.UpdateById).Methods("PUT")
	r.HandleFunc("/api/task/{id}", handlers.FindById).Methods("GET")
	r.HandleFunc("/api/task", handlers.FindAll).Methods("GET")
	r.HandleFunc("/api/task", handlers.Create).Methods("POST")
}
