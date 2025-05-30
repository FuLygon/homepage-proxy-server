package config

import (
	"fmt"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	Port           string   `env:"PORT" envDefault:"8080"`
	GinMode        string   `env:"GIN_MODE" envDefault:"debug"`
	LogLevel       string   `env:"LOG_LEVEL" envDefault:"info"`
	TrustedProxies []string `env:"TRUSTED_PROXIES" envDefault:"10.0.0.0/8,172.16.0.0/12,192.168.0.0/16" envSeparator:","`
	ServicesConfig
}

type ServicesConfig struct {
	AdGuardHome struct {
		Enabled  bool   `env:"SERVICE_AGH_ENABLED" envDefault:"false"`
		Url      string `env:"SERVICE_AGH_URL"`
		Username string `env:"SERVICE_AGH_USERNAME"`
		Password string `env:"SERVICE_AGH_PASSWORD"`
	}
	NginxProxyManager struct {
		Enabled  bool   `env:"SERVICE_NPM_ENABLED" envDefault:"false"`
		Url      string `env:"SERVICE_NPM_URL"`
		Username string `env:"SERVICE_NPM_USERNAME"`
		Password string `env:"SERVICE_NPM_PASSWORD"`
	}
	Portainer struct {
		Enabled bool   `env:"SERVICE_PORTAINER_ENABLED" envDefault:"false"`
		Url     string `env:"SERVICE_PORTAINER_URL"`
		Key     string `env:"SERVICE_PORTAINER_KEY"`
	}
	WUD struct {
		Enabled  bool   `env:"SERVICE_WUD_ENABLED" envDefault:"false"`
		Url      string `env:"SERVICE_WUD_URL"`
		Username string `env:"SERVICE_WUD_USERNAME"`
		Password string `env:"SERVICE_WUD_PASSWORD"`
	}
	Gotify struct {
		Enabled bool   `env:"SERVICE_GOTIFY_ENABLED" envDefault:"false"`
		Url     string `env:"SERVICE_GOTIFY_URL"`
		Key     string `env:"SERVICE_GOTIFY_KEY"`
	}
	UptimeKuma struct {
		Enabled bool   `env:"SERVICE_UPTIME_KUMA_ENABLED" envDefault:"false"`
		Url     string `env:"SERVICE_UPTIME_KUMA_URL"`
	}
	Linkwarden struct {
		Enabled bool   `env:"SERVICE_LINKWARDEN_ENABLED" envDefault:"false"`
		Url     string `env:"SERVICE_LINKWARDEN_URL"`
		Key     string `env:"SERVICE_LINKWARDEN_KEY"`
	}
	YourSpotify struct {
		Enabled bool   `env:"SERVICE_YOUR_SPOTIFY_ENABLED" envDefault:"false"`
		Url     string `env:"SERVICE_YOUR_SPOTIFY_URL"`
		Token   string `env:"SERVICE_YOUR_SPOTIFY_TOKEN"`
	}
	WireGuard struct {
		Enabled         bool   `env:"SERVICE_WIREGUARD_ENABLED" envDefault:"false"`
		Method          string `env:"SERVICE_WIREGUARD_METHOD"`
		Interface       string `env:"SERVICE_WIREGUARD_INTERFACE"`
		Timeout         int    `env:"SERVICE_WIREGUARD_TIMEOUT" envDefault:"5"`
		DockerContainer string `env:"SERVICE_WIREGUARD_DOCKER_CONTAINER"`
	}
	Komodo struct {
		Enabled    bool     `env:"SERVICE_KOMODO_ENABLED" envDefault:"false"`
		Url        string   `env:"SERVICE_KOMODO_URL"`
		Key        string   `env:"SERVICE_KOMODO_KEY"`
		Secret     string   `env:"SERVICE_KOMODO_SECRET"`
		ExtraStats []string `env:"SERVICE_KOMODO_EXTRA_STATS" envSeparator:","`
	}
	ASF struct {
		Enabled     bool   `env:"SERVICE_ASF_ENABLED" envDefault:"false"`
		Url         string `env:"SERVICE_ASF_URL"`
		IPCPassword string `env:"SERVICE_ASF_IPC_PASSWORD"`
	}
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	// Validate services configuration
	if err := validateServicesConfig(cfg); err != nil {
		return nil, fmt.Errorf("services configuration validation failed: %w", err)
	}

	return cfg, nil
}

// validateServicesConfig validates the service configuration
func validateServicesConfig(cfg *Config) error {
	// Validate AdGuard Home
	if config := cfg.AdGuardHome; config.Enabled {
		if config.Url == "" || config.Username == "" || config.Password == "" {
			return fmt.Errorf("missing configuration for AdGuard Home")
		}
	}

	// Validate Nginx Proxy Manager
	if config := cfg.NginxProxyManager; config.Enabled {
		if config.Url == "" || config.Username == "" || config.Password == "" {
			return fmt.Errorf("missing configuration for Nginx Proxy Manager")
		}
	}

	// Validate Portainer
	if config := cfg.Portainer; config.Enabled {
		if config.Url == "" || config.Key == "" {
			return fmt.Errorf("missing configuration for Portainer")
		}
	}

	// Validate WUD
	if config := cfg.WUD; config.Enabled {
		if config.Url == "" || config.Username == "" || config.Password == "" {
			return fmt.Errorf("missing configuration for WUD")
		}
	}

	// Validate Gotify
	if config := cfg.Gotify; config.Enabled {
		if config.Url == "" || config.Key == "" {
			return fmt.Errorf("missing configuration for Gotify")
		}
	}

	// Validate Uptime Kuma
	if config := cfg.UptimeKuma; config.Enabled {
		if config.Url == "" {
			return fmt.Errorf("missing configuration for Uptime Kuma")
		}
	}

	// Validate Linkwarden
	if config := cfg.Linkwarden; config.Enabled {
		if config.Url == "" || config.Key == "" {
			return fmt.Errorf("missing configuration for Linkwarden")
		}
	}

	// Validate Your Spotify
	if config := cfg.YourSpotify; config.Enabled {
		if config.Url == "" || config.Token == "" {
			return fmt.Errorf("missing configuration for Your Spotify")
		}
	}

	// Validate WireGuard
	if config := cfg.WireGuard; config.Enabled {
		if config.Method == "" {
			return fmt.Errorf("missing configuration for WireGuard")
		}

		if config.Method != "docker" && config.Method != "local" && config.Method != "external" {
			return fmt.Errorf("misconfigured WireGuard method")
		}

		if (config.Method == "docker" || config.Method == "local") && config.Interface == "" {
			return fmt.Errorf("missing WireGuard interface name")
		}

		if config.Method == "docker" && config.DockerContainer == "" {
			return fmt.Errorf("missing Docker container name for WireGuard")
		}
	}

	// Validate Komodo
	if config := cfg.Komodo; config.Enabled {
		if config.Url == "" || config.Key == "" || config.Secret == "" {
			return fmt.Errorf("missing configuration for Komodo")
		}

		validExtraStats := map[string]struct{}{
			"stack":         {},
			"build":         {},
			"repo":          {},
			"action":        {},
			"builder":       {},
			"deployment":    {},
			"procedure":     {},
			"resource-sync": {},
		}

		if len(config.ExtraStats) > 0 {
			for _, stat := range config.ExtraStats {
				if _, ok := validExtraStats[stat]; !ok {
					return fmt.Errorf("invalid komodo extra stats config: %s", stat)
				}
			}
		}
	}

	// Validate ASF
	if config := cfg.ASF; config.Enabled {
		if config.Url == "" || config.IPCPassword == "" {
			return fmt.Errorf("missing configuration for ASF")
		}
	}

	return nil
}
