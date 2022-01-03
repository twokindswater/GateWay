package model

type Account struct {
	Id        string   `json:"id"`
	Name      string   `json:"name"`
	Image     string   `json:"image"`
	SSID      string   `json:"ssid"`
	BSSID     string   `json:"bssid"`
	Street    string   `json:"street"`
	InitDate  string   `json:"initDate"`
	Latitude  float64  `json:"latitude"`
	Longitude float64  `json:"longitude"`
	Friends   []string `json:"friends"`
	AtHome    bool     `json:"atHome"`
	Token     string   `json:"token"`
}

type Friend struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Image  string `json:"image"`
	AtHome bool   `json:"atHome"`
}
