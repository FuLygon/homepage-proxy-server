package services

import (
	"encoding/json"
	"fmt"
	"homepage-proxy-server/internal/models"
	"net/http"
	"time"
)

type AdGuardHomeService interface {
	GetStats(baseUrl, username, password string) (*models.AdGuardHomeResponse, error)
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

func (s *adGuardHomeService) GetStats(baseUrl, username, password string) (*models.AdGuardHomeResponse, error) {
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
	var stats models.AdGuardHomeStats
	if err := json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("failed to parse stats response: %w", err)
	}

	// Map response
	response := &models.AdGuardHomeResponse{
		Queries:  stats.NumDNSQueries,
		Blocked:  stats.NumBlockedFiltering,
		Filtered: stats.NumReplacedSafeBrowsing + stats.NumReplacedSafeSearch + stats.NumReplacedParental,
		Latency:  stats.AvgProcessingTime * 1000, // Convert to milliseconds
	}

	return response, nil
}
