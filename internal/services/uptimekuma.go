package services

import (
	"encoding/json"
	"fmt"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/models"
	"net/http"
)

type UptimeKumaService interface {
	GetStats(slug string) (*models.UptimeKumaStatsResponse, error)
	GetStatsHeartbeat(slug string) (*models.UptimeKumaStatsHeartbeatResponse, error)
}

type uptimeKumaService struct {
	baseUrl string
}

func NewUptimeKumaService(serviceConfig config.ServicesConfig) UptimeKumaService {
	baseConfig := serviceConfig.UptimeKuma
	return &uptimeKumaService{
		baseUrl: baseConfig.Url,
	}
}

// GetStats implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/uptimekuma/component.jsx
func (s *uptimeKumaService) GetStats(slug string) (*models.UptimeKumaStatsResponse, error) {
	// Get stats data
	statsUrl := fmt.Sprintf("%s/api/status-page/%s", s.baseUrl, slug)
	statsResp, err := http.Get(statsUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch stats: %w", err)
	}
	defer statsResp.Body.Close()

	// Return error if status code is not 200
	if statsResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch stats with status: %s", statsResp.Status)
	}

	// Parse stats response
	var response models.UptimeKumaStatsResponse
	if err = json.NewDecoder(statsResp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode stats response: %w", err)
	}

	return &response, nil
}

// GetStatsHeartbeat implement from https://github.com/gethomepage/homepage/blob/main/src/widgets/uptimekuma/component.jsx
func (s *uptimeKumaService) GetStatsHeartbeat(slug string) (*models.UptimeKumaStatsHeartbeatResponse, error) {
	// Get heartbeat data
	heartbeatUrl := fmt.Sprintf("%s/api/status-page/heartbeat/%s", s.baseUrl, slug)
	heartbeatResp, err := http.Get(heartbeatUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch heartbeat stats: %w", err)
	}
	defer heartbeatResp.Body.Close()

	// Return error if status code is not 200
	if heartbeatResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch heartbeat stats with status: %s", heartbeatResp.Status)
	}

	// Parse heartbeat stats response
	var response models.UptimeKumaStatsHeartbeatResponse
	if err = json.NewDecoder(heartbeatResp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf("failed to decode heartbeat stats response: %w", err)
	}

	return &response, nil
}
