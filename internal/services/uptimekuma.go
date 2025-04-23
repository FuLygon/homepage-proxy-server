package services

import (
	"encoding/json"
	"fmt"
	"homepage-widgets-gateway/internal/models"
	"net/http"
)

type UptimeKumaService interface {
	GetStats(baseUrl, slug string) (*models.UptimeKumaStatsResponse, error)
	GetStatsHeartbeat(baseUrl, slug string) (*models.UptimeKumaStatsHeartbeatResponse, error)
}

type uptimeKumaService struct{}

func NewUptimeKumaService() UptimeKumaService {
	return &uptimeKumaService{}
}

// GetStats implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/uptimekuma/component.jsx
func (s *uptimeKumaService) GetStats(baseUrl, slug string) (*models.UptimeKumaStatsResponse, error) {
	// Get stats data
	statsUrl := fmt.Sprintf("%s/api/status-page/%s", baseUrl, slug)
	statsResp, err := http.Get(statsUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stats: %w", err)
	}
	defer statsResp.Body.Close()

	// Parse stats response
	var response models.UptimeKumaStatsResponse
	if err = json.NewDecoder(statsResp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode stats response: %w", err)
	}

	return &response, nil
}

// GetStatsHeartbeat implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/uptimekuma/component.jsx
func (s *uptimeKumaService) GetStatsHeartbeat(baseUrl, slug string) (*models.UptimeKumaStatsHeartbeatResponse, error) {
	// Get heartbeat data
	heartbeatUrl := fmt.Sprintf("%s/api/status-page/heartbeat/%s", baseUrl, slug)
	heartbeatResp, err := http.Get(heartbeatUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch heartbeat stats: %w", err)
	}
	defer heartbeatResp.Body.Close()

	// Parse heartbeat stats response
	var response models.UptimeKumaStatsHeartbeatResponse
	if err = json.NewDecoder(heartbeatResp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode heartbeat stats response: %w", err)
	}

	return &response, nil
}
