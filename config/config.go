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
		Env     int    `env:"SERVICE_PORTAINER_ENV"`
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
		Slug    string `env:"SERVICE_UPTIME_KUMA_SLUG"`
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
		if config.Url == "" || config.Env <= 0 || config.Key == "" {
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
		if config.Url == "" || config.Slug == "" {
			return fmt.Errorf("missing configuration for Uptime Kuma")
		}
	}

	return nil
}
