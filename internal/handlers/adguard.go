package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type AdGuardHandler interface {
	Handle(c *gin.Context)
}

type adGuardHandler struct {
	config  *config.Config
	service services.AdGuardHomeService
}

func NewAdGuardHandler(config *config.Config, service services.AdGuardHomeService) AdGuardHandler {
	return &adGuardHandler{
		config:  config,
		service: service,
	}
}

func (h *adGuardHandler) Handle(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.AdGuardHome
	stats, err := h.service.GetStats(baseConfig.Url, baseConfig.Username, baseConfig.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
