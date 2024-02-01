package handlers

import (
	"fmt"
	"net/http"
)

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
