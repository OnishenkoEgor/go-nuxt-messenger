package controller

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Errors bool `json:"errors"`
	Data   any  `json:"data"`
}

func ReadRequest[T any](r *http.Request, entity *T) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(entity)
	if err != nil {
		return err
	}

	return nil
}

func WriteResponse[T any | []T](errors bool, data T, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	if errors {
		w.WriteHeader(http.StatusInternalServerError)

	} else {
		w.WriteHeader(http.StatusOK)
	}

	err := json.NewEncoder(w).Encode(Response{
		Errors: errors,
		Data:   data,
	})

	if err != nil {
		panic(err)
	}
}
