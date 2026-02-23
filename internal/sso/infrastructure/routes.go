package infrastructure

import (
	"net/http"
	"sso/application"
	"sso/infrastructure/user/rest/controller"
)

type Route struct {
	path    string
	handler func(w http.ResponseWriter, r *http.Request)
	method  string
}

func NewRoute(path string, handler func(w http.ResponseWriter, r *http.Request), method string) Route {
	return Route{
		path:    path,
		handler: handler,
		method:  method,
	}
}

type Routes []Route

func NewRoutes(app application.Application) Routes {
	userController := controller.NewUserController(app)

	return Routes{
		NewRoute("/api/users", userController.GetAll, http.MethodGet),
		NewRoute("/api/users/{id}", userController.Get, http.MethodGet),
		NewRoute("/api/users/create", userController.Create, http.MethodPost),
		NewRoute("/api/users/{id}", userController.Update, http.MethodPut),
		NewRoute("/api/users/{id}", userController.Delete, http.MethodDelete),
	}
}
