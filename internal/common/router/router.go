package router

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

type Router struct {
	mux    *mux.Router
	routes []Route
}

func NewRouter(routes []Route) (Router, error) {
	m := mux.NewRouter()

	return Router{
		mux:    m,
		routes: routes,
	}, nil
}

func (r *Router) Serve(port string) error {
	if port[:1] != ":" {
		return errors.New("incorrect port")
	}

	for _, route := range r.routes {
		r.mux.HandleFunc(route.path, route.handler).Methods(route.method)
	}

	corsHandler := cors.AllowAll().Handler(r.mux)
	err := http.ListenAndServe(port, corsHandler)

	if err != nil {
		return err
	}

	return nil
}

func Variables(r *http.Request) map[string]string {
	return mux.Vars(r)
}
