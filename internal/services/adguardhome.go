package services

import (
	"encoding/json"
	"fmt"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/models"
	"net/http"
	"time"
)

type AdGuardHomeService interface {
	GetStats() (*models.AdguardHomeStatsResponse, error)
}

type adGuardHomeService struct {
	client   *http.Client
	baseUrl  string
	username string
	password string
}

func NewAdGuardHomeService(serviceConfig config.ServicesConfig) AdGuardHomeService {
	baseConfig := serviceConfig.AdGuardHome
	return &adGuardHomeService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl:  baseConfig.Url,
		username: baseConfig.Username,
		password: baseConfig.Password,
	}
}

// GetStats implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/adguard/component.jsx
func (s *adGuardHomeService) GetStats() (*models.AdguardHomeStatsResponse, error) {
	// Prepare stats request
	statsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/control/stats", s.baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare stats request: %w", err)
	}
	statsReq.SetBasicAuth(s.username, s.password)

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
	var stats models.AdguardHomeStatsResponse
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("failed to parse stats response: %w", err)
	}

	return &stats, nil
}
