package routers

import (
	"github.com/FernandoCagale/serverless-go/src/infra"
	"github.com/FernandoCagale/serverless-go/src/middleware"
	"github.com/gorilla/mux"
)

func InitRoutesPublic() *mux.Router {
	router := new()

	authRoutes(router)

	return router
}

func InitRoutesAuth() *mux.Router {
	router := new()

	taskRoutes(router)

	return router
}

func new() *mux.Router {
	router := mux.NewRouter()
	router.NotFoundHandler = infra.MethodNotFoundHandler()
	router.MethodNotAllowedHandler = infra.MethodNotAllowedHandler()
	router.Use(middleware.BindDb)
	return router
}
