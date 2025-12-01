package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Auth struct {
	email    string
	password string
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	fmt.Println(r.Body)
	var auth Auth
	err := decoder.Decode(&auth)
	fmt.Println(auth)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("foo"))
}
