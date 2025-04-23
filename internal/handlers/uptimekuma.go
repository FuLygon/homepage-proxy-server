package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
)

type UptimeKumaHandler interface {
	Handle(c *gin.Context)
}

type uptimeKumaHandler struct {
	config  *config.Config
	service services.UptimeKumaService
}

func NewUptimeKumaHandler(config *config.Config, service services.UptimeKumaService) UptimeKumaHandler {
	return &uptimeKumaHandler{
		config:  config,
		service: service,
	}
}

func (h *uptimeKumaHandler) Handle(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.UptimeKuma
	stats, err := h.service.GetStats(baseConfig.Url, baseConfig.Slug)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}
