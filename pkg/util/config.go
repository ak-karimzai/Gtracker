// Package util provides utility functions for configuration handling.
package util

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

// ENV_DEVELOPMENT represents the development environment.
const ENV_DEVELOPMENT = "DEV"

// ENV_PRODUCTION represents the production environment.
const ENV_PRODUCTION = "PROD"

// Config represents the application configuration.
type Config struct {
	ServerPort          string        `mapstructure:"SERVER_PORT"`
	DBHost              string        `mapstructure:"DB_HOST"`
	DBPort              string        `mapstructure:"DB_PORT"`
	DBUsername          string        `mapstructure:"DB_USERNAME"`
	DBPassword          string        `mapstructure:"DB_PASSWORD"`
	DBName              string        `mapstructure:"DB_DBNAME"`
	SSLMode             string        `mapstructure:"DB_SSLMODE"`
	TokenSecretKey      string        `mapstructure:"TOKEN_SECRET_KEY"`
	TokenValidationTime time.Duration `mapstructure:"TOKEN_VALIDATION_TIME"`
	MigrationUrl        string        `mapstructure:"MIGRATION_URL"`
	LoggerFilePath      string        `mapstructure:"LOGGER_FILE_PATH"`
	Environment         string        `mapstructure:"ENV"`
	ReadOnly            bool          `mapstructure:"READ_ONLY"`
	BasePath            string        `mapstructure:"BASE_PATH"`
}

// NewConfig creates a new configuration instance by reading from environment variables and .env file.
// It returns a Config instance and an error if any.
func NewConfig() (Config, error) {
	var config Config

	currentDir, err := os.Getwd()
	if err != nil {
		return Config{}, err
	}
	viper.AutomaticEnv()
	viper.SetConfigFile(currentDir + "/.env")
	err = viper.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	viper.Unmarshal(&config)
	return config, err
}
