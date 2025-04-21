package models

// WUDStats lists available container with update
type WUDStats struct {
	UpdateAvailable bool `json:"updateAvailable"`
}

// WUDResponse minimalized response from WUD API
type WUDResponse struct {
	Monitoring int `json:"monitoring"`
	Updates    int `json:"updates"`
}
