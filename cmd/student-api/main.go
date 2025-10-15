package main

import (
	"fmt"
	"github/rithikrajkumar4/student-backend-go/internal/config"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Student API Started")
	// Load Config
	cfg := config.MustLoad()
	// Logger
	// Database setup
	// router setup
	router := http.NewServeMux()

	router.HandleFunc("GET /", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte("Welcome to Student API's"))
	})
	// server setup
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Printf("Server Started.... %s", cfg.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to Start Server")
	}

}
