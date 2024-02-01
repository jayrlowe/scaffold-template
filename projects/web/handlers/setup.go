package handlers

import "net/http"

func SetupHandlers(mux *http.ServeMux) {
	mux.HandleFunc("/api", apiHandler)
	mux.HandleFunc("/healthcheck", healthcheckHandler)
	
	
	mux.HandleFunc("/", indexHandler)
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", fs))
	
}
