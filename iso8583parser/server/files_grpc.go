package server

import (
	"context"
	"errors"
	"github.com/xenedium/iso8583parser/parser"
	protocolBuffer "github.com/xenedium/iso8583parser/server/gRPC"
	"log"
	"os"
	"path"
)

func (s *gRPCServer) FilesParse(ctx context.Context, req *protocolBuffer.FilesRequest) (*protocolBuffer.Response, error) {
	log.Println("Received FilesParse request")

	if len(req.Files) == 0 {
		log.Printf("No files specified\n")
		return nil, errors.New("no files specified")
	}

	log.Printf("Received %d File(s)\n", len(req.Files))

	log.Printf("Creating temp dir\n")
	tempDir, err := os.MkdirTemp(os.TempDir(), "*")
	if err != nil {
		log.Printf("Failed to create temp dir: %v\n", err)
		return nil, err
	}

	for _, file := range req.Files {
		log.Printf("Saving file: %v\n", file.Name)
		err := os.WriteFile(path.Join(tempDir, file.Name), []byte(file.Content), 0644)
		if err != nil {
			log.Printf("Failed to save file: %v\n", err)
			return nil, err
		}
	}

	log.Printf("Parsing files\n")
	logParser := parser.NewParser(tempDir)
	logParser.Parse(true)

	log.Printf("Sending response\n")
	return &protocolBuffer.Response{Messages: logParser.Messages}, nil
}
