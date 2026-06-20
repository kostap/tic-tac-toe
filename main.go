package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Tic-Tac-Toe Game Server")
		if err != nil {
			log.Printf("Error writing response: %v", err)
		}
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// Run server in goroutine
	go func() {
		fmt.Println("Server starting on :8080...")

		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Channel for OS signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block
	sig := <-quit
	fmt.Println("Received signal: %v. Shutting down...", sig)

	// 30 second for finishing active connections
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	fmt.Println("Server stopped successfully")
}
