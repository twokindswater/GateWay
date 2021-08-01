package data

type AccountInfo struct {
	Id        string  `json:"id"`
	Name      string  `json:"name"`
	Image     string  `json:"image"`
	SSID      string  `json:"ssid"`
	BSSID     string  `json:"bssid"`
	Street    string  `json:"street"`
	InitDate  int     `json:"initDate"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
