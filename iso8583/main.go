package main

import (
	"github.com/xenedium/iso8583parser/server"
	"os"
)

func main() {
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
