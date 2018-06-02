package main

import (
	"github.com/FernandoCagale/serverless-go/src/handlers"
	"github.com/FernandoCagale/serverless-go/src/middleware"
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

		r.HandleFunc("/task", handlers.Create).Methods("POST")
		r.HandleFunc("/task", handlers.FindAll).Methods("GET")
		r.HandleFunc("/task/{id}", handlers.FindById).Methods("GET")
		r.HandleFunc("/task/{id}", handlers.UpdateById).Methods("PUT")
		r.HandleFunc("/task/{id}", handlers.DeleteById).Methods("DELETE")

		r.Use(middleware.BindDb)

		gorillaLambda = gorillamux.New(r)
		initialized = true
	}

	return gorillaLambda.Proxy(req)
}

func main() {
	lambda.Start(Handler)
}
