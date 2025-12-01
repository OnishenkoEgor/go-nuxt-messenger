package main

import (
	"auth/handlers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	r := mux.NewRouter()

	initRoutes(r)

	handler := cors.AllowAll().Handler(r)
	err := http.ListenAndServe(":8080", handler)

	if err != nil {
		return
	}
}

func initRoutes(r *mux.Router) {
	r.HandleFunc("/api/login", handlers.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/test", handlers.TestHandler).Methods(http.MethodGet)
}
