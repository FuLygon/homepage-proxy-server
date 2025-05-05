package services

import (
	"encoding/json"
	"fmt"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/models"
	"net/http"
	"time"
)

type LinkwardenService interface {
	GetCollections() (*models.LinkwardenStatsResponse, error)
	GetTags() (map[string]interface{}, error)
}

type linkwardenService struct {
	client  *http.Client
	baseUrl string
	apiKey  string
}

func NewLinkwardenService(serviceConfig config.ServicesConfig) LinkwardenService {
	baseConfig := serviceConfig.Linkwarden
	return &linkwardenService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl: baseConfig.Url,
		apiKey:  baseConfig.Key,
	}
}

// GetCollections implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/linkwarden/component.jsx
func (s *linkwardenService) GetCollections() (*models.LinkwardenStatsResponse, error) {
	// Prepare collections request
	collectionsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/collections", s.baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare collections request: %w", err)
	}
	collectionsReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

	// Make collections request
	collectionsResp, err := s.client.Do(collectionsReq)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch collections: %w", err)
	}
	defer collectionsResp.Body.Close()

	// Return error if status code is not 200
	if collectionsResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch collections with status: %s", collectionsResp.Status)
	}

	// Parse collections response
	var response models.LinkwardenStatsResponse
	if err = json.NewDecoder(collectionsResp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to parse collections response: %w", err)
	}

	return &response, nil
}

// GetTags implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/linkwarden/component.jsx
func (s *linkwardenService) GetTags() (map[string]interface{}, error) {
	// Prepare tags request
	tagsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/v1/tags", s.baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare tags request: %w", err)
	}
	tagsReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", s.apiKey))

	// Make tags request
	tagsResp, err := s.client.Do(tagsReq)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tags: %w", err)
	}
	defer tagsResp.Body.Close()

	// Return error if status code is not 200
	if tagsResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch tags with status: %s", tagsResp.Status)
	}

	// Parse tags response
	var tagsStats models.LinkwardenStatsResponse
	if err = json.NewDecoder(tagsResp.Body).Decode(&tagsStats); err != nil {
		return nil, fmt.Errorf("failed to parse tags response: %w", err)
	}

	// Create a fake response with the same length as tags responses
	messages := make([]struct{}, len(tagsStats.Response))
	response := make(map[string]interface{})
	response["response"] = messages

	return response, nil
}
