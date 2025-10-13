package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// tAppConfig defines static application-level metadata.
// This YAML file should not contain any secrets — only static info like app name, version, etc.
type tAppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
}


// loadConfig reads and parses the YAML configuration file located at `path`.
// It returns a TAppConfig instance or an error if loading fails.
func loadConfig(path string) (*tAppConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg tAppConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
