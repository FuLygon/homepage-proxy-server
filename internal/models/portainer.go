package models

type PortainerStats struct {
	State  string `json:"State"`
	Status string `json:"Status"`
}

type PortainerResponse struct {
	Total   int `json:"total"`
	Running int `json:"running"`
	Stopped int `json:"stopped"`
}
