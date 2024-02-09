package client

import (
	"encoding/json"
	"fmt"
	. "github.com/carlwilk32/dmv-service-gin/models"
	"net/http"
)

type Offices []FieldOffice

func GetFieldOffices() Offices {
	req, _ := http.NewRequest("GET", "https://www.dmv.ca.gov/portal/wp-json/dmv/v1/field-offices", nil)
	req.Header.Set("user-agent", "curl/8.1.2")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		panic(err)
	}
	defer resp.Body.Close()

	var offices Offices
	err = json.NewDecoder(resp.Body).Decode(&offices)
	if err != nil {
		panic(err)
	}
	fmt.Println("total:", len(offices))

	return offices
}
