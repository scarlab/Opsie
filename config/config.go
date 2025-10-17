package config

import "log"

// ENV is a global singleton that holds all environment-based configuration.
// It’s initialized once when the package is imported.
var ENV = loadEnv()

// AppConfig holds the static metadata loaded from app.config.yaml.
//
// Panics at startup if the file is missing or malformed —
// this is intentional, as app metadata should never fail silently.
var AppConfig *TAppConfig

// IsDev is a global convenience flag for development-only logic.
// e.g., if config.IsDev { enableDebugLogging() }
var IsDev bool

// init runs automatically when the package is imported.
// It initializes global configuration objects.
func init() {
	var err error
	AppConfig, err = loadConfig("app.config.yaml")
	if err != nil {
		log.Panicf("failed to load app.config.yaml: %v", err)
	}

	IsDev = ENV.GoEnv == "development"
}
