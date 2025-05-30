package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/cache"
	"homepage-widgets-gateway/internal/models"
	"net/http"
	"time"
)

const npmAuthTokenCacheKey = "npm_auth_token"
const npmAuthTokenExpiry = "npm_auth_expiry"

type NPMService interface {
	GetStats(authToken string) (*[]models.NPMStatsResponse, error)
	Login() (*models.NPMAuthResponse, error)
}

type npmService struct {
	client   *http.Client
	cache    cache.Cache
	baseUrl  string
	username string
	password string
}

func NewNPMService(serviceConfig config.ServicesConfig, cache cache.Cache) NPMService {
	baseConfig := serviceConfig.NginxProxyManager
	return &npmService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		cache:    cache,
		baseUrl:  baseConfig.Url,
		username: baseConfig.Username,
		password: baseConfig.Password,
	}
}

// GetStats implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/npm/component.jsx
func (s *npmService) GetStats(authToken string) (*[]models.NPMStatsResponse, error) {
	// Prepare stats request
	statsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/nginx/proxy-hosts", s.baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare stats request: %w", err)
	}
	statsReq.Header.Add("Authorization", authToken)

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
	var stats []models.NPMStatsResponse
	if err = json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("failed to parse stats response: %w", err)
	}

	return &stats, nil
}

// Login implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/npm/proxy.js
func (s *npmService) Login() (*models.NPMAuthResponse, error) {
	var npmAuthResponse models.NPMAuthResponse

	// Attempt to retrieve auth token and expiry from cache
	// Notes that Homepage also does caching for auth on their side, but we cache it on our side anyway in case Homepage was restarted
	npmAuthTokenCache, foundToken := s.cache.Get(npmAuthTokenCacheKey)
	npmAuthExpiryCache, foundExpiry := s.cache.Get(npmAuthTokenExpiry)

	if foundToken && foundExpiry {
		npmAuthResponse.Token = npmAuthTokenCache.(string)
		npmAuthResponse.Expires = npmAuthExpiryCache.(time.Time)
	} else {
		loginReq := models.NPMAuthRequest{
			Identity: s.username,
			Secret:   s.password,
		}

		// Login request JSON
		loginJSON, err := json.Marshal(loginReq)
		if err != nil {
			return &npmAuthResponse, fmt.Errorf("failed to prepare login request: %w", err)
		}

		// Prepare login request
		loginResp, err := s.client.Post(fmt.Sprintf("%s/api/tokens", s.baseUrl), "application/json", bytes.NewBuffer(loginJSON))
		if err != nil {
			return &npmAuthResponse, fmt.Errorf("failed to prepare login request: %w", err)
		}
		defer loginResp.Body.Close()

		if loginResp.StatusCode != http.StatusOK {
			return &npmAuthResponse, fmt.Errorf("login failed with status: %d", loginResp.StatusCode)
		}

		if err = json.NewDecoder(loginResp.Body).Decode(&npmAuthResponse); err != nil {
			return &npmAuthResponse, fmt.Errorf("failed to parse login response: %w", err)
		}

		// Get ttl for cache
		ttl := time.Until(npmAuthResponse.Expires)
		if ttl <= 0 {
			// Default ttl if there is any issue with expiration time
			ttl = 1 * time.Hour
		}

		// Cache token and expiry
		s.cache.Set(npmAuthTokenCacheKey, npmAuthResponse.Token, ttl)
		s.cache.Set(npmAuthTokenExpiry, npmAuthResponse.Expires, ttl)
	}

	return &npmAuthResponse, nil
}
