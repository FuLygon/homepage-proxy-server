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
	GetMessages(baseUrl, key string) (map[string]interface{}, error)
	GetApplications(baseUrl, key string) (interface{}, error)
	GetClients(baseUrl, key string) (interface{}, error)
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

type dummy struct{}

// GetApplications implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/gotify/component.jsx
func (s *gotifyService) GetApplications(baseUrl, key string) (interface{}, error) {
	// Prepare stats request
	applicationStatsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/application", baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare application stats request: %w", err)
	}

	applicationStatsReq.Header.Add("X-Gotify-Key", key)

	// Make stats request
	resp, err := s.client.Do(applicationStatsReq)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch application stats: %w", err)
	}
	defer resp.Body.Close()

	// Parse stats response
	var applicationsStats []map[string]interface{}
	if err = json.NewDecoder(resp.Body).Decode(&applicationsStats); err != nil {
		return nil, fmt.Errorf("failed to parse application stats response: %w", err)
	}

	// Create a fake response with the same length as applicationsStats
	response := make([]dummy, len(applicationsStats))
	return response, nil
}

// GetClients implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/gotify/component.jsx
func (s *gotifyService) GetClients(baseUrl, key string) (interface{}, error) {
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

	// Create a fake response with the same length as clientsStats
	response := make([]dummy, len(clientsStats))
	return response, nil
}

// GetMessages partially implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/gotify/component.jsx
// Because the current implementation by Homepage has an issue where messages are capped at 100
func (s *gotifyService) GetMessages(baseUrl, key string) (map[string]interface{}, error) {
	var (
		totalMessages int
		offset        int
	)
	for {
		size, since, err := func() (int, int, error) {
			// Prepare stats request
			reqUrl, err := url.Parse(fmt.Sprintf("%s/message", baseUrl))
			if err != nil {
				log.Fatal(err)
			}

			queryParams := reqUrl.Query()
			queryParams.Set("limit", "200") // Limitation by Gotify API
			queryParams.Set("since", fmt.Sprintf("%d", offset))
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
		}()
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

	// Create a fake response with the same length as totalMessages
	messages := make([]dummy, totalMessages)
	response := make(map[string]interface{})
	response["messages"] = messages

	return response, nil
}
