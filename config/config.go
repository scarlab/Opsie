package config

import (
	"log"
)

// ENV is a global singleton that holds all environment-based configuration.
// It’s initialized once when the package is imported.
var ENV = loadEnv()

// App holds the static metadata loaded from app.config.yaml.
//
// Panics at startup if the file is missing or malformed —
// this is intentional, as app metadata should never fail silently.
var App *TAppConfig

// IsDev is a global convenience flag for development-only logic.
// e.g., if config.IsDev { enableDebugLogging() }
var IsDev bool
var IsProd bool


// Files and dirs
// Default relative path for dev
var DefaultStaticDir = "uploads"


// init runs automatically when the package is imported.
// It initializes global configuration objects.
func init() {
	var err error
	App, err = loadAppConfig()
	if err != nil {
		log.Panicf("failed to load app.config.yaml: %v", err)
	}

	IsDev = ENV.Env == "development"
	IsProd = ENV.Env == "production"

	// Production Static dir
	if IsProd {
		DefaultStaticDir = App.StaticDir
	}
}
