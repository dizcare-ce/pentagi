// Package config provides configuration management for the pentagi application.
// It loads configuration from environment variables and .env files.
package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds all configuration values for the application.
type Config struct {
	// Server configuration
	ServerHost string
	ServerPort int
	DebugMode  bool

	// Database configuration
	DatabaseURL      string
	DatabaseHost     string
	DatabasePort     int
	DatabaseName     string
	DatabaseUser     string
	DatabasePassword string
	DatabaseSSLMode  string

	// Authentication
	JWTSecret          string
	JWTExpirationHours int

	// AI Provider configuration
	OpenAIAPIKey    string
	OpenAIModel     string
	AnthropicAPIKey string
	AnthropicModel  string

	// Storage configuration
	StoragePath string

	// Logging
	LogLevel  string
	LogFormat string
}

// Load reads configuration from environment variables.
// If a .env file exists in the current directory, it will be loaded first.
func Load() (*Config, error) {
	// Attempt to load .env file, ignore error if it doesn't exist
	_ = godotenv.Load()

	cfg := &Config{
		ServerHost:         getEnv("SERVER_HOST", "0.0.0.0"),
		ServerPort:         getEnvAsInt("SERVER_PORT", 8080),
		DebugMode:          getEnvAsBool("DEBUG_MODE", false),
		DatabaseHost:       getEnv("DB_HOST", "localhost"),
		DatabasePort:       getEnvAsInt("DB_PORT", 5432),
		DatabaseName:       getEnv("DB_NAME", "pentagi"),
		DatabaseUser:       getEnv("DB_USER", "pentagi"),
		DatabasePassword:   getEnv("DB_PASSWORD", ""),
		DatabaseSSLMode:    getEnv("DB_SSLMODE", "disable"),
		JWTSecret:          getEnv("JWT_SECRET", ""),
		JWTExpirationHours: getEnvAsInt("JWT_EXPIRATION_HOURS", 24),
		OpenAIAPIKey:       getEnv("OPENAI_API_KEY", ""),
		OpenAIModel:        getEnv("OPENAI_MODEL", "gpt-4o"),
		AnthropicAPIKey:    getEnv("ANTHROPIC_API_KEY", ""),
		AnthropicModel:     getEnv("ANTHROPIC_MODEL", "claude-3-5-sonnet-20241022"),
		StoragePath:        getEnv("STORAGE_PATH", "./data"),
		LogLevel:           getEnv("LOG_LEVEL", "info"),
		LogFormat:          getEnv("LOG_FORMAT", "json"),
	}

	// Build database URL if not explicitly provided
	cfg.DatabaseURL = getEnv("DATABASE_URL", cfg.buildDatabaseURL())

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return cfg, nil
}

// buildDatabaseURL constructs a PostgreSQL connection string from individual fields.
func (c *Config) buildDatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.DatabaseUser,
		c.DatabasePassword,
		c.DatabaseHost,
		c.DatabasePort,
		c.DatabaseName,
		c.DatabaseSSLMode,
	)
}

// validate checks that required configuration values are present.
func (c *Config) validate() error {
	var missing []string

	if c.JWTSecret == "" {
		missing = append(missing, "JWT_SECRET")
	}

	if len(missing) > 0 {
		return fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	return nil
}

// getEnv returns the value of an environment variable or a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

// getEnvAsInt returns the integer value of an environment variable or a default value.
func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

// getEnvAsBool returns the boolean value of an environment variable or a default value.
func getEnvAsBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		if boolValue, err := strconv.ParseBool(value); err == nil {
			return boolValue
		}
	}
	return defaultValue
}
