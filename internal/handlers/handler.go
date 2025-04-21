package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-proxy-server/config"
)

type ServiceHandler struct {
	config *config.Config
}

func NewServiceHandler(cfg *config.Config) *ServiceHandler {
	return &ServiceHandler{
		config: cfg,
	}
}

// SetupRoutes configures API routes
func (h *ServiceHandler) SetupRoutes(r *gin.Engine) {
	// Service status endpoints
	servicesConfig := h.config.ServicesConfig

	// Register routes based on service configuration
	h.registerServiceRoute(r.GET, "/adguard-home", servicesConfig.AdGuardHome.Enabled, h.getAdGuardHomeStatus)
	h.registerServiceRoute(r.GET, "/nginx-proxy-manager", servicesConfig.NginxProxyManager.Enabled, h.getNginxProxyManagerStatus)
	h.registerServiceRoute(r.GET, "/portainer", servicesConfig.Portainer.Enabled, h.getPortainerStatus)
	h.registerServiceRoute(r.GET, "/wud", servicesConfig.WUD.Enabled, h.getWUDStatus)
	h.registerServiceRoute(r.GET, "/gotify", servicesConfig.Gotify.Enabled, h.getGotifyStatus)
	h.registerServiceRoute(r.GET, "/uptime-kuma", servicesConfig.UptimeKuma.Enabled, h.getUptimeKumaStatus)
}

// registerServiceRoute registers a route if the service is enabled
func (h *ServiceHandler) registerServiceRoute(registerFunc func(string, ...gin.HandlerFunc) gin.IRoutes, path string, enabled bool, handler gin.HandlerFunc) {
	if enabled {
		registerFunc(path, handler)
	}
}

func (h *ServiceHandler) getNginxProxyManagerStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "nginx-proxy-manager",
		"status":  "ok",
	})
}

func (h *ServiceHandler) getAdGuardHomeStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "adguard-home",
		"status":  "ok",
	})
}

func (h *ServiceHandler) getPortainerStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "portainer",
		"status":  "ok",
	})
}

func (h *ServiceHandler) getWUDStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "wud",
		"status":  "ok",
	})
}

func (h *ServiceHandler) getGotifyStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "gotify",
		"status":  "ok",
	})
}

func (h *ServiceHandler) getUptimeKumaStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "uptime-kuma",
		"status":  "ok",
	})
}
