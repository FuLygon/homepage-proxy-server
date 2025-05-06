package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type KomodoHandler interface {
	Handle(c *gin.Context)
}

type komodoHandler struct {
	service services.KomodoService
}

func NewKomodoHandler(service services.KomodoService) KomodoHandler {
	return &komodoHandler{
		service: service,
	}
}

func (h *komodoHandler) Handle(c *gin.Context) {
	stats, err := h.service.GetStats(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
