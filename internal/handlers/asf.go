package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type ASFHandler interface {
	Handle(c *gin.Context)
}

type asfHandler struct {
	service services.ASFService
}

func NewASFHandler(service services.ASFService) ASFHandler {
	return &asfHandler{
		service: service,
	}
}

func (h *asfHandler) Handle(c *gin.Context) {
	stats, err := h.service.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
