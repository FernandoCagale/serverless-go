package infra

import (
	"net/http"

	errors "github.com/FernandoCagale/serverless-go/src/error"
	"github.com/FernandoCagale/serverless-go/src/render"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
)

func GetConfigJWT() *jwtmiddleware.Options {
	return &jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte("secret"), nil
		},
		ErrorHandler: func(w http.ResponseWriter, r *http.Request, err string) {
			render.ResponseError(w, errors.AddUnauthorizedError(err))
		},
		SigningMethod: jwt.SigningMethodHS512,
	}
}

func MethodNotFoundHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.ResponseError(w, errors.AddNotFoundError("Not Found"))
	})
}

func MethodNotAllowedHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		render.ResponseError(w, errors.AddMethodNotAllowedError("Method Not Allowed"))
	})
}
