package data

type AccountInfo struct {
	Id       string   `json:"id"`
	Image    string   `json:"image"`
	SSID     string   `json:"ssid"`
	BSSID    string   `json:"bssid"`
	TimeInfo TimeInfo `json:"time_info"`
}

type TimeInfo struct {
	Total int
	Week  int
	Day   int
}
