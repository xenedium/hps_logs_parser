package parser

import (
	"fmt"
	"github.com/xenedium/hps_logs_parser/iso8583/scanner"
	"github.com/xenedium/hps_logs_parser/iso8583/types"
	"os"
	"path"
	"sync"
)

type parser struct {
	scanners             []*scanner.Scanner
	Messages             []*types.Message
	Files                []*os.File
	ParsedDumpPostilions []*types.Message
	ParsedDumpXmls       []*types.Message
	ParsedDumpIsos       []*types.Message
	ParsedDumpTlvBuffers []*types.Message
}

func NewParser(logDir string) *Parser {
	files, err := os.ReadDir(logDir)
	if err != nil {
		fmt.Println(err)
		return nil
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

	return &parser{
		Files: filesToParse,
	}
}

func (p *parser) Parse() []*types.Message {
	defer func() {
		for _, file := range p.Files {
			_ = file.Close()
		}
	}()
	scanWaitGroup := sync.WaitGroup{}
	for _, file := range p.Files {
		scanWaitGroup.Add(1)
		fileScanner := &scanner.Scanner{File: file}
		go func() {
			defer scanWaitGroup.Done()
			fileScanner.Scan()
		}()
		p.scanners = append(p.scanners, fileScanner)
	}
	scanWaitGroup.Wait()

	for _, fileScanner := range p.scanners {
		for _, postilionDump := range fileScanner.DumpPostilions {
			p.parseDumpPostilion(&postilionDump, fileScanner.File.Name())
		}
		for _, xmlDump := range fileScanner.DumpXmls {
			p.parseDumpXml(&xmlDump, fileScanner.File.Name())
		}
		for _, isoDump := range fileScanner.DumpIsos {
			p.parseDumpIso(&isoDump, fileScanner.File.Name())
		}
		for _, tlvBufferDump := range fileScanner.DumpTlvBuffers {
			p.parseDumpTlvBuffer(&tlvBufferDump, fileScanner.File.Name())
		}

	}

	for _, parsedMessage := range p.ParsedDumpPostilions {
		p.Messages = append(p.Messages, parsedMessage)
	}
	for _, parsedMessage := range p.ParsedDumpXmls {
		p.Messages = append(p.Messages, parsedMessage)
	}
	for _, parsedMessage := range p.ParsedDumpIsos {
		p.Messages = append(p.Messages, parsedMessage)
	}
	for _, parsedMessage := range p.ParsedDumpTlvBuffers {
		p.Messages = append(p.Messages, parsedMessage)
	}

	return p.Messages
}

type Parser = parser
