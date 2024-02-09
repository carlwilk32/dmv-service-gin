package api

import (
	"encoding/json"
	dmv "github.com/carlwilk32/dmv-service-gin/client"
	"net/http"
)

// todo just a sample implementqtion for now
func Test(w http.ResponseWriter, r *http.Request) {
	offices := dmv.GetFieldOffices()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(offices[0]) //write one
	if err != nil {
		panic(err)
	}
}
