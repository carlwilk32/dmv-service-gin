package main

import (
	"github.com/carlwilk32/dmv-service-gin/api"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/test", api.Test)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
