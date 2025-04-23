package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
)

type PortainerHandler interface {
	Handle(c *gin.Context)
}

type portainerHandler struct {
	config  *config.Config
	service services.PortainerService
}

func NewPortainerHandler(config *config.Config, service services.PortainerService) PortainerHandler {
	return &portainerHandler{
		config:  config,
		service: service,
	}
}

func (h *portainerHandler) Handle(c *gin.Context) {
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
