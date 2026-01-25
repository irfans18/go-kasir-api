package main

import (
	"fmt"
	"kasir-api/server"
	"net/http"
)

func main() {
	// Server Config
	PORT := "8080"
	HOST := "localhost"

	mux := server.NewRouter()

	fmt.Println("Server running di", fmt.Sprintf("%s:%s", HOST, PORT))

	err := http.ListenAndServe(fmt.Sprintf("%s:%s", HOST, PORT), mux)
	if err != nil {
		fmt.Println("Gagal running server")
	}
}
