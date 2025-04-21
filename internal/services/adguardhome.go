package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"homepage-proxy-server/internal/models"
	"net/http"
	"time"
)

type AdGuardHomeService struct {
	client *http.Client
}

func NewAdGuardHomeService() *AdGuardHomeService {
	return &AdGuardHomeService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetStats request and maps the AdGuard Home response
func (s *AdGuardHomeService) GetStats(baseUrl, username, password string) (*models.AdGuardHomeResponse, error) {
	// Get agh_session cookie
	sessionCookie, err := s.login(baseUrl, username, password)
	if err != nil {
		return nil, fmt.Errorf("failed to login: %w", err)
	}

	// Prepare stats request
	statsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/control/stats", baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare stats request: %w", err)
	}

	// Add cookie to the request
	statsReq.AddCookie(&http.Cookie{
		Name:  "agh_session",
		Value: sessionCookie,
	})

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

// login performs login request and return agh_session cookie for stats request
func (s *AdGuardHomeService) login(baseUrl, username, password string) (string, error) {
	loginReq := models.AdGuardHomeLoginRequest{
		Name:     username,
		Password: password,
	}

	// Login request JSON
	loginJSON, err := json.Marshal(loginReq)
	if err != nil {
		return "", fmt.Errorf("failed to prepare login request: %w", err)
	}

	// Prepare login request
	loginResp, err := s.client.Post(fmt.Sprintf("%s/control/login", baseUrl), "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		return "", fmt.Errorf("failed to prepare login request: %w", err)
	}
	defer loginResp.Body.Close()

	if loginResp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("login failed with status: %d", loginResp.StatusCode)
	}

	// Retrieve agh_session cookie
	for _, cookie := range loginResp.Cookies() {
		if cookie.Name == "agh_session" {
			return cookie.Value, nil
		}
	}

	return "", fmt.Errorf("unable to retrieve agh_session cookie")
}
