package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
)

type PortainerHandler struct {
	config  *config.Config
	service services.PortainerService
}

func NewPortainerHandler(config *config.Config, service services.PortainerService) *PortainerHandler {
	return &PortainerHandler{
		config:  config,
		service: service,
	}
}

func (h *PortainerHandler) Handle(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.Portainer
	stats, err := h.service.GetStats(baseConfig.Url, baseConfig.Key, baseConfig.Env)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}
