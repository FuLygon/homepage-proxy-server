package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type AdGuardHandler interface {
	Handle(c *gin.Context)
}

type adGuardHandler struct {
	service services.AdGuardHomeService
}

func NewAdGuardHandler(service services.AdGuardHomeService) AdGuardHandler {
	return &adGuardHandler{
		service: service,
	}
}

func (h *adGuardHandler) Handle(c *gin.Context) {
	stats, err := h.service.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
