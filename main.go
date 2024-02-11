package main

import (
	"github.com/carlwilk32/dmv-service-gin/app"
	"log"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", app.Server())
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
