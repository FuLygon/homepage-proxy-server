package handlers

import (
	"github.com/gin-gonic/gin"
	"homepage-proxy-server/config"
	"homepage-proxy-server/internal/services"
)

type ServiceHandler interface {
	SetupRoutes(r *gin.Engine)
}

type serviceHandler struct {
	config         *config.Config
	adguardService services.AdGuardHomeService
	npmService     services.NPMService
}

func NewServiceHandler(
	cfg *config.Config,
	adguardService services.AdGuardHomeService,
	npmService services.NPMService,
) ServiceHandler {
	return &serviceHandler{
		config:         cfg,
		adguardService: adguardService,
		npmService:     npmService,
	}
}

// SetupRoutes configures API routes
func (h *serviceHandler) SetupRoutes(r *gin.Engine) {
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
func (h *serviceHandler) registerServiceRoute(registerFunc func(string, ...gin.HandlerFunc) gin.IRoutes, path string, enabled bool, handler gin.HandlerFunc) {
	if enabled {
		registerFunc(path, handler)
	}
}

func (h *serviceHandler) getNginxProxyManagerStatus(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.NginxProxyManager
	stats, err := h.npmService.GetStats(baseConfig.Url, baseConfig.Username, baseConfig.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}

func (h *serviceHandler) getAdGuardHomeStatus(c *gin.Context) {
	baseConfig := h.config.ServicesConfig.AdGuardHome
	stats, err := h.adguardService.GetStats(baseConfig.Url, baseConfig.Username, baseConfig.Password)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, stats)
}

func (h *serviceHandler) getPortainerStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "portainer",
		"status":  "ok",
	})
}

func (h *serviceHandler) getWUDStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "wud",
		"status":  "ok",
	})
}

func (h *serviceHandler) getGotifyStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "gotify",
		"status":  "ok",
	})
}

func (h *serviceHandler) getUptimeKumaStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"service": "uptime-kuma",
		"status":  "ok",
	})
}
