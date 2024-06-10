package main

import (
	"log"

	"bamboo/internal/app/services"
	"bamboo/internal/config"
	"bamboo/internal/server"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Could not load config from config yaml: %v", err)
	}

	openAIService := services.NewOpenAIService(cfg.OpenAI.APIKey)

	server.Start(openAIService, cfg.Server.Port)
}
