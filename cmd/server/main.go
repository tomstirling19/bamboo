package main

import (
	"bamboo/internal/app/services"
	"bamboo/internal/config"
	"bamboo/internal/server"
	"log"
)

func main() {
	cfg, err := config.LoadConfig("config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

    openAIService := services.NewOpenAIService(cfg.OpenAI.APIKey)
    lessonService := services.NewLessonService()

	    server.Start(openAIService, lessonService, cfg.Server.Port)
	}
