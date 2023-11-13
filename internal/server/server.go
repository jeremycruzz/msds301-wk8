package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func Start(router *chi.Mux, port string) {

	fmt.Printf("Starting server on port:%s\n", port)
	err := http.ListenAndServe(port, router)
	if err != nil {
		log.Fatalf("Error starting server: %s\n", err)
	}
}
