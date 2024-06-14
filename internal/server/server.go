package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bamboo/internal/app/resolvers"
	"bamboo/internal/app/services"

	"github.com/graph-gophers/graphql-go/relay"
)

func Start(
	openAIService *services.OpenAIService,
	graphQLService *services.GraphQLService,
	lessonService *services.LessonService,
	port string,
) {
	server := setupServer(openAIService, graphQLService, lessonService, port)

	log.Println("Bamboo server is starting...")

	go func() {
		log.Printf("Bamboo server started successfully on port: %s", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
		log.Println("Bamboo server stopped accepting requests")
	}()

	handleShutdown(server)
}

func setupServer(
	openAIService *services.OpenAIService,
	graphQLService *services.GraphQLService,
	lessonService *services.LessonService,
	port string,
) *http.Server {
	schema, err := graphQLService.LoadSchema(&resolvers.LessonResolver{
		LessonService: lessonService,
	})
	if err != nil {
		log.Fatalf("Failed to create schema: %v", err)
	}

	http.Handle("/graphql", Logging(&relay.Handler{Schema: schema}))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	return server
}

func handleShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("Shutdown received. Stopping Bamboo server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed: %v", err)
	}

	log.Println("Bamboo server stopped gracefully.")
}
