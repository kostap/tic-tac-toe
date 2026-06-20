package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/kostap/tic-tac-toe/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	gameHandler := handler.NewGameHandler()
	mux := http.NewServeMux()

	// API endpoints
	mux.HandleFunc("/api/new-game", gameHandler.NewGame)
	mux.HandleFunc("/api/move", gameHandler.MakeMove)
	mux.HandleFunc("/api/state", gameHandler.GetState)

	// Main page
	fs := http.FileServer(http.Dir("static"))
	mux.Handle("/", fs)

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
