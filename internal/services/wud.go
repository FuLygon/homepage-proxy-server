package services

import (
	"encoding/json"
	"fmt"
	"homepage-widgets-gateway/internal/models"
	"net/http"
	"time"
)

type WUDService interface {
	GetStats(baseUrl, username, password string) (*models.WUDResponse, error)
}

type wudService struct {
	client *http.Client
}

func NewWUDService() WUDService {
	return &wudService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
func (s *wudService) GetStats(baseUrl, username, password string) (*models.WUDResponse, error) {
	// Prepare stats request
	statsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/containers", baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare stats request: %w", err)
	}
	statsReq.SetBasicAuth(username, password)

	// Make stats request
	resp, err := s.client.Do(statsReq)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stats: %w", err)
	}
	defer resp.Body.Close()

	// Parse stats response
	var stats []models.WUDStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("failed to parse stats response: %w", err)
	}

	response := &models.WUDResponse{
		Monitoring: len(stats),
	}

	for _, container := range stats {
		if container.UpdateAvailable {
			response.Updates++
		}
	}

	return response, nil
}
