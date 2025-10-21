package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// TAppConfig defines static application-level metadata.
// This YAML file should not contain any secrets — only static info like app name, version, etc.
type TAppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	SessionDays int `yaml:"session_days"`

	// Directory & Files
	LogDir string `yaml:"log_dir"`
	StaticDir string `yaml:"static_dir"`
	DataDir string `yaml:"data_dir"`

	ConfigFile string `yaml:"config_file"`
	RuntimeFile string `yaml:"runtime_dir"`
	Binary string `yaml:"binary"`
}


// loadAppConfig reads and parses the YAML configuration file located at `path`.
// It returns a TAppConfig instance or an error if loading fails.
func loadAppConfig() (*TAppConfig, error) {
	data, err := os.ReadFile("app.config.yaml")
	if err != nil {
		return nil, err
	}

	var cfg TAppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
