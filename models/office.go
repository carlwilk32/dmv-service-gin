package models

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
