package services

import (
	"encoding/json"
	"fmt"
	"homepage-widgets-gateway/internal/models"
	"net/http"
	"time"
)

type PortainerService interface {
	GetStats(baseUrl, key string, env int) (*[]models.PortainerResponse, error)
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

// GetStats implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/portainer/component.jsx
func (s *portainerService) GetStats(baseUrl, key string, env int) (*[]models.PortainerResponse, error) {
	// Prepare stats request
	// Hardcoded query param all=1 instead of taking it from the request
	statsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/endpoints/%d/docker/containers/json?all=1", baseUrl, env), nil)
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

	// Return error if status code is not 200
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch stats with status: %s", resp.Status)
	}

	// Parse stats response
	var response []models.PortainerResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to parse stats response: %w", err)
	}

	return &response, nil
}
