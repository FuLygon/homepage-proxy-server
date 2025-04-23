package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
)

type GotifyHandler interface {
	Handle(c *gin.Context)
}

type gotifyHandler struct {
	config  *config.Config
	service services.GotifyService
}

func NewGotifyHandler(config *config.Config, service services.GotifyService) GotifyHandler {
	return &gotifyHandler{
		config:  config,
		service: service,
	}
}

func (h *gotifyHandler) Handle(c *gin.Context) {
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
