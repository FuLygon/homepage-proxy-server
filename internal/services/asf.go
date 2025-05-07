package services

import (
	"encoding/json"
	"fmt"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/models"
	"net/http"
	"time"
)

type ASFService interface {
	GetStats() (*models.ASFStatsResponse, error)
}

type asfService struct {
	client      *http.Client
	baseUrl     string
	ipcPassword string
}

func NewASFService(serviceConfig config.ServicesConfig) ASFService {
	baseConfig := serviceConfig.ASF
	return &asfService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl:     baseConfig.Url,
		ipcPassword: baseConfig.IPCPassword,
	}
}

func (s *asfService) GetStats() (*models.ASFStatsResponse, error) {
	// Prepare stats request
	statsReq, err := http.NewRequest("GET", fmt.Sprintf("%s/api/bot/asf", s.baseUrl), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare stats request: %w", err)
	}
	statsReq.Header.Add("Authentication", s.ipcPassword)

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
	var stats models.ASFBotInfo
	if err = json.NewDecoder(resp.Body).Decode(&stats); err != nil {
		return nil, fmt.Errorf("failed to parse stats response: %w", err)
	}

	response := models.ASFStatsResponse{
		Total: len(stats.Result),
	}

	for _, botInfo := range stats.Result {
		if botInfo.IsConnectedAndLoggedOn {
			response.Online++
		}

		for _, game := range botInfo.CardsFarmer.GamesToFarm {
			response.CardsRemaining += game.CardsRemaining
		}
	}

	return &response, nil
}
