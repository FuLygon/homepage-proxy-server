package models

// PortainerStats lists of containers and stats
type PortainerStats struct {
	State  string `json:"State"`
	Status string `json:"Status"`
}

// PortainerResponse minimalized response from Portainer API
type PortainerResponse struct {
	Total   int `json:"total"`
	Running int `json:"running"`
	Stopped int `json:"stopped"`
}
