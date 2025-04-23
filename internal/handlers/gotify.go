package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
)

type GotifyHandler interface {
	HandleApplication(c *gin.Context)
	HandleClient(c *gin.Context)
	HandleMessage(c *gin.Context)
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

func (h *gotifyHandler) HandleApplication(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.Gotify
	stats, err := h.service.GetApplications(baseConfig.Url, baseConfig.Key)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}

func (h *gotifyHandler) HandleClient(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.Gotify
	stats, err := h.service.GetClients(baseConfig.Url, baseConfig.Key)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}

func (h *gotifyHandler) HandleMessage(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.Gotify
	stats, err := h.service.GetMessages(baseConfig.Url, baseConfig.Key)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}
