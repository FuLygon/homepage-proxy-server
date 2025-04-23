package services

import (
	"encoding/json"
	"fmt"
	"homepage-widgets-gateway/internal/models"
	"net/http"
	"time"
)

type PortainerService interface {
	GetStats(baseUrl, key string, env int) (*models.PortainerResponse, error)
}

type portainerService struct {
	client *http.Client
}

func NewPortainerService() PortainerService {
	return &portainerService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
func (s *portainerService) GetStats(baseUrl, key string, env int) (*models.PortainerResponse, error) {
	// Prepare stats request
	statsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/endpoints/%d/docker/containers/json?all=true", baseUrl, env), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare stats request: %w", err)
	}

	statsReq.Header.Add("X-API-Key", key)

	// Make stats request
	resp, err := s.client.Do(statsReq)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stats: %w", err)
	}
	defer resp.Body.Close()

	// Parse stats response
	var stats []models.PortainerStats
	if err = json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("failed to parse stats response: %w", err)
	}

	response := &models.PortainerResponse{
		Total: len(stats),
	}

	for _, container := range stats {
		switch container.State {
		case "running":
			response.Running++
		case "exited", "dead":
			response.Stopped++
		}
	}

	return response, nil
}
