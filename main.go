package main

import (
	"github.com/xenedium/hps_logs_parser/server"
)

func main() {
	server.NewServer(":8080", "123").Run()
}

/*
	logParser := parser.NewParser(tempDir)
	logParser.Parse()

	for _, message := range logParser.Messages {
		_, ok := message.Fields["039"]
		if message.MTI.Class != 2 || !ok {
			continue
		}
		fmt.Println("RNN:", message.Fields["037"].Value, "\nBITMAP: ", message.Bitmap, "\nRESPONSE CODE: ",
			message.Fields["039"].Value, "\nMTI: ", message.MTI.String(), "\nFILENAME: ", message.LogFileName,
			"\nRESPONSE MESSAGE: ", types.ResponseCodeMap[message.Fields["039"].Value],
			"\nCard acceptor name/location: ", message.Fields["043"].Value,
			"\n")
	}
*/
