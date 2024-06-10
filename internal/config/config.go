package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	OpenAI OpenAIConfig `yaml:"openai"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

type OpenAIConfig struct {
	APIKey string `yaml:"apiKey"`
}

func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	log.Printf("Configuration loaded from %s\n", filename)
	return &config, nil
}
