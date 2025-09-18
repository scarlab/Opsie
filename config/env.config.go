package config

import (
	"os"

	"github.com/joho/godotenv"
)

// envConfig holds all environment configuration variables
// that your application needs to run.
type envConfig struct {
	GoEnv      	string
	Addr      	string
	JwtSecret 	string

	MainDBPath 	string
	AgentDBPath string
}


// ENV (singleton) is a globally accessible variable
var ENV = loadEnv()
var IsDev = loadEnv().GoEnv == "development"



// loadEnv loads environment variables into an EnvConfig struct.
// If an environment variable is not found, it uses the provided default value.
func loadEnv() *envConfig {
	// Load the env variables.
	godotenv.Load()
	
	return &envConfig{
		GoEnv:      	getEnv("GO_ENV", "development"),
		Addr:      		getEnv("ADDR", ":3905"),
		JwtSecret: 		getEnv("JWT_SECRET", "9as879das7d86$a87das89nd89asd7as+6da9snd9asd"),

		MainDBPath:  	getEnv("MAIN_DB_PATH", "/var/lib/watchtower/database.sqlite"),
		AgentDBPath:  	getEnv("AGENT_DB_PATH", "/var/lib/watchtower/wtagent.db"),
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


