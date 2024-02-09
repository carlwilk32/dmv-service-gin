package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AllOffices []FieldOffice

type FieldOffice struct {
	ID       int    `json:"id,omitempty"`
	Slug     string `json:"slug,omitempty"`
	Link     string `json:"link,omitempty"`
	Title    Title  `json:"title,omitempty"`
	Meta     Meta   `json:"meta,omitempty"`
	Distance bool   `json:"distance,omitempty"`
}

type Title struct {
	Rendered string `json:"rendered,omitempty"`
}

type Sunday struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type Monday struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type Tuesday struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type Wednesday struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type Thursday struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type Friday struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type Saturday struct {
	Start string `json:"start,omitempty"`
	End   string `json:"end,omitempty"`
}

type DmvFieldOfficeHours struct {
	Sunday    Sunday    `json:"sunday,omitempty"`
	Monday    Monday    `json:"monday,omitempty"`
	Tuesday   Tuesday   `json:"tuesday,omitempty"`
	Wednesday Wednesday `json:"wednesday,omitempty"`
	Thursday  Thursday  `json:"thursday,omitempty"`
	Friday    Friday    `json:"friday,omitempty"`
	Saturday  Saturday  `json:"saturday,omitempty"`
}

type Meta struct {
	DmvFieldOfficePublicID    string              `json:"dmv_field_office_public_id,omitempty"`
	DmvFieldOfficeAppointment string              `json:"dmv_field_office_appointment,omitempty"`
	DmvFieldOfficeStreet      string              `json:"dmv_field_office_street,omitempty"`
	DmvFieldOfficeCity        string              `json:"dmv_field_office_city,omitempty"`
	DmvFieldOfficeZipcode     string              `json:"dmv_field_office_zipcode,omitempty"`
	DmvFieldOfficeLatitude    string              `json:"dmv_field_office_latitude,omitempty"`
	DmvFieldOfficeLongitude   string              `json:"dmv_field_office_longitude,omitempty"`
	DmvFieldOfficeHours       DmvFieldOfficeHours `json:"dmv_field_office_hours,omitempty"`
	DmvFieldOfficePhone       string              `json:"dmv_field_office_phone,omitempty"`
}

func getFieldOffices(c *gin.Context) {
	req, _ := http.NewRequest("GET", "https://www.dmv.ca.gov/portal/wp-json/dmv/v1/field-offices", nil)
	req.Header.Set("user-agent", "curl/8.1.2")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		panic(err)
	}

	var offices AllOffices
	err = json.NewDecoder(resp.Body).Decode(&offices)
	if err != nil {
		panic(err)
	}

	fmt.Println("total:", len(offices))

	c.Status(http.StatusOK)
	//c.JSON(200, offices)
}

func main() {
	router := gin.Default()
	router.GET("/offices", getFieldOffices)

	router.Run("localhost:8080")
}
