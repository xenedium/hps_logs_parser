package parser

import (
	"fmt"
	"github.com/xenedium/hps_logs_parser/scanner"
	"os"
	"path"
	"sync"

	protocolBuffer "github.com/xenedium/hps_logs_parser/server/gRPC"
)

type parser struct {
	scanners             []*scanner.Scanner
	Messages             []*protocolBuffer.Message
	Files                []*os.File
	ParsedDumpPostilions []*protocolBuffer.Message
	ParsedDumpXmls       []*protocolBuffer.Message
	ParsedDumpIsos       []*protocolBuffer.Message
	ParsedDumpTlvBuffers []*protocolBuffer.Message
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

func (p *parser) Parse(ignorePing bool) []*protocolBuffer.Message {
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
		for lineNumber, postilionDump := range fileScanner.DumpPostilions {
			p.parseDumpPostilion(&postilionDump, fileScanner.File.Name(), lineNumber)
		}
		for lineNumber, xmlDump := range fileScanner.DumpXmls {
			p.parseDumpXml(&xmlDump, fileScanner.File.Name(), lineNumber)
		}
		for lineNumber, isoDump := range fileScanner.DumpIsos {
			p.parseDumpIso(&isoDump, fileScanner.File.Name(), lineNumber)
		}
		for lineNumber, tlvBufferDump := range fileScanner.DumpTlvBuffers {
			p.parseDumpTlvBuffer(&tlvBufferDump, fileScanner.File.Name(), lineNumber)
		}

	}

	for _, parsedMessage := range p.ParsedDumpPostilions {
		if ignorePing && parsedMessage.Mti.Class == 8 {
			continue
		}
		p.Messages = append(p.Messages, parsedMessage)
	}
	for _, parsedMessage := range p.ParsedDumpXmls {
		if ignorePing && parsedMessage.Mti.Class == 8 {
			continue
		}
		p.Messages = append(p.Messages, parsedMessage)
	}
	for _, parsedMessage := range p.ParsedDumpIsos {
		if ignorePing && parsedMessage.Mti.Class == 8 {
			continue
		}
		p.Messages = append(p.Messages, parsedMessage)
	}
	for _, parsedMessage := range p.ParsedDumpTlvBuffers {
		if ignorePing && parsedMessage.Mti.Class == 8 {
			continue
		}
		p.Messages = append(p.Messages, parsedMessage)
	}

	return p.Messages
}

type Parser = parser
