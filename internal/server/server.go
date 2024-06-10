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

func Start(openAIService *services.OpenAIService, port string) {
	server := setupServer(openAIService, port)

	log.Println("Bamboo server is starting...")

	go func() {
		log.Printf("Bamboo server started successfully on port: %s.", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", server.Addr, err)
		}
		log.Println("Bamboo server stopped accepting requests.")
	}()

	handleShutdown(server)
}

func setupServer(openAIService *services.OpenAIService, port string) *http.Server {
	resolver := &resolvers.OpenAIResolver{OpenAIService: openAIService}
	schema := resolvers.NewSchema(resolver)

	http.Handle("/query", loggingMiddleware(&relay.Handler{Schema: schema}))

	server := &http.Server{
		Addr:    ":" + port,
		Handler: nil,
	}

	log.Printf("Bamboo server configured to listen on port: %s\n", port)
	return server
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Received %s request for %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
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
