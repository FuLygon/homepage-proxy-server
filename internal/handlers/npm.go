package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type NPMHandler interface {
	HandleLogin(c *gin.Context)
	HandleStats(c *gin.Context)
}

type npmHandler struct {
	config  *config.Config
	service services.NPMService
}

func NewNPMHandler(config *config.Config, service services.NPMService) NPMHandler {
	return &npmHandler{
		config:  config,
		service: service,
	}
}

func (h *npmHandler) HandleLogin(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.NginxProxyManager
	stats, err := h.service.Login(baseConfig.Url, baseConfig.Username, baseConfig.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *npmHandler) HandleStats(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.NginxProxyManager
	// auth token from Homepage
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Missing authorization",
		})
		return
	}

	stats, err := h.service.GetStats(baseConfig.Url, authToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
