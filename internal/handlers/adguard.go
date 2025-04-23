package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
)

type AdGuardHandler struct {
	config  *config.Config
	service services.AdGuardHomeService
}

func NewAdGuardHandler(config *config.Config, service services.AdGuardHomeService) *AdGuardHandler {
	return &AdGuardHandler{
		config:  config,
		service: service,
	}
}

func (h *AdGuardHandler) Handle(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.AdGuardHome
	stats, err := h.service.GetStats(baseConfig.Url, baseConfig.Username, baseConfig.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}
