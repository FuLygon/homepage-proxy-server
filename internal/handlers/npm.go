package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
)

type NPMHandler struct {
	config  *config.Config
	service services.NPMService
}

func NewNPMHandler(config *config.Config, service services.NPMService) *NPMHandler {
	return &NPMHandler{
		config:  config,
		service: service,
	}
}

func (h *NPMHandler) HandleLogin(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.NginxProxyManager
	stats, err := h.service.Login(baseConfig.Url, baseConfig.Username, baseConfig.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}

func (h *NPMHandler) HandleStats(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.NginxProxyManager
	// auth token from Homepage
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		c.JSON(403, gin.H{
			"error": "Missing authorization",
		})
		return
	}

	stats, err := h.service.GetStats(baseConfig.Url, authToken)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}
