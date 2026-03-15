package infrastructure

import (
	"messenger/router"
	"messenger/sso/application"
	"messenger/sso/infrastructure/user/rest/controller"
	"net/http"
)

func NewRoutes(app application.Application) []router.Route {
	userController := controller.NewUserController(app)

	return []router.Route{
		router.NewRoute("/api/users", userController.GetAll, http.MethodGet),
		router.NewRoute("/api/users/{id}", userController.Get, http.MethodGet),
		router.NewRoute("/api/users/create", userController.Create, http.MethodPost),
		router.NewRoute("/api/users/{id}", userController.Update, http.MethodPut),
		router.NewRoute("/api/users/{id}", userController.Delete, http.MethodDelete),
	}
}
