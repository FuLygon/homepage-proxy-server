package main

import (
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/cache"
	"homepage-widgets-gateway/internal/handlers"
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
	r := gin.Default()

	// Set trusted proxies
	if len(conf.TrustedProxies) > 0 {
		err = r.SetTrustedProxies(conf.TrustedProxies)
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

	// Setup handlers
	handler := handlers.NewServiceHandler(conf, adguardService, npmService, portainerService, wudService, gotifyService, uptimeKumaService)
	handler.SetupRoutes(r)

	// Start server
	if err = r.Run(":" + conf.Port); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
