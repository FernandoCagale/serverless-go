package main

import (
	"github.com/FernandoCagale/serverless-go/src/infra"
	"github.com/FernandoCagale/serverless-go/src/routers"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/negroni"
	"github.com/urfave/negroni"
)

var initialized = false

var negroniLambda *negroniadapter.NegroniAdapter

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if !initialized {
		public := routers.InitRoutesPublic()
		auth := routers.InitRoutesAuth()

		mw := jwtmiddleware.New(*infra.GetConfigJWT())

		an := negroni.New(negroni.HandlerFunc(mw.HandlerWithNext), negroni.Wrap(auth))
		public.PathPrefix("/api").Handler(an)

		n := negroni.Classic()
		n.UseHandler(public)

		negroniLambda = negroniadapter.New(n)
		initialized = true
	}

	return negroniLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
