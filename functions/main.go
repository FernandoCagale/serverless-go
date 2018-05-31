package main

import (
	"github.com/FernandoCagale/serverless-go/src/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

var initialized = false

var gorillaLambda *gorillamux.GorillaMuxAdapter

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	if !initialized {
		r := mux.NewRouter()

		r.HandleFunc("/api", handlers.Api).Methods("GET")
		r.HandleFunc("/public", handlers.Public).Methods("GET")

		gorillaLambda = gorillamux.New(r)
		initialized = true
	}

	return gorillaLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
