package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// TenantBasePath Default value
var (
	TenantBasePath = "./sites"
)

// LoadConfig loads .env based on GO_ENV (local, docker, production)
func LoadConfig() {
	env := os.Getenv("GO_ENV")

	var envFile string

	switch env {
	case "production":
		envFile = ".env.production"
	case "docker":
		envFile = ".env.docker"
	case "local":
		envFile = ".env.local"
	default:
		envFile = ".env"
	}

	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("Could not load %s: %v\n", envFile, err)
	} else {
		log.Println("Loaded env file:", envFile)
	}

	// Set config variable if present
	if val := os.Getenv("TENANT_PATH"); val != "" {
		TenantBasePath = val
		log.Println("Tenant base path:", TenantBasePath)
	}
}
