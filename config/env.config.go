package config

import (
	"os"

	"github.com/joho/godotenv"
)

// TEnvConfig holds all environment configuration variables
// that your application needs to run.
type TEnvConfig struct {
	Env      	string
	Addr      	string
	JwtSecret 	string

	// Database [PostgreSQL]
	PG_User string
	PG_Password string
	PG_Host string
	PG_Port string
	PG_Database string

	DevUIHost string

	// --------------------------------------- Agent ---------------------------------------
	ServerHost 	string
}


// loadEnv loads environment variables into an EnvConfig struct.
// If an environment variable is not found, it uses the provided default value.
func loadEnv() *TEnvConfig {
	// Load the env variables.
	godotenv.Load()
	
	return &TEnvConfig{
		Env:      			getEnv("OPSIE_ENV", "development"),
		Addr:      			getEnv("ADDR", ":3905"),
		JwtSecret: 			getEnv("JWT_SECRET", "9as879das7d8a9snd9asd"),

		// Database [PostgreSQL]
		PG_User:   			getEnv("POSTGRES_USER", "postgres"),
		PG_Password: 		getEnv("POSTGRES_PASSWORD", "postgres_password"),
		PG_Host:   			getEnv("POSTGRES_HOST", "127.0.0.1"),
		PG_Port:   			getEnv("POSTGRES_PORT", "5432"),
		PG_Database:     	getEnv("POSTGRES_DB", "opsie"),

		DevUIHost: 			getEnv("DEV_UI_HOST", "opsie-ui-1"),


		
		// --------------------------------------- Agent ---------------------------------------
		ServerHost: 		getEnv("SERVER_HOST", "localhost:3905"),

	}
}

// getEnv retrieves the value of an environment variable by its key.
// If the variable is not set, it returns the fallback value instead.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}


