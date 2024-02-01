package main

import (
	"log"
	"net/http"
	"github.com/username/api-backend/handlers"
)

func main() {
	mux := http.NewServeMux()
	handlers.SetupHandlers(mux)

	log.Println("Starting HTTP server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
