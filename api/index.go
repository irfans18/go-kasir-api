package api

import (
	"kasir-api/server"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	mux := server.NewRouter()
	mux.ServeHTTP(w, r)
}
