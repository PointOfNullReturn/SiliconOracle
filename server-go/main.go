package main

import (
	"fmt"
	"log"
	"server/server/routes"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func LoadConfig(path string) (config Config, err error) {

	// Set default values for the configuration
	viper.SetDefault("server.host", "localhost")
	viper.SetDefault("server.port", "8080")

	// Set the path to the configuration file
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the configuration
	err = viper.ReadInConfig()
	if err != nil {

	} else {
		log.Printf("Using config file: %s", viper.ConfigFileUsed())
	}

	// Unmarshal the configuration into the Config struct
	err = viper.Unmarshal(&config) // Unmarshal config into the Config struct
	if err != nil {
		return Config{}, fmt.Errorf("unable to decode into struct, %v", err)
	}

	// Log loaded conifguration
	log.Printf("Loaded configuration: %+v", config)

	return config, nil
}

func main() {
	// Load configuration from config.yaml in the same directory.
	appConfig, err := LoadConfig(".")

	if err != nil {
		log.Fatalf("FATAL: Error loading config: %v", err)
	}

	// Initialize the Gin Engine
	router := gin.Default()

	router = routes.ConfigRoutes(router) // Configure the routes

	// Construct the server address from the config
	listenAddr := fmt.Sprintf("%s:%s", appConfig.Server.Host, appConfig.Server.Port)
	log.Printf("Starting server on %s...\n", listenAddr)

	// Start the server
	err = router.Run(listenAddr)

	if err != nil {
		log.Fatalf("FATAL: Failed to start server: %v", err)
	} else {
		log.Printf("INFO: Server started on %s", listenAddr)
	}

}
