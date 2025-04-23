package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
)

type WUDHandler struct {
	config  *config.Config
	service services.WUDService
}

func NewWUDHandler(config *config.Config, service services.WUDService) *WUDHandler {
	return &WUDHandler{
		config:  config,
		service: service,
	}
}

func (h *WUDHandler) Handle(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.WUD
	stats, err := h.service.GetStats(baseConfig.Url, baseConfig.Username, baseConfig.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}
