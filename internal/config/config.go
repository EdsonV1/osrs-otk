package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Assets AssetsConfig `yaml:"assets"`
	CORS   CORSConfig   `yaml:"cors"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
	Host string `yaml:"host"`
}

type AssetsConfig struct {
	SkillDataPath string `yaml:"skill_data_path"`
	ImagesPath    string `yaml:"images_path"`
}

type CORSConfig struct {
	AllowedOrigins []string `yaml:"allowed_origins"`
	AllowedMethods []string `yaml:"allowed_methods"`
	AllowedHeaders []string `yaml:"allowed_headers"`
}

// Load loads configuration from environment-specific YAML file
func Load() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	configPath := filepath.Join("internal", "config", "environments", fmt.Sprintf("%s.yaml", env))

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file %s: %w", configPath, err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &config, nil
}

// GetAddr returns the full server address
func (c *Config) GetAddr() string {
	return c.Server.Host + ":" + c.Server.Port
}
