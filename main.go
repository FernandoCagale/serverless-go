package main

import (
	"github.com/FernandoCagale/serverless-go/src/infra"
	"github.com/FernandoCagale/serverless-go/src/routers"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/urfave/negroni"
)

func main() {
	public := routers.InitRoutesPublic()
	auth := routers.InitRoutesAuth()

	mw := jwtmiddleware.New(*infra.GetConfigJWT())

	an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(auth))
	public.PathPrefix("/api").Handler(an)

	n := negroni.Classic()
	n.UseHandler(public)

	n.Run("127.0.0.1:3000")
}
