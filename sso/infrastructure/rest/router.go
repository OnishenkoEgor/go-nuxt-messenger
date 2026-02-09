package rest

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"sso/application"
)

type Router struct {
	app       application.Application
	muxRouter *mux.Router
}

func NewRouter(r *mux.Router, app application.Application) Router {
	return Router{
		app:       app,
		muxRouter: r,
	}
}

func (r Router) Serve(port string) error {
	if port[:1] != ":" {
		return errors.New("incorrect port")
	}

	corsHandler := cors.AllowAll().Handler(r.muxRouter)
	err := http.ListenAndServe(port, corsHandler)

	if err != nil {
		return err
	}

	return nil
}

func (r Router) initRoutes() error {
	routes := NewRoutes(r.app)

	for _, route := range routes {
		r.muxRouter.HandleFunc(route.path, route.handler).Methods(route.method)
	}

	return nil
}
