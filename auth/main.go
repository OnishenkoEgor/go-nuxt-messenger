package main

import (
	"encoding/json"
	_ "fmt"
)
import "net/http"

type Auth struct {
	email    string
	password string
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var auth Auth

	err := decoder.Decode(&auth)

	if err != nil {
		panic(err)
	}

}

func main() {
	http.HandleFunc("/api/login", loginHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		return
	}
}
