package client

import (
	"fmt"
	"strconv"
)

type Branch struct {
	PublicID     string `json:"publicId,omitempty"`
	Name         string `json:"name,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressZip   string `json:"addressZip,omitempty"`
}

type Availability struct {
	PublicID     string `json:"publicId,omitempty"`
	Name         string `json:"name,omitempty"`
	AddressLine1 string `json:"addressLine1,omitempty"`
	AddressZip   string `json:"addressZip,omitempty"`
}

type DmvService struct {
	PublicID      string `json:"publicId,omitempty"`
	Name          string `json:"name,omitempty"`
	Active        bool   `json:"active,omitempty"`
	PublicEnabled bool   `json:"publicEnabled,omitempty"`
}

type AuthCheck struct {
	User      string `json:"user,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	UID       bool   `json:"uid,omitempty"`
}

type Appointment struct {
	NumberItems           int    `json:"numberItems"`
	OfficeID              string `json:"officeId"`
	RequestDate           string `json:"requestDate"`
	RequestTime           string `json:"requestTime"`
	RequestTask           string `json:"requestTask"`
	FirstName             string `json:"firstName"`
	LastName              string `json:"lastName"`
	TelNumber             string `json:"telNumber"`
	Token                 string `json:"token"`
	HasPrevPermit         bool   `json:"hasPrevPermit"`
	Dob                   string `json:"dob"`
	DlNumber              string `json:"dlNumber"`
	SafetyCourseCompleted string `json:"safetyCourseCompleted"`
}

type AppointmentResponse struct {
	FindApptResponse struct {
		Cheader []any `json:"cheader,omitempty"`
		Status  struct {
			Success string `json:"bsuccess,omitempty"`
		} `json:"bstatus,omitempty"`
		Payload struct {
			Hold          []any `json:"dhold,omitempty"`
			Office        []any `json:"doffice,omitempty"`
			FoaStatusCode struct {
				StatusCode string `json:"dstatusCode,omitempty"`
			} `json:"dfoaStatusCode,omitempty"`
		} `json:"dpayload,omitempty"`
	} `json:"dFindApptResponse,omitempty"`
}

type FieldOffice struct {
	ID       int    `json:"id,omitempty"`
	Slug     string `json:"slug,omitempty"`
	Link     string `json:"link,omitempty"`
	Title    Title  `json:"title,omitempty"`
	Meta     Meta   `json:"meta,omitempty"`
	Distance bool   `json:"distance,omitempty"`
}

func (o FieldOffice) LatLon() (lat, lon float64) {
	meta := o.Meta
	lat, _ = strconv.ParseFloat(meta.Latitude, 64)
	lon, _ = strconv.ParseFloat(meta.Longitude, 64)
	return
}

func (o FieldOffice) String() string {
	return fmt.Sprintf("DMV %v", o.Title.Rendered)
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
	PublicID    string              `json:"dmv_field_office_public_id,omitempty"`
	Appointment string              `json:"dmv_field_office_appointment,omitempty"`
	Street      string              `json:"dmv_field_office_street,omitempty"`
	City        string              `json:"dmv_field_office_city,omitempty"`
	Zipcode     string              `json:"dmv_field_office_zipcode,omitempty"`
	Latitude    string              `json:"dmv_field_office_latitude,omitempty"`
	Longitude   string              `json:"dmv_field_office_longitude,omitempty"`
	Hours       DmvFieldOfficeHours `json:"dmv_field_office_hours,omitempty"`
	Phone       string              `json:"dmv_field_office_phone,omitempty"`
}
