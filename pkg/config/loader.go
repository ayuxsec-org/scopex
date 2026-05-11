package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func Load(path string) (Config, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return Config{}, ErrFileNotFound
		} else {
			return Config{}, fmt.Errorf("'os.Stat': %v", err)
		}
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to read config: %v", err)
	}
	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return Config{}, fmt.Errorf("failed to unmarshal config: %v", err)
	}
	return cfg, nil
}

func Create(path string) error {
	data, _ := yaml.Marshal(New())
	return os.WriteFile(path, data, 0644)
}
