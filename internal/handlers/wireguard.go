package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type WireGuardHandler interface {
	Handle(c *gin.Context)
}

type wireGuardHandler struct {
	config  config.ServicesConfig
	service services.WireGuardService
}

func NewWireGuardHandler(config config.ServicesConfig, service services.WireGuardService) WUDHandler {
	return &wireGuardHandler{
		config:  config,
		service: service,
	}
}

func (h *wireGuardHandler) Handle(c *gin.Context) {
	switch h.config.WireGuard.Method {
	case "local":
		response, err := h.service.GetLocalStats()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, response)
	case "docker":
		response, err := h.service.GetDockerStats(c.Request.Context())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, response)
	case "external":
		response, err := h.service.GetExternalStats()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, response)
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid data fetch method",
		})
		return
	}
}
