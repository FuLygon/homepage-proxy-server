package services

import (
	"encoding/json"
	"fmt"
	"homepage-proxy-server/internal/models"
	"net/http"
)

type UptimeKumaService interface {
	GetStats(baseUrl, slug string) (*models.UptimeKumaResponse, error)
}

type uptimeKumaService struct{}

func NewUptimeKumaService() UptimeKumaService {
	return &uptimeKumaService{}
}
func (s *uptimeKumaService) GetStats(baseUrl, slug string) (*models.UptimeKumaResponse, error) {
	// Get heartbeat data
	heartbeatURL := fmt.Sprintf("%s/api/status-page/heartbeat/%s", baseUrl, slug)
	heartbeatResp, err := http.Get(heartbeatURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch heartbeat stats: %w", err)
	}
	defer heartbeatResp.Body.Close()

	// Parse heartbeat stats response
	var heartbeatData models.UptimeKumaHeartbeatStats
	if err = json.NewDecoder(heartbeatResp.Body).Decode(&heartbeatData); err != nil {
		return nil, fmt.Errorf("failed to decode heartbeat stats response: %w", err)
	}

	response := &models.UptimeKumaResponse{}
	for _, siteList := range heartbeatData.HeartbeatList {
		if len(siteList) > 0 {
			lastHeartbeat := siteList[len(siteList)-1]
			if lastHeartbeat.Status == 1 {
				response.SitesUp++
			} else {
				response.SitesDown++
			}
		}
	}

	// Calculate uptime percentage
	if len(heartbeatData.UptimeList) > 0 {
		var sum float64
		for _, uptime := range heartbeatData.UptimeList {
			sum += uptime
		}
		response.Uptime = (sum / float64(len(heartbeatData.UptimeList))) * 100
	}

	return response, nil
}
