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

	server.NewServer(os.Getenv("BACKEND_ADDRESS")).Run()
}
