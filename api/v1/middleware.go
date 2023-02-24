package v1

import (
	"encoding/json"
	"net/http"
)

func writeJSON(w http.ResponseWriter, status int, v any) error {
	/*
		Response writer middleware
	*/
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func makeHttpHandler(f ApiFunc) http.HandlerFunc {
	/*
		Middleware that returns HandleFunc
	*/
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			return
		}
	}
}
