package routes

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/handlers"
)

type Routes interface {
	RegisterRoutes()
}

type routes struct {
	router            *gin.Engine
	config            *config.Config
	adguardHandler    handlers.AdGuardHandler
	npmHandler        handlers.NPMHandler
	portainerHandler  handlers.PortainerHandler
	wudHandler        handlers.WUDHandler
	gotifyHandler     handlers.GotifyHandler
	uptimeKumaHandler handlers.UptimeKumaHandler
}

func NewRoutes(
	router *gin.Engine,
	config *config.Config,
	adguardHandler handlers.AdGuardHandler,
	npmHandler handlers.NPMHandler,
	portainerHandler handlers.PortainerHandler,
	wudHandler handlers.WUDHandler,
	gotifyHandler handlers.GotifyHandler,
	uptimeKumaHandler handlers.UptimeKumaHandler,
) Routes {
	return &routes{
		router:            router,
		config:            config,
		adguardHandler:    adguardHandler,
		npmHandler:        npmHandler,
		portainerHandler:  portainerHandler,
		wudHandler:        wudHandler,
		gotifyHandler:     gotifyHandler,
		uptimeKumaHandler: uptimeKumaHandler,
	}
}

// RegisterRoutes configures API routes
func (r *routes) RegisterRoutes() {
	router := r.router
	// Base services config
	servicesConfig := r.config.ServicesConfig

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// Adguard Home
	r.registerServiceRoute(router.GET, "/adguard-home/control/stats", servicesConfig.AdGuardHome.Enabled, r.adguardHandler.Handle)

	// Nginx Proxy Manager
	r.registerServiceRoute(router.POST, "/nginx-proxy-manager/api/tokens", servicesConfig.NginxProxyManager.Enabled, r.npmHandler.HandleLogin)
	r.registerServiceRoute(router.GET, "/nginx-proxy-manager/api/nginx/proxy-hosts", servicesConfig.NginxProxyManager.Enabled, r.npmHandler.HandleStats)

	// Portainer
	r.registerServiceRoute(router.GET, "/portainer/api/endpoints/:env/docker/containers/json", servicesConfig.Portainer.Enabled, r.portainerHandler.Handle)

	// WUD (What's Up Docker)
	r.registerServiceRoute(router.GET, "/wud/api/containers", servicesConfig.WUD.Enabled, r.wudHandler.Handle)

	// Gotify
	r.registerServiceRoute(router.GET, "/gotify", servicesConfig.Gotify.Enabled, r.gotifyHandler.Handle)

	// Uptime Kuma
	r.registerServiceRoute(router.GET, "/uptime-kuma/api/status-page/:slug", servicesConfig.UptimeKuma.Enabled, r.uptimeKumaHandler.HandleStats)
	r.registerServiceRoute(router.GET, "/uptime-kuma/api/status-page/heartbeat/:slug", servicesConfig.UptimeKuma.Enabled, r.uptimeKumaHandler.HandleStatsHeartbeat)
}

// registerServiceRoute registers a route if the service is enabled
func (r *routes) registerServiceRoute(registerFunc func(string, ...gin.HandlerFunc) gin.IRoutes, path string, enabled bool, handler gin.HandlerFunc) {
	if enabled {
		registerFunc(path, handler)
	}
}
