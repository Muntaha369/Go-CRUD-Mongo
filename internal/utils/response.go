package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	ErrorType string
	Status string
}

func WriteJson(w http.ResponseWriter, status int, data any) error {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)

}

func Genralerror(err error) Error {
	return Error{
		ErrorType: err.Error(),
		Status:    "Error",
	}
}