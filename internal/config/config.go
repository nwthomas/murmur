package config

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all configuration for the application
type Config struct {
	// API Configuration
	OpenAIAPIKey string
	Model        string
	
	// Application Configuration
	Debug        bool
	LogLevel     string
	
	// UI Configuration
	Theme        string
	MaxRetries   int
	Timeout      int
}

// Load loads configuration from environment variables
func Load() (*Config, error) {
	config := &Config{
		OpenAIAPIKey: getEnv("OPENAI_API_KEY", ""),
		Model:        getEnv("OPENAI_MODEL", "gpt-3.5-turbo"),
		Debug:        getEnvBool("DEBUG", false),
		LogLevel:     getEnv("LOG_LEVEL", "info"),
		Theme:        getEnv("THEME", "default"),
		MaxRetries:   getEnvInt("MAX_RETRIES", 3),
		Timeout:      getEnvInt("TIMEOUT", 30),
	}

	// Validate required configuration
	if config.OpenAIAPIKey == "" {
		return nil, fmt.Errorf("OPENAI_API_KEY environment variable is required")
	}

	return config, nil
}

// getEnv gets an environment variable with a fallback value
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// getEnvBool gets a boolean environment variable with a fallback value
func getEnvBool(key string, fallback bool) bool {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
	}
	return fallback
}

// getEnvInt gets an integer environment variable with a fallback value
func getEnvInt(key string, fallback int) int {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.Atoi(value); err == nil {
			return parsed
		}
	}
	return fallback
}
