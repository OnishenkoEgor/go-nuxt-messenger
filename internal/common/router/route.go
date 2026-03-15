package router

import "net/http"

type Handler func(w http.ResponseWriter, r *http.Request)

type Route struct {
	path    string
	handler Handler
	method  string
}

func NewRoute(path string, handler Handler, method string) Route {
	return Route{
		path:    path,
		handler: handler,
		method:  method,
	}
}
