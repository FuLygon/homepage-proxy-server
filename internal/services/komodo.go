package services

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/sync/errgroup"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/models"
	"net/http"
	"time"
)

type KomodoService interface {
	GetStats(ctx context.Context) (*models.KomodoStatsResponse, error)
}

type summaryModel interface {
	SummaryRequest() map[string]interface{}
}

type komodoService struct {
	client            *http.Client
	baseUrl           string
	key               string
	secret            string
	extraStack        bool
	extraBuild        bool
	extraRepo         bool
	extraAction       bool
	extraBuilder      bool
	extraDeployment   bool
	extraProcedure    bool
	extraResourceSync bool
}

func NewKomodoService(serviceConfig config.ServicesConfig) KomodoService {
	baseConfig := serviceConfig.Komodo
	extraStatsBool := make([]bool, 8)

	func() {
		for _, stat := range baseConfig.ExtraStats {
			switch stat {
			case "stack":
				extraStatsBool[0] = true
			case "build":
				extraStatsBool[1] = true
			case "repo":
				extraStatsBool[2] = true
			case "action":
				extraStatsBool[3] = true
			case "builder":
				extraStatsBool[4] = true
			case "deployment":
				extraStatsBool[5] = true
			case "procedure":
				extraStatsBool[6] = true
			case "resource-sync":
				extraStatsBool[7] = true
			}
		}
	}()

	return &komodoService{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		baseUrl:           baseConfig.Url,
		key:               baseConfig.Key,
		secret:            baseConfig.Secret,
		extraStack:        extraStatsBool[0],
		extraBuild:        extraStatsBool[1],
		extraRepo:         extraStatsBool[2],
		extraAction:       extraStatsBool[3],
		extraBuilder:      extraStatsBool[4],
		extraDeployment:   extraStatsBool[5],
		extraProcedure:    extraStatsBool[6],
		extraResourceSync: extraStatsBool[7],
	}
}

func (s *komodoService) GetStats(ctx context.Context) (*models.KomodoStatsResponse, error) {
	var response models.KomodoStatsResponse
	g, ctx := errgroup.WithContext(ctx)

	// Fetch container summary stats
	g.Go(func() error {
		summary, err := s.getSummaryStats(ctx, &models.KomodoContainerStats{})
		if err != nil {
			return fmt.Errorf("failed to fetch container summary stats: %w", err)
		}
		response.Container = summary.(*models.KomodoContainerStats)
		return nil
	})

	// Fetch stack summary stats
	if s.extraStack {
		g.Go(func() error {
			summary, err := s.getSummaryStats(ctx, &models.KomodoStackStats{})
			if err != nil {
				return fmt.Errorf("failed to fetch stack summary stats: %w", err)
			}
			response.Stack = summary.(*models.KomodoStackStats)
			return nil
		})
	}

	// Fetch build summary stats
	if s.extraBuild {
		g.Go(func() error {
			summary, err := s.getSummaryStats(ctx, &models.KomodoBuildStats{})
			if err != nil {
				return fmt.Errorf("failed to fetch build summary stats: %w", err)
			}
			response.Build = summary.(*models.KomodoBuildStats)
			return nil
		})
	}

	// Fetch repo summary stats
	if s.extraRepo {
		g.Go(func() error {
			summary, err := s.getSummaryStats(ctx, &models.KomodoRepoStats{})
			if err != nil {
				return fmt.Errorf("failed to fetch repo summary stats: %w", err)
			}
			response.Repo = summary.(*models.KomodoRepoStats)
			return nil
		})
	}

	// Fetch action summary stats
	if s.extraAction {
		g.Go(func() error {
			summary, err := s.getSummaryStats(ctx, &models.KomodoActionStats{})
			if err != nil {
				return fmt.Errorf("failed to fetch action summary stats: %w", err)
			}
			response.Action = summary.(*models.KomodoActionStats)
			return nil
		})
	}

	// Fetch builder summary stats
	if s.extraBuilder {
		g.Go(func() error {
			summary, err := s.getSummaryStats(ctx, &models.KomodoBuilderStats{})
			if err != nil {
				return fmt.Errorf("failed to fetch builder summary stats: %w", err)
			}
			response.Builder = summary.(*models.KomodoBuilderStats)
			return nil
		})
	}

	// Fetch deployment summary stats
	if s.extraDeployment {
		g.Go(func() error {
			summary, err := s.getSummaryStats(ctx, &models.KomodoDeploymentStats{})
			if err != nil {
				return fmt.Errorf("failed to fetch deployment summary stats: %w", err)
			}
			response.Deployment = summary.(*models.KomodoDeploymentStats)
			return nil
		})
	}

	// Fetch procedure summary stats
	if s.extraProcedure {
		g.Go(func() error {
			summary, err := s.getSummaryStats(ctx, &models.KomodoProcedureStats{})
			if err != nil {
				return fmt.Errorf("failed to fetch procedure summary stats: %w", err)
			}
			response.Procedure = summary.(*models.KomodoProcedureStats)
			return nil
		})
	}

	// Fetch resource sync summary stats
	if s.extraResourceSync {
		g.Go(func() error {
			summary, err := s.getSummaryStats(ctx, &models.KomodoResourceSyncStats{})
			if err != nil {
				return fmt.Errorf("failed to fetch resource sync summary stats: %w", err)
			}
			response.ResourceSync = summary.(*models.KomodoResourceSyncStats)
			return nil
		})
	}

	// Wait for goroutines
	if err := g.Wait(); err != nil {
		return nil, err
	}

	return &response, nil
}

func (s *komodoService) getSummaryStats(ctx context.Context, summary summaryModel) (interface{}, error) {
	payloadJson, err := json.Marshal(summary.SummaryRequest())
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	// Prepare stats request
	statsReq, err := http.NewRequestWithContext(ctx, "POST", fmt.Sprintf("%s/read", s.baseUrl), bytes.NewBuffer(payloadJson))
	if err != nil {
		return nil, fmt.Errorf("failed to prepare stats request: %w", err)
	}
	statsReq.Header.Add("X-Api-Key", s.key)
	statsReq.Header.Add("X-Api-Secret", s.secret)
	statsReq.Header.Set("Content-Type", "application/json")

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
	if err = json.NewDecoder(resp.Body).Decode(summary); err != nil {
		return nil, fmt.Errorf("failed to parse stats response: %w", err)
	}

	return summary, nil
}
