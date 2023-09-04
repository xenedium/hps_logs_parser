package main

import (
	"github.com/xenedium/hps_logs_parser/iso8583parser/server"
	"log"
	"os"
)

func main() {
	if os.Getenv("GRPC_ADDRESS") == "" {
		err := os.Setenv("GRPC_ADDRESS", ":8080")
		if err != nil {
			log.Fatalf("Error setting GRPC_ADDRESS environment variable: %v", err)
		}
	}
	server.NewGRPCServer(os.Getenv("GRPC_ADDRESS"))
}

/*
tempDir := "logs"
	logParser := parser.NewParser(tempDir)
	logParser.Parse()

	for _, message := range logParser.Messages {
		_, ok := message.Fields["039"]
		if message.Mti.Class != 2 || !ok || message.Fields["043"] == nil {
			continue
		}
		fmt.Println("RNN:", message.Fields["037"].Value, "\nBITMAP: ", message.Bitmap, "\nRESPONSE CODE: ",
			message.Fields["039"].Value, "\nMTI: ", message.Mti.String(), "\nFILENAME: ", message.LogFileName,
			"\nRESPONSE MESSAGE: ", types.ResponseCodeMap[message.Fields["039"].Value],
			"\nCard acceptor name/location: ", message.Fields["043"].Value,
			"\n")
	}
*/
