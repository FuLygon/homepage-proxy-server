package services

import (
	"encoding/json"
	"fmt"
	"homepage-widgets-gateway/internal/models"
	"log"
	"net/http"
	"net/url"
	"time"
)

type GotifyService interface {
	GetStats(baseUrl, key string) (*models.GotifyResponse, error)
}

type gotifyService struct {
	client *http.Client
}

func NewGotifyService() GotifyService {
	return &gotifyService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
func (s *gotifyService) GetStats(baseUrl, key string) (*models.GotifyResponse, error) {
	totalApplications, err := s.getTotalApplications(baseUrl, key)
	if err != nil {
		return nil, fmt.Errorf("failed to get total applications: %w", err)
	}

	totalClients, err := s.getTotalClients(baseUrl, key)
	if err != nil {
		return nil, fmt.Errorf("failed to get total clients: %w", err)
	}

	var (
		totalMessages int
		offset        int
	)
	for {
		size, since, err := s.getMessages(baseUrl, key, offset)
		if err != nil {
			return nil, fmt.Errorf("failed to get total messages: %w", err)
		}

		totalMessages += size
		if since == 0 {
			break
		} else {
			offset = since
		}
	}

	response := &models.GotifyResponse{
		Applications: totalApplications,
		Clients:      totalClients,
		Messages:     totalMessages,
	}

	return response, nil
}

func (s *gotifyService) getTotalApplications(baseUrl, key string) (int, error) {
	// Prepare stats request
	applicationStatsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/application", baseUrl), nil)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare application stats request: %w", err)
	}

	applicationStatsReq.Header.Add("X-Gotify-Key", key)

	// Make stats request
	resp, err := s.client.Do(applicationStatsReq)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch application stats: %w", err)
	}
	defer resp.Body.Close()

	// Parse stats response
	var applicationsStats []map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&applicationsStats); err != nil {
		return 0, fmt.Errorf("failed to parse application stats response: %w", err)
	}

	return len(applicationsStats), nil
}

func (s *gotifyService) getTotalClients(baseUrl, key string) (int, error) {
	// Prepare stats request
	clientStatsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/client", baseUrl), nil)
	if err != nil {
		return 0, fmt.Errorf("failed to prepare client stats request: %w", err)
	}

	clientStatsReq.Header.Add("X-Gotify-Key", key)

	// Make stats request
	resp, err := s.client.Do(clientStatsReq)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch client stats: %w", err)
	}
	defer resp.Body.Close()

	// Parse stats response
	var clientsStats []map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&clientsStats); err != nil {
		return 0, fmt.Errorf("failed to parse client stats response: %w", err)
	}

	return len(clientsStats), nil
}

func (s *gotifyService) getMessages(baseUrl, key string, since int) (int, int, error) {
	// Prepare stats request
	reqUrl, err := url.Parse(fmt.Sprintf("%s/message", baseUrl))
	if err != nil {
		log.Fatal(err)
	}

	queryParams := reqUrl.Query()
	queryParams.Set("limit", "200")
	queryParams.Set("since", fmt.Sprintf("%d", since))
	reqUrl.RawQuery = queryParams.Encode()

	clientStatsReq, err := http.NewRequest("GET", reqUrl.String(), nil)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to prepare message stats request: %w", err)
	}

	clientStatsReq.Header.Add("X-Gotify-Key", key)

	// Make stats request
	resp, err := s.client.Do(clientStatsReq)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to fetch message stats: %w", err)
	}
	defer resp.Body.Close()

	// Parse stats response
	var messageStats models.GotifyMessageStats
	if err = json.NewDecoder(resp.Body).Decode(&messageStats); err != nil {
		return 0, 0, fmt.Errorf("failed to parse message stats response: %w", err)
	}

	return messageStats.Paging.Size, messageStats.Paging.Since, nil
}
