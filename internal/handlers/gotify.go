package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type GotifyHandler interface {
	HandleApplication(c *gin.Context)
	HandleClient(c *gin.Context)
	HandleMessage(c *gin.Context)
}

type gotifyHandler struct {
	service services.GotifyService
}

func NewGotifyHandler(service services.GotifyService) GotifyHandler {
	return &gotifyHandler{
		service: service,
	}
}

func (h *gotifyHandler) HandleApplication(c *gin.Context) {
	stats, err := h.service.GetApplications()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *gotifyHandler) HandleClient(c *gin.Context) {
	stats, err := h.service.GetClients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *gotifyHandler) HandleMessage(c *gin.Context) {
	stats, err := h.service.GetMessages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
