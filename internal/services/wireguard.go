package services

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"homepage-widgets-gateway/internal/docker"
	"homepage-widgets-gateway/internal/models"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type WireGuardService interface {
	GetLocalStats(interfaceName string, timeout int) (*models.WireGuardStatsResponse, error)
	GetDockerStats(ctx context.Context, interfaceName, container string, timeout int) (*models.WireGuardStatsResponse, error)
	GetExternalStats() (*models.WireGuardStatsResponse, error)
}

type wireGuardService struct {
	docker docker.Docker
}

func NewWireGuardService(docker docker.Docker) WireGuardService {
	return &wireGuardService{docker: docker}
}

const wgExternalClientsDir = "wireguard-clients"

func (w *wireGuardService) GetLocalStats(interfaceName string, timeout int) (*models.WireGuardStatsResponse, error) {
	var (
		stdout bytes.Buffer
		stderr bytes.Buffer
	)

	cmd := exec.Command("wg", "show", interfaceName, "dump")
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		// Check for errors in stderr
		if stderr.Len() > 0 {
			return nil, fmt.Errorf("error executing wg command: %s", stderr.String())
		}
		return nil, fmt.Errorf("error executing wg command: %w", err)
	}

	total, connected, err := w.processOutput(&stdout, time.Duration(timeout))
	if err != nil {
		return nil, fmt.Errorf("error processing WireGuard output: %w", err)
	}

	return &models.WireGuardStatsResponse{
		Total:     total,
		Connected: connected,
	}, nil
}

func (w *wireGuardService) GetDockerStats(ctx context.Context, interfaceName, container string, timeout int) (*models.WireGuardStatsResponse, error) {
	stdout, err := w.docker.Exec(ctx, container, []string{"wg", "show", interfaceName, "dump"})
	if err != nil {
		return nil, fmt.Errorf("error executing wg command: %w", err)
	}

	total, connected, err := w.processOutput(stdout, time.Duration(timeout))
	if err != nil {
		return nil, fmt.Errorf("error processing WireGuard output: %w", err)
	}

	return &models.WireGuardStatsResponse{
		Total:     total,
		Connected: connected,
	}, nil
}

func (w *wireGuardService) GetExternalStats() (*models.WireGuardStatsResponse, error) {
	entries, err := os.ReadDir(wgExternalClientsDir)
	if err != nil {
		return nil, fmt.Errorf("error reading wireguard-clients directory: %w", err)
	}

	var (
		total     int
		connected int
	)

	// Process through each file in the directory
	for _, entry := range entries {
		// Skip subfolders
		if entry.IsDir() {
			continue
		}

		// Read file content
		filePath := filepath.Join(wgExternalClientsDir, entry.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, fmt.Errorf("error reading client file: %w", err)
		}

		// Process content
		if trimmedContent := string(bytes.TrimSpace(content)); trimmedContent == "online" || trimmedContent == "offline" {
			total++
			if trimmedContent == "online" {
				connected++
			}
		}
	}

	return &models.WireGuardStatsResponse{
		Total:     total,
		Connected: connected,
	}, nil
}

func (w *wireGuardService) processOutput(stdout *bytes.Buffer, timeout time.Duration) (total int, connected int, err error) {
	scanner := bufio.NewScanner(stdout)

	// Skip the first line
	if scanner.Scan() {
	}

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "\t")
		total++

		handshakeUnix, err := strconv.ParseInt(fields[4], 10, 64)
		if err != nil {
			return total, connected, fmt.Errorf("failed to parse handshake timestamp: %w", err)
		}

		// If the last handshake is within the timeout, consider it connected
		if handshakeUnix > 0 {
			latestHandshake := time.Unix(handshakeUnix, 0)
			if time.Now().Sub(latestHandshake) <= timeout*time.Minute {
				connected++
			}
		}
	}

	return
}
