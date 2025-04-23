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
	adguardHandler    *handlers.AdGuardHandler
	npmHandler        *handlers.NPMHandler
	portainerHandler  *handlers.PortainerHandler
	wudHandler        *handlers.WUDHandler
	gotifyHandler     *handlers.GotifyHandler
	uptimeKumaHandler *handlers.UptimeKumaHandler
}

func NewRoutes(
	router *gin.Engine,
	config *config.Config,
	adguardHandler *handlers.AdGuardHandler,
	npmHandler *handlers.NPMHandler,
	portainerHandler *handlers.PortainerHandler,
	wudHandler *handlers.WUDHandler,
	gotifyHandler *handlers.GotifyHandler,
	uptimeKumaHandler *handlers.UptimeKumaHandler,
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
	ginRouter := r.router
	// Base services config
	servicesConfig := r.config.ServicesConfig

	// Health check endpoint
	ginRouter.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	// Adguard Home
	r.registerServiceRoute(ginRouter.GET, "/adguard-home/control/stats", servicesConfig.AdGuardHome.Enabled, r.adguardHandler.Handle)
	r.registerServiceRoute(ginRouter.GET, "/nginx-proxy-manager", servicesConfig.NginxProxyManager.Enabled, r.npmHandler.Handle)
	r.registerServiceRoute(ginRouter.GET, "/portainer", servicesConfig.Portainer.Enabled, r.portainerHandler.Handle)
	r.registerServiceRoute(ginRouter.GET, "/wud", servicesConfig.WUD.Enabled, r.wudHandler.Handle)
	r.registerServiceRoute(ginRouter.GET, "/gotify", servicesConfig.Gotify.Enabled, r.gotifyHandler.Handle)
	r.registerServiceRoute(ginRouter.GET, "/uptime-kuma", servicesConfig.UptimeKuma.Enabled, r.uptimeKumaHandler.Handle)
}

// registerServiceRoute registers a route if the service is enabled
func (r *routes) registerServiceRoute(registerFunc func(string, ...gin.HandlerFunc) gin.IRoutes, path string, enabled bool, handler gin.HandlerFunc) {
	if enabled {
		registerFunc(path, handler)
	}
}
