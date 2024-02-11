package app

import (
	"net/http"
)

func addRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /distance", ByDistance)
}
