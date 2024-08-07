package config

import (
	"errors"
	"log"
	"os"
	"sync"

	"github.com/goccy/go-yaml"
)

type Config struct {
	Port       int               `yaml:"port"`
	ClientPort int               `yaml:"client_port"`
	Version    map[string]string `yaml:"version"`
}

var ConfigData *Config = &Config{}

var mutex = &sync.Mutex{}

func LoadConfig() (*Config, error) {
	filePath, ok := os.LookupEnv("MVN_CONFIG_FILE")
	if !ok {
		log.Fatal("MVN_CONFIG_FILE is not set")
		return nil, errors.New("MVN_CONFIG_FILE is not set")
	}
	yamlFile, error := os.ReadFile(filePath)

	if error != nil {
		log.Fatal("Error Fetching Config ", error)
		return nil, error
	}

	mutex.Lock()
	err := yaml.Unmarshal(yamlFile, ConfigData)
	if err != nil {
		log.Fatal("Error Loading Config ", err)
		return nil, err
	}
	mutex.Unlock()

	return ConfigData, nil
}
