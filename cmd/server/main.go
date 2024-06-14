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

	openAIService := services.NewOpenAIService(&cfg.OpenAIConfig)
	graphQLService := services.NewGraphQLService(&cfg.GraphQLConfig)
	lessonService := services.NewLessonService(openAIService)

	server.Start(openAIService, graphQLService, lessonService, cfg.ServerConfig.Port)
}
