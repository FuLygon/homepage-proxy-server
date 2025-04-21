package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"homepage-proxy-server/internal/cache"
	"homepage-proxy-server/internal/models"
	"net/http"
	"time"
)

const npmAuthTokenCacheKey = "npm_token"

type NPMService interface {
	GetStats(baseUrl, username, password string) (*models.NPMResponse, error)
}

type npmService struct {
	client *http.Client
	cache  cache.Cache
}

func NewNPMService(cache cache.Cache) NPMService {
	return &npmService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		cache: cache,
	}
}
func (s *npmService) GetStats(baseUrl, username, password string) (*models.NPMResponse, error) {
	var npmAuthToken string
	// Attempt to retrieve auth token from cache
	npmAuthTokenCache, found := s.cache.Get(npmAuthTokenCacheKey)
	if found {
		npmAuthToken = npmAuthTokenCache.(string)
	} else {
		// If cache key value not found, get aghSession from login
		npmAuthTokenResp, ttl, err := s.login(baseUrl, username, password)
		if err != nil {
			return nil, fmt.Errorf("failed to login: %w", err)
		}

		// cache agh_session cookie
		s.cache.Set(npmAuthTokenCacheKey, npmAuthTokenResp, ttl)
		npmAuthToken = npmAuthTokenResp

	}

	// Prepare stats request
	statsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/nginx/proxy-hosts", baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare stats request: %w", err)
	}

	statsReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", npmAuthToken))

	// Make stats request
	resp, err := s.client.Do(statsReq)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stats: %w", err)
	}
	defer resp.Body.Close()

	// Parse stats response
	var proxyHostsList []models.NPMProxyHostsStats
	if err := json.NewDecoder(resp.Body).Decode(&proxyHostsList); err != nil {
		return nil, fmt.Errorf("failed to parse stats response: %w", err)
	}

	response := &models.NPMResponse{
		Total: len(proxyHostsList),
	}

	for _, host := range proxyHostsList {
		if host.Enabled {
			response.Enabled++
		} else {
			response.Disabled++
		}
	}

	return response, nil
}

func (s *npmService) login(baseUrl, username, password string) (string, time.Duration, error) {
	loginReq := models.NpmAuthRequest{
		Identity: username,
		Secret:   password,
	}

	// Login request JSON
	loginJSON, err := json.Marshal(loginReq)
	if err != nil {
		return "", 0, fmt.Errorf("failed to prepare login request: %w", err)
	}

	// Prepare login request
	loginResp, err := s.client.Post(fmt.Sprintf("%s/api/tokens", baseUrl), "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		return "", 0, fmt.Errorf("failed to prepare login request: %w", err)
	}
	defer loginResp.Body.Close()

	if loginResp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("login failed with status: %d", loginResp.StatusCode)
	}

	var authResponse models.NpmAuthResponse
	if err = json.NewDecoder(loginResp.Body).Decode(&authResponse); err != nil {
		return "", 0, fmt.Errorf("failed to parse login response: %w", err)
	}

	// Get ttl for cache
	ttl := time.Until(authResponse.Expires)
	if ttl <= 0 {
		// Default ttl if there is any issue with expiration time
		ttl = 1 * time.Hour
	}

	return authResponse.Token, ttl, nil
}
