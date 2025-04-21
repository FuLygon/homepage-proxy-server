package main

import (
	"github.com/gin-gonic/gin"
	"homepage-proxy-server/config"
	"homepage-proxy-server/internal/cache"
	"homepage-proxy-server/internal/handlers"
	"homepage-proxy-server/internal/services"
	"log"
)

func main() {
	// Load env
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal("failed to load configuration: ", err)
	}

	// Set Gin mode
	if conf.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
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
	adguardService := services.NewAdGuardHomeService(cacheInstance)

	// Setup handlers
	handler := handlers.NewServiceHandler(conf, adguardService)
	handler.SetupRoutes(r)

	// Start server
	if err = r.Run(":" + conf.Port); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
