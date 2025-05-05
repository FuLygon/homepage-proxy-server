package services

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/errgroup"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/cache"
	"homepage-widgets-gateway/internal/models"
	"net/http"
	"net/url"
	"time"
)

type YourSpotifyService interface {
	GetStats(ctx context.Context, timeRange string) (*models.YourSpotifyResponse, error)
}

type yourSpotifyService struct {
	client  *http.Client
	cache   cache.Cache
	baseUrl string
	token   string
}

func NewYourSpotifyService(serviceConfig config.ServicesConfig, cache cache.Cache) YourSpotifyService {
	baseConfig := serviceConfig.YourSpotify
	return &yourSpotifyService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		cache:   cache,
		baseUrl: baseConfig.Url,
		token:   baseConfig.Token,
	}
}

const yourSpotifyResponseCacheKey = "your_spotify_response_%s"

func (s *yourSpotifyService) GetStats(ctx context.Context, timeRange string) (*models.YourSpotifyResponse, error) {
	// Return cached response if available
	if cachedResponse, found := s.cache.Get(fmt.Sprintf(yourSpotifyResponseCacheKey, timeRange)); found {
		response := cachedResponse.(models.YourSpotifyResponse)
		return &response, nil
	}

	startTime, endTime, err := s.getTimeRange(timeRange)
	if err != nil {
		return nil, err
	}

	var response models.YourSpotifyResponse
	g, ctx := errgroup.WithContext(ctx)
	// Get songs listened stats
	g.Go(func() error {
		songsListenedEndpoint, err := s.getRequestUrl(s.baseUrl, "/api/spotify/songs_per", s.token, startTime, endTime)
		if err != nil {
			return fmt.Errorf("failed to get songs listened endpoint: %w", err)
		}
		count, err := s.getSongsListened(ctx, songsListenedEndpoint)
		if err != nil {
			return err
		}
		response.SongsListened = count
		return nil
	})

	// Get time listened stats
	g.Go(func() error {
		timeListenedEndpoint, err := s.getRequestUrl(s.baseUrl, "/api/spotify/time_per", s.token, startTime, endTime)
		if err != nil {
			return fmt.Errorf("failed to get time listened endpoint: %w", err)
		}
		count, err := s.getTimeListened(ctx, timeListenedEndpoint)
		if err != nil {
			return err
		}
		response.TimeListened = count
		return nil
	})

	// Get artists listened stats
	g.Go(func() error {
		artistsListenedEndpoint, err := s.getRequestUrl(s.baseUrl, "/api/spotify/different_artists_per", s.token, startTime, endTime)
		if err != nil {
			return fmt.Errorf("failed to get artists listened endpoint: %w", err)
		}
		count, err := s.getArtistsListened(ctx, artistsListenedEndpoint)
		if err != nil {
			return err
		}
		response.ArtistsListened = count
		return nil
	})

	// Wait for goroutines
	if err = g.Wait(); err != nil {
		return nil, err
	}

	// Cache the response for 5 min, since Your Spotify also doesn't fetch new data regularly
	s.cache.Set(fmt.Sprintf(yourSpotifyResponseCacheKey, timeRange), response, 5*time.Minute)

	return &response, nil
}

func (s *yourSpotifyService) getSongsListened(ctx context.Context, reqUrl string) (int64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create song listened stats request: %w", err)
	}

	// Prepare song listened stats request
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch songs listened stats: %w", err)
	}
	defer resp.Body.Close()

	// Return error if status code is not 200
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch songs listened stats with status: %s", resp.Status)
	}

	// Parse song listened stats response
	var response []models.YourSpotifySongsResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("failed to decode songs listened stats response: %w", err)
	}

	if len(response) == 0 {
		return 0, nil
	}

	return response[0].Count, nil
}

func (s *yourSpotifyService) getTimeListened(ctx context.Context, reqUrl string) (int64, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create time listened stats request: %w", err)
	}

	// Prepare time listened stats request
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch time listened stats: %w", err)
	}
	defer resp.Body.Close()

	// Return error if status code is not 200
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch time listened stats with status: %s", resp.Status)
	}

	// Parse time listened stats response
	var response []models.YourSpotifyTimeResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("failed to decode time listened stats response: %w", err)
	}

	if len(response) == 0 {
		return 0, nil
	}

	return response[0].Count / 1000 / 60, nil
}

func (s *yourSpotifyService) getArtistsListened(ctx context.Context, reqUrl string) (int, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqUrl, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create artists listened stats request: %w", err)
	}

	// Prepare artists listened stats request
	resp, err := s.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch artists listened stats: %w", err)
	}
	defer resp.Body.Close()

	// Return error if status code is not 200
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("failed to fetch artists listened stats with status: %s", resp.Status)
	}

	// Parse artists listened stats response
	var response []models.YourSpotifyArtistsResponse
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return 0, fmt.Errorf("failed to decode artists listened stats response: %w", err)
	}

	if len(response) == 0 {
		return 0, nil
	}

	return len(response[0].Artists), nil
}

func (s *yourSpotifyService) getRequestUrl(baseUrl, path, token, startTime, endTime string) (string, error) {
	reqUrl, err := url.Parse(baseUrl + path)
	if err != nil {
		return "", err
	}

	queryParams := reqUrl.Query()
	queryParams.Set("start", startTime)
	queryParams.Set("end", endTime)
	queryParams.Set("token", token)
	queryParams.Set("timeSplit", "all")
	reqUrl.RawQuery = queryParams.Encode()

	return reqUrl.String(), nil
}

func (s *yourSpotifyService) getTimeRange(timeRange string) (string, string, error) {
	var startTime time.Time
	endTime := time.Now()

	switch timeRange {
	case "day":
		startTime = endTime.AddDate(0, 0, -1)
	case "week":
		startTime = endTime.AddDate(0, 0, -7)
	case "month":
		startTime = endTime.AddDate(0, -1, 0)
	case "year":
		startTime = endTime.AddDate(-1, 0, 0)
	case "all":
		startTime = time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	default:
		return "", "", fmt.Errorf("invalid time range: %s", timeRange)
	}

	return startTime.Format(time.RFC3339Nano), endTime.Format(time.RFC3339Nano), nil
}
