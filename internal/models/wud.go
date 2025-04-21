package models

type WUDStats struct {
	UpdateAvailable bool `json:"updateAvailable"`
}

type WUDResponse struct {
	Monitoring int `json:"monitoring"`
	Updates    int `json:"updates"`
}
