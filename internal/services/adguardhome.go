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

const aghSessionCacheKey = "agh_session"

type AdGuardHomeService interface {
	GetStats(baseUrl, username, password string) (*models.AdGuardHomeResponse, error)
}

type adGuardHomeService struct {
	client *http.Client
	cache  cache.Cache
}

func NewAdGuardHomeService(cache cache.Cache) AdGuardHomeService {
	return &adGuardHomeService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		cache: cache,
	}
}

func (s *adGuardHomeService) GetStats(baseUrl, username, password string) (*models.AdGuardHomeResponse, error) {
	var aghSession string
	// Attempt to retrieve session from cache
	aghSessionCache, found := s.cache.Get(aghSessionCacheKey)
	if found {
		aghSession = aghSessionCache.(string)
	} else {
		// If cache key value not found, get aghSession from login
		aghSessionCookie, ttl, err := s.login(baseUrl, username, password)
		if err != nil {
			return nil, fmt.Errorf("failed to login: %w", err)
		}
		// cache agh_session cookie
		s.cache.Set(aghSessionCacheKey, aghSessionCookie, ttl)
		aghSession = aghSessionCookie
	}

	// Prepare stats request
	statsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/control/stats", baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare stats request: %w", err)
	}

	// Add cookie to the request
	statsReq.AddCookie(&http.Cookie{
		Name:  "agh_session",
		Value: aghSession,
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

// login performs login request, return agh_session cookie for stats request and ttl for caching
func (s *adGuardHomeService) login(baseUrl, username, password string) (string, time.Duration, error) {
	loginReq := models.AdGuardHomeLoginRequest{
		Name:     username,
		Password: password,
	}

	// Login request JSON
	loginJSON, err := json.Marshal(loginReq)
	if err != nil {
		return "", 0, fmt.Errorf("failed to prepare login request: %w", err)
	}

	// Prepare login request
	loginResp, err := s.client.Post(fmt.Sprintf("%s/control/login", baseUrl), "application/json", bytes.NewBuffer(loginJSON))
	if err != nil {
		return "", 0, fmt.Errorf("failed to prepare login request: %w", err)
	}
	defer loginResp.Body.Close()

	if loginResp.StatusCode != http.StatusOK {
		return "", 0, fmt.Errorf("login failed with status: %d", loginResp.StatusCode)
	}

	// Retrieve agh_session cookie and its expiration time
	var sessionCookie string
	var expirationTime time.Time
	for _, cookie := range loginResp.Cookies() {
		if cookie.Name == "agh_session" {
			sessionCookie = cookie.Value
			expirationTime = cookie.Expires
			break
		}
	}

	if sessionCookie == "" {
		return "", 0, fmt.Errorf("unable to retrieve agh_session cookie")
	}

	// Get ttl for cache
	ttl := time.Until(expirationTime)
	if ttl <= 0 {
		// Default ttl if there is any issue with expiration time
		ttl = 1 * time.Hour
	}

	return sessionCookie, ttl, nil
}
