package config

import (
	"os"
)

func Load() *Config {
	return &Config{
		Service: Service{
			BaseURL: getEnv("BASE_URL", "https://api.mercadobitcoin.net/api/v4/"),
			Port:    getEnv("PORT", "8080"),
		},
		Database: Database{
			Host:     getEnv("DB_HOST", "db"),
			Port:     getEnv("DB_PORT", "5432"),
			Name:     getEnv("DB_NAME", ""),
			Password: getEnv("DB_PASSWORD", ""),
			User:     getEnv("DB_USER", ""),
		},
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
