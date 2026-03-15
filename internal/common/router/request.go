package router

import (
	"encoding/json"
	"net/http"
)

func ParseRequest[T any](r *http.Request, entity *T) error {
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(entity)
	if err != nil {
		return err
	}

	return nil
}
