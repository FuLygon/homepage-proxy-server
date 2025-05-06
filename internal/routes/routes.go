package routes

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/handlers"
	"net/http"
)

type Routes interface {
	RegisterRoutes()
}

type routes struct {
	router             *gin.Engine
	config             *config.Config
	adguardHandler     handlers.AdGuardHandler
	npmHandler         handlers.NPMHandler
	portainerHandler   handlers.PortainerHandler
	wudHandler         handlers.WUDHandler
	gotifyHandler      handlers.GotifyHandler
	uptimeKumaHandler  handlers.UptimeKumaHandler
	linkwardenHandler  handlers.LinkwardenHandler
	yourSpotifyHandler handlers.YourSpotifyHandler
	wireGuardHandler   handlers.WireGuardHandler
	komodoHandler      handlers.KomodoHandler
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
	linkwardenHandler handlers.LinkwardenHandler,
	yourSpotifyHandler handlers.YourSpotifyHandler,
	wireGuardHandler handlers.WireGuardHandler,
	komodoHandler handlers.KomodoHandler,
) Routes {
	return &routes{
		router:             router,
		config:             config,
		adguardHandler:     adguardHandler,
		npmHandler:         npmHandler,
		portainerHandler:   portainerHandler,
		wudHandler:         wudHandler,
		gotifyHandler:      gotifyHandler,
		uptimeKumaHandler:  uptimeKumaHandler,
		linkwardenHandler:  linkwardenHandler,
		yourSpotifyHandler: yourSpotifyHandler,
		wireGuardHandler:   wireGuardHandler,
		komodoHandler:      komodoHandler,
	}
}

// RegisterRoutes configures API routes
func (r *routes) RegisterRoutes() {
	router := r.router
	// Base services config
	servicesConfig := r.config.ServicesConfig

	// Health check endpoint
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	// Adguard Home - endpoint references from: https://github.com/gethomepage/homepage/blob/main/src/widgets/adguard/widget.js
	r.registerServiceRoute(router.GET, "/adguard-home/control/stats", servicesConfig.AdGuardHome.Enabled, r.adguardHandler.Handle)

	// Nginx Proxy Manager - endpoints references from: https://github.com/gethomepage/homepage/blob/main/src/widgets/npm/widget.js
	r.registerServiceRoute(router.POST, "/nginx-proxy-manager/api/tokens", servicesConfig.NginxProxyManager.Enabled, r.npmHandler.HandleLogin)
	r.registerServiceRoute(router.GET, "/nginx-proxy-manager/api/nginx/proxy-hosts", servicesConfig.NginxProxyManager.Enabled, r.npmHandler.HandleStats)

	// Portainer - endpoint references from: https://github.com/gethomepage/homepage/blob/main/src/widgets/portainer/widget.js
	r.registerServiceRoute(router.GET, "/portainer/api/endpoints/:env/docker/containers/json", servicesConfig.Portainer.Enabled, r.portainerHandler.Handle)

	// WUD (What's Up Docker) - endpoint references from: https://github.com/gethomepage/homepage/blob/main/src/widgets/whatsupdocker/widget.js
	r.registerServiceRoute(router.GET, "/wud/api/containers", servicesConfig.WUD.Enabled, r.wudHandler.Handle)

	// Gotify - endpoints references from: https://github.com/gethomepage/homepage/blob/main/src/widgets/gotify/widget.js
	r.registerServiceRoute(router.GET, "/gotify/application", servicesConfig.Gotify.Enabled, r.gotifyHandler.HandleApplication)
	r.registerServiceRoute(router.GET, "/gotify/client", servicesConfig.Gotify.Enabled, r.gotifyHandler.HandleClient)
	r.registerServiceRoute(router.GET, "/gotify/message", servicesConfig.Gotify.Enabled, r.gotifyHandler.HandleMessage)

	// Uptime Kuma - endpoints references from: https://github.com/gethomepage/homepage/blob/main/src/widgets/uptimekuma/widget.js
	r.registerServiceRoute(router.GET, "/uptime-kuma/api/status-page/:slug", servicesConfig.UptimeKuma.Enabled, r.uptimeKumaHandler.HandleStats)
	r.registerServiceRoute(router.GET, "/uptime-kuma/api/status-page/heartbeat/:slug", servicesConfig.UptimeKuma.Enabled, r.uptimeKumaHandler.HandleStatsHeartbeat)

	// Linkwarden - endpoints references from: https://github.com/gethomepage/homepage/blob/main/src/widgets/linkwarden/widget.js
	r.registerServiceRoute(router.GET, "/linkwarden/api/v1/collections", servicesConfig.Linkwarden.Enabled, r.linkwardenHandler.HandleCollections)
	r.registerServiceRoute(router.GET, "/linkwarden/api/v1/tags", servicesConfig.Linkwarden.Enabled, r.linkwardenHandler.HandleTags)

	// Your Spotify
	r.registerServiceRoute(router.GET, "/your-spotify", servicesConfig.YourSpotify.Enabled, r.yourSpotifyHandler.Handle)

	// WireGuard
	r.registerServiceRoute(router.GET, "/wireguard", servicesConfig.WireGuard.Enabled, r.wireGuardHandler.Handle)

	// Komodo
	r.registerServiceRoute(router.GET, "/komodo", servicesConfig.Komodo.Enabled, r.komodoHandler.Handle)
}

// registerServiceRoute registers a route if the service is enabled
func (r *routes) registerServiceRoute(registerFunc func(string, ...gin.HandlerFunc) gin.IRoutes, path string, enabled bool, handler gin.HandlerFunc) {
	if enabled {
		registerFunc(path, handler)
	}
}
