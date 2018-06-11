package main

import (
	"net/http"

	"github.com/FernandoCagale/serverless-go/src/render"

	errors "github.com/FernandoCagale/serverless-go/src/error"
	"github.com/FernandoCagale/serverless-go/src/handlers"
	"github.com/FernandoCagale/serverless-go/src/middleware"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	r := mux.NewRouter()

	r.Use(middleware.BindDb)

	r.HandleFunc("/task", handlers.FindAll).Methods("GET")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	ar := mux.NewRouter()

	mw := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			render.ResponseError(w, errors.AddUnauthorizedError(err))
		},
		SigningMethod: jwt.SigningMethodHS512,
	})

	ar.Use(middleware.BindDb)
	ar.HandleFunc("/api/task", handlers.FindAll).Methods("GET")

	an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(ar))
	r.PathPrefix("/api").Handler(an)

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run("127.0.0.1:3000")
}
