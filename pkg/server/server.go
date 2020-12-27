package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/benfdking/jen/pkg/url"
)

// Serve server on given port which exposes the default jwks
func Serve(port int) error {
	h := GetHandler()
	return http.ListenAndServe(":"+strconv.Itoa(port), h)
}

// GetHandler returns the handler for the server
func GetHandler() http.Handler {
	s := http.NewServeMux()
	s.HandleFunc("/a", handler("a"))
	s.HandleFunc("/b", handler("b"))
	s.HandleFunc("/c", handler("c"))
	return s
}

func handler(letter string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			jwks, _, err := url.ReturnJWKSAndPrivateKeyFromData(letter)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			jsonbuf, err := json.MarshalIndent(jwks, "", "  ")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(jsonbuf)
			return
		default:
			w.WriteHeader(http.StatusNotFound)
			return
		}
	}
}
