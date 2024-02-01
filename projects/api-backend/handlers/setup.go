package handlers

import "net/http"

func SetupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/healthcheck", healthcheckHandler)
	
	
}
