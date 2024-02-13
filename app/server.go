package app

import "net/http"

func NewServer() http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux)
	var handler http.Handler = mux
	//handler = someMiddleware(handler)
	return handler
}
