package main

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/cache"
	"homepage-widgets-gateway/internal/handlers"
	"homepage-widgets-gateway/internal/routes"
	"homepage-widgets-gateway/internal/services"
	"log"
)

func main() {
	// Load env
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load configuration: ", err)
	}

	// Set Gin mode
	gin.SetMode(conf.GinMode)
	router := gin.Default()

	// Set trusted proxies
	if len(conf.TrustedProxies) > 0 {
		err = router.SetTrustedProxies(conf.TrustedProxies)
		if err != nil {
			log.Printf("error setting trusted proxies: %v", err)
		}
	}

	// Initialize cache
	cacheInstance := cache.NewCache()

	// Setup services
	adguardService := services.NewAdGuardHomeService()
	npmService := services.NewNPMService(cacheInstance)
	portainerService := services.NewPortainerService()
	wudService := services.NewWUDService()
	gotifyService := services.NewGotifyService()
	uptimeKumaService := services.NewUptimeKumaService()
	linkwardenService := services.NewLinkwardenService()

	// Setup handlers
	adguardHandler := handlers.NewAdGuardHandler(conf, adguardService)
	npmHandler := handlers.NewNPMHandler(conf, npmService)
	portainerHandler := handlers.NewPortainerHandler(conf, portainerService)
	wudHandler := handlers.NewWUDHandler(conf, wudService)
	gotifyHandler := handlers.NewGotifyHandler(conf, gotifyService)
	uptimeKumaHandler := handlers.NewUptimeKumaHandler(conf, uptimeKumaService)
	linkwardenHandler := handlers.NewLinkwardenHandler(conf, linkwardenService)

	// Setup routes
	r := routes.NewRoutes(
		router,
		conf,
		adguardHandler,
		npmHandler,
		portainerHandler,
		wudHandler,
		gotifyHandler,
		uptimeKumaHandler,
		linkwardenHandler,
	)

	// Register routes
	r.RegisterRoutes()

	// Start server
	if err = router.Run(":" + conf.Port); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
