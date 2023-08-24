package main

import (
	"fmt"
	"github.com/xenedium/hps_logs_parser/iso8583/parser"
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
		if message.MTI.Class != 2 {
			continue
		}
		fmt.Println(message.Fields["037"].Value)
	}
}
