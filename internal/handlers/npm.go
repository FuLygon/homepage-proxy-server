package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type NPMHandler interface {
	HandleLogin(c *gin.Context)
	HandleStats(c *gin.Context)
}

type npmHandler struct {
	service services.NPMService
}

func NewNPMHandler(service services.NPMService) NPMHandler {
	return &npmHandler{
		service: service,
	}
}

func (h *npmHandler) HandleLogin(c *gin.Context) {
	stats, err := h.service.Login()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *npmHandler) HandleStats(c *gin.Context) {
	// auth token from Homepage
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Missing authorization",
		})
		return
	}

	stats, err := h.service.GetStats(authToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
