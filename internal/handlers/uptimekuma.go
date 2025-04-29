package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type UptimeKumaHandler interface {
	HandleStats(c *gin.Context)
	HandleStatsHeartbeat(c *gin.Context)
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

func (h *uptimeKumaHandler) HandleStats(c *gin.Context) {
	// extract slug param from the Homepage's request
	reqSlug := c.Param("slug")
	if reqSlug == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing slug parameter",
		})
		return
	}

	baseConfig := h.config.ServicesConfig.UptimeKuma
	stats, err := h.service.GetStats(baseConfig.Url, reqSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *uptimeKumaHandler) HandleStatsHeartbeat(c *gin.Context) {
	// extract slug param from the Homepage's request
	reqSlug := c.Param("slug")
	if reqSlug == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing slug parameter",
		})
		return
	}

	baseConfig := h.config.ServicesConfig.UptimeKuma
	stats, err := h.service.GetStatsHeartbeat(baseConfig.Url, reqSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
