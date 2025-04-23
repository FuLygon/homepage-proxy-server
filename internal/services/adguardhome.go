package services

import (
	"encoding/json"
	"fmt"
	"homepage-widgets-gateway/internal/models"
	"net/http"
	"time"
)

type AdGuardHomeService interface {
	GetStats(baseUrl, username, password string) (*models.AdguardHomeStatsResponse, error)
}

type adGuardHomeService struct {
	client *http.Client
}

func NewAdGuardHomeService() AdGuardHomeService {
	return &adGuardHomeService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetStats implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/adguard/component.jsx
func (s *adGuardHomeService) GetStats(baseUrl, username, password string) (*models.AdguardHomeStatsResponse, error) {
	// Prepare stats request
	statsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/control/stats", baseUrl), nil)
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
	var stats models.AdguardHomeStatsResponse
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("failed to parse stats response: %w", err)
	}

	return &stats, nil
}
