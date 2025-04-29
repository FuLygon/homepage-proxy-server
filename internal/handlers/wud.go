package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type WUDHandler interface {
	Handle(c *gin.Context)
}

type wudHandler struct {
	config  *config.Config
	service services.WUDService
}

func NewWUDHandler(config *config.Config, service services.WUDService) WUDHandler {
	return &wudHandler{
		config:  config,
		service: service,
	}
}

func (h *wudHandler) Handle(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.WUD
	stats, err := h.service.GetStats(baseConfig.Url, baseConfig.Username, baseConfig.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, stats)
}
