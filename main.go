package main

import (
	"fmt"
	"github.com/xenedium/hps_logs_parser/iso8583/parser"
	"github.com/xenedium/hps_logs_parser/iso8583/types"
	"os"
	"path"
)

func main() {

	var logDir = "logs"

	files, err := os.ReadDir(logDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	var filesToParse []*os.File

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		f, err := os.Open(path.Join(logDir, file.Name()))
		if err != nil {
			fmt.Println(err)
			break
		}
		filesToParse = append(filesToParse, f)
	}

	logParser := parser.Parser{
		Files: filesToParse,
	}

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
}
