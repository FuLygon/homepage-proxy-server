package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type LinkwardenHandler interface {
	HandleCollections(c *gin.Context)
	HandleTags(c *gin.Context)
}

type linkwardenHandler struct {
	config  *config.Config
	service services.LinkwardenService
}

func NewLinkwardenHandler(config *config.Config, service services.LinkwardenService) LinkwardenHandler {
	return &linkwardenHandler{
		config:  config,
		service: service,
	}
}

func (h *linkwardenHandler) HandleCollections(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.Linkwarden
	collections, err := h.service.GetCollections(baseConfig.Url, baseConfig.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, collections)
}

func (h *linkwardenHandler) HandleTags(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.Linkwarden
	tags, err := h.service.GetTags(baseConfig.Url, baseConfig.Key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tags)
}
