package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type UptimeKumaHandler interface {
	HandleStats(c *gin.Context)
	HandleStatsHeartbeat(c *gin.Context)
}

type uptimeKumaHandler struct {
	service services.UptimeKumaService
}

func NewUptimeKumaHandler(service services.UptimeKumaService) UptimeKumaHandler {
	return &uptimeKumaHandler{
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

	stats, err := h.service.GetStats(reqSlug)
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

	stats, err := h.service.GetStatsHeartbeat(reqSlug)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
