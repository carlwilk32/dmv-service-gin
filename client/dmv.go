package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchDmvData[T any](req *http.Request, model *T) {
	req.Header.Set("user-agent", "curl/8.1.2")
	req.Header.Set("accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		panic(err)
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&model)
	if err != nil {
		panic(err)
	}
}

func GetFieldOffices(query string) (offices []FieldOffice) {
	req, _ := http.NewRequest("GET", "https://www.dmv.ca.gov/portal/wp-json/dmv/v1/field-offices", nil)

	if len(query) > 0 {
		q := req.URL.Query()
		q.Add("q", query)
		req.URL.RawQuery = q.Encode()
	}

	fetchDmvData(req, &offices)
	return
}

func GetServices() (services []DmvService) {
	req, _ := http.NewRequest("GET", "https://www.dmv.ca.gov/portal/wp-json/dmv/v1/appointment/services", nil)
	fetchDmvData(req, &services)
	return
}

func GetAvailable(service DmvService) (availabilityList []Availability) {
	req, _ := http.NewRequest("GET", "https://www.dmv.ca.gov/portal/wp-json/dmv/v1/appointment/available", nil)

	q := req.URL.Query()
	q.Add("services[]", service.PublicID)
	req.URL.RawQuery = q.Encode()

	fetchDmvData(req, &availabilityList)
	return
}

// GetDates ["2024-02-16T00:00:00", "2024-03-22T00:00:00"]
func GetDates(service DmvService, office FieldOffice) (dates []string) {
	url := fmt.Sprintf("https://www.dmv.ca.gov/portal/wp-json/dmv/v1/appointment/branches/%s/dates", office.Meta.PublicID)
	req, _ := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	q.Add("services[]", service.PublicID)
	q.Add("numberOfCustomers", "1")
	req.URL.RawQuery = q.Encode()

	fetchDmvData(req, &dates)
	return
}

// GetTimes 24h format, f.e. ["10:20","11:00","11:20","13:40","14:20","14:40","15:20","15:40"]
func GetTimes(service DmvService, office FieldOffice, date string) (times []string) {
	url := fmt.Sprintf("https://www.dmv.ca.gov/portal/wp-json/dmv/v1/appointment/branches/%s/times", office.Meta.PublicID)
	req, _ := http.NewRequest("GET", url, nil)

	q := req.URL.Query()
	q.Add("services[]", service.PublicID)
	q.Add("numberOfCustomers", "1")
	req.URL.RawQuery = q.Encode()

	fetchDmvData(req, &times)
	return
}

func GetBranches() (branches []Branch) {
	req, _ := http.NewRequest("GET", "https://www.dmv.ca.gov/portal/wp-json/dmv/v1/appointment/branches", nil)
	fetchDmvData(req, &branches)
	return
}

func GetAuthCheck() (auth AuthCheck) {
	req, _ := http.NewRequest("GET", "https://www.dmv.ca.gov/portal/wp-json/dmv/v1/auth-check", nil)
	fetchDmvData(req, &auth)
	return
}

func PostAppointment(appointment Appointment) (response AppointmentResponse) {
	body, err := json.Marshal(appointment)
	if err != nil {
		panic(err)
	}

	req, _ := http.NewRequest("POST", "https://www.dmv.ca.gov/portal/wp-json/dmv/v1/appointment/hold-appointment", bytes.NewBuffer(body))
	req.Header.Set("content-type", "application/json")

	fetchDmvData(req, &response)
	return
}
