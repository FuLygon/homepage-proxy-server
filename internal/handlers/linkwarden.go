package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/internal/services"
	"net/http"
)

type LinkwardenHandler interface {
	HandleCollections(c *gin.Context)
	HandleTags(c *gin.Context)
}

type linkwardenHandler struct {
	service services.LinkwardenService
}

func NewLinkwardenHandler(service services.LinkwardenService) LinkwardenHandler {
	return &linkwardenHandler{
		service: service,
	}
}

func (h *linkwardenHandler) HandleCollections(c *gin.Context) {
	collections, err := h.service.GetCollections()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, collections)
}

func (h *linkwardenHandler) HandleTags(c *gin.Context) {
	tags, err := h.service.GetTags()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, tags)
}
