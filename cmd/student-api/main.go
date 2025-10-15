package main

import (
	"context"
	"fmt"
	"github/rithikrajkumar4/student-backend-go/internal/config"
	"github/rithikrajkumar4/student-backend-go/internal/http/handlers/student"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("Student API Started")
	// Load Config
	cfg := config.MustLoad()
	// Logger
	// Database setup
	// router setup
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())

	// server setup
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server Started", slog.String("address", cfg.Addr))
	fmt.Printf("Server Started.... %s", cfg.Addr)

	done := make(chan os.Signal, 1) // Create a channel

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Failed to Start Server")
		}
	}()

	<-done

	slog.Info("Shuting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("Failed to Shutdown the server!", slog.String("error", err.Error()))
	}

	slog.Info("Server Shut down Successfully")
}
