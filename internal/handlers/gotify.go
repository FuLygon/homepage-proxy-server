package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
)

type GotifyHandler struct {
	config  *config.Config
	service services.GotifyService
}

func NewGotifyHandler(config *config.Config, service services.GotifyService) *GotifyHandler {
	return &GotifyHandler{
		config:  config,
		service: service,
	}
}

func (h *GotifyHandler) Handle(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.Gotify
	stats, err := h.service.GetStats(baseConfig.Url, baseConfig.Key)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}
