package main

import (
	"github.com/xenedium/hps_logs_parser/iso8583backend/server"
	"log"
	"os"
)

func main() {
	if os.Getenv("BACKEND_ADDRESS") == "" {
		err := os.Setenv("BACKEND_ADDRESS", ":8000")
		if err != nil {
			log.Fatalf("Error setting BACKEND_ADDRESS environment variable: %v", err)
		}
	}

	if os.Getenv("GRPC_ADDRESS") == "" {
    log.Println("GRPC_ADDRESS not found, using default value")
		err := os.Setenv("GRPC_ADDRESS", "127.0.0.1:8080")
		if err != nil {
			log.Fatalf("Error setting GRPC_ADDRESS environment variable: %v", err)
		}
	}

	if os.Getenv("REDIS_ADDRESS") == "" {
    log.Println("REDIS_ADDRESS not found, using default value")
		err := os.Setenv("REDIS_ADDRESS", "127.0.0.1:6379")
		if err != nil {
			log.Fatalf("Error setting REDIS_ADDRESS environment variable: %v", err)
		}
	}

	server.NewServer(os.Getenv("BACKEND_ADDRESS")).Run()
}
