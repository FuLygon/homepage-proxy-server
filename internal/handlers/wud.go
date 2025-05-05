package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type WUDHandler interface {
	Handle(c *gin.Context)
}

type wudHandler struct {
	service services.WUDService
}

func NewWUDHandler(service services.WUDService) WUDHandler {
	return &wudHandler{
		service: service,
	}
}

func (h *wudHandler) Handle(c *gin.Context) {
	stats, err := h.service.GetStats()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
