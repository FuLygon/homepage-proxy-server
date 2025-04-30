package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type YourSpotifyHandler interface {
	Handle(c *gin.Context)
}

type yourSpotifyHandler struct {
	config  *config.Config
	service services.YourSpotifyService
}

func NewYourSpotifyHandler(config *config.Config, service services.YourSpotifyService) YourSpotifyHandler {
	return &yourSpotifyHandler{
		config:  config,
		service: service,
	}
}

func (h *yourSpotifyHandler) Handle(c *gin.Context) {
	timeRange := c.Query("time_range")
	if timeRange == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing time_range parameter",
		})
		return
	}

	baseConfig := h.config.ServicesConfig.YourSpotify
	stats, err := h.service.GetStats(baseConfig.Url, baseConfig.Token, timeRange)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
