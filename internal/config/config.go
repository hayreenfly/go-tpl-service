package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config represents the application configuration structure
type Config struct {
    App      AppConfig
    Database DatabaseConfig
}

// AppConfig holds app-level configuration
type AppConfig struct {
    Name        string
    Port        int
    Environment string
}

// DatabaseConfig holds database-related configuration
type DatabaseConfig struct {
    Host     string
    Port     int
    Username string
    Password string
    Name     string
}

// LoadConfig reads configuration from file or environment variables
func LoadConfig() (*Config, error) {
    var config Config

    viper.SetConfigName("config")   // Name of config file (without extension)
    viper.SetConfigType("yaml")     // Type of config file (yaml)
    viper.AddConfigPath("./internal/config") // Path to look for the config file
    viper.AutomaticEnv()            // Automatically override values from the environment variables

    // Set environment variable prefix (optional)
    viper.SetEnvPrefix("MYAPP")

    // Read the config file
    if err := viper.ReadInConfig(); err != nil {
        log.Printf("Error reading config file, %s", err)
        return nil, err
    }

    // Unmarshal the config into the struct
    if err := viper.Unmarshal(&config); err != nil {
        log.Printf("Unable to decode into struct, %v", err)
        return nil, err
    }

    return &config, nil
}


// GetDSN constructs the PostgreSQL DSN (Data Source Name)
func (dbConfig *DatabaseConfig) GetDSN() string {
	// Format the DSN string: "postgres://username:password@host:port/dbname"
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)
}