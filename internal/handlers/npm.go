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

func (h *NPMHandler) Handle(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.NginxProxyManager
	stats, err := h.service.GetStats(baseConfig.Url, baseConfig.Username, baseConfig.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}
