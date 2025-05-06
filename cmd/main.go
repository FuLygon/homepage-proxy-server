package main

import (
	"context"
	"github.com/docker/docker/client"
	"github.com/gin-gonic/gin"
	"homepage-widgets-gateway/config"
	"homepage-widgets-gateway/internal/cache"
	"homepage-widgets-gateway/internal/docker"
	"homepage-widgets-gateway/internal/handlers"
	"homepage-widgets-gateway/internal/routes"
	"homepage-widgets-gateway/internal/services"
	"log"
	"time"
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

	// Create Docker Client
	var dockerClient *client.Client
	if conf.WireGuard.Method == "docker" {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		dockerClient, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
		if err != nil {
			log.Fatalf("error creating docker client: %v", err)
		}
		defer dockerClient.Close()

		_, err = dockerClient.Ping(ctx)
		if err != nil {
			log.Fatalf("error pinging docker daemon: %v", err)
		}
	}

	// Initialize cache
	cacheInstance := cache.NewCache()
	// Initialize docker
	dockerInstance := docker.NewDocker(dockerClient)

	// Setup services
	serviceConfig := conf.ServicesConfig
	adguardService := services.NewAdGuardHomeService(serviceConfig)
	npmService := services.NewNPMService(serviceConfig, cacheInstance)
	portainerService := services.NewPortainerService(serviceConfig)
	wudService := services.NewWUDService(serviceConfig)
	gotifyService := services.NewGotifyService(serviceConfig)
	uptimeKumaService := services.NewUptimeKumaService(serviceConfig)
	linkwardenService := services.NewLinkwardenService(serviceConfig)
	yourSpotifyService := services.NewYourSpotifyService(serviceConfig, cacheInstance)
	wireguardService := services.NewWireGuardService(serviceConfig, dockerInstance)
	komodoService := services.NewKomodoService(serviceConfig)

	// Setup handlers
	adguardHandler := handlers.NewAdGuardHandler(adguardService)
	npmHandler := handlers.NewNPMHandler(npmService)
	portainerHandler := handlers.NewPortainerHandler(portainerService)
	wudHandler := handlers.NewWUDHandler(wudService)
	gotifyHandler := handlers.NewGotifyHandler(gotifyService)
	uptimeKumaHandler := handlers.NewUptimeKumaHandler(uptimeKumaService)
	linkwardenHandler := handlers.NewLinkwardenHandler(linkwardenService)
	yourSpotifyHandler := handlers.NewYourSpotifyHandler(yourSpotifyService)
	wireguardHandler := handlers.NewWireGuardHandler(serviceConfig, wireguardService)
	komodoHandler := handlers.NewKomodoHandler(komodoService)

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
		yourSpotifyHandler,
		wireguardHandler,
		komodoHandler,
	)

	// Register routes
	r.RegisterRoutes()

	// Start server
	if err = router.Run(":" + conf.Port); err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
