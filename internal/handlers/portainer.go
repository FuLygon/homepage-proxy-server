package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/internal/services"
	"net/http"
	"strconv"
)

type PortainerHandler interface {
	Handle(c *gin.Context)
}

type portainerHandler struct {
	service services.PortainerService
}

func NewPortainerHandler(service services.PortainerService) PortainerHandler {
	return &portainerHandler{
		service: service,
	}
}

func (h *portainerHandler) Handle(c *gin.Context) {
	// extract env param from the Homepage's request
	reqEnvStr := c.Param("env")
	if reqEnvStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing env parameter",
		})
		return
	}

	reqEnv, err := strconv.Atoi(reqEnvStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid env parameter",
		})
		return
	}

	stats, err := h.service.GetStats(reqEnv)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
