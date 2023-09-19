package parser

import (
	"fmt"
	"github.com/xenedium/hps_logs_parser/iso8583parser/scanner"
	"math"
	"os"
	"path"
	"sort"
	"strconv"
	"sync"

	protocolBuffer "github.com/xenedium/hps_logs_parser/gRPC"
)

type parser struct {
	scanners             []*scanner.Scanner
	Messages             []*protocolBuffer.Message
	Files                []*os.File
	ParsedDumpPostilions []*protocolBuffer.Message
	ParsedDumpXmls       []*protocolBuffer.Message
	ParsedDumpIsos       []*protocolBuffer.Message
	ParsedDumpTlvBuffers []*protocolBuffer.Message
	ParsedDumpBuffers    []*protocolBuffer.Message
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
		for lineNumber, isoDump := range fileScanner.DumpBuffers {
			p.parseDumpBuffer(&isoDump, fileScanner.File.Name(), lineNumber)
		}
	}

	tempArray := make([]*protocolBuffer.Message, 0)

	for _, parsedMessage := range p.ParsedDumpPostilions {
		if ignorePing && parsedMessage.Mti.Class == 8 {
			continue
		}
		tempArray = append(tempArray, parsedMessage)
	}
	for _, parsedMessage := range p.ParsedDumpXmls {
		if ignorePing && parsedMessage.Mti.Class == 8 {
			continue
		}
		tempArray = append(tempArray, parsedMessage)
	}
	for _, parsedMessage := range p.ParsedDumpIsos {
		if ignorePing && parsedMessage.Mti.Class == 8 {
			continue
		}
		tempArray = append(tempArray, parsedMessage)
	}
	for _, parsedMessage := range p.ParsedDumpTlvBuffers {
		if ignorePing && parsedMessage.Mti.Class == 8 {
			continue
		}
		tempArray = append(tempArray, parsedMessage)
	}

	for index, message := range tempArray {
		for _, bufferDump := range p.ParsedDumpBuffers {
			if message.LogFileName == bufferDump.LogFileName &&
				message.ThreadId == bufferDump.ThreadId &&
				diff(message.Timestamp, bufferDump.Timestamp) < 10 {
				tempArray[index].Raw = bufferDump.Raw
				break
			}
		}
	}

	for _, parsedMessage := range tempArray {
		duplicateMessage, index := isDuplicate(p.Messages, parsedMessage)
		if duplicateMessage != nil {
			p.Messages[index] = mergeMessages(duplicateMessage, parsedMessage)
			continue
		}
		p.Messages = append(p.Messages, parsedMessage)
	}

	sort.Slice(p.Messages, func(i, j int) bool {
		return p.Messages[i].Timestamp < p.Messages[j].Timestamp
	})

	return p.Messages
}

func isDuplicate(messages []*protocolBuffer.Message, parsedMessage *protocolBuffer.Message) (*protocolBuffer.Message, int) {
	for index, message := range messages {
		if parsedMessage.ThreadId == message.ThreadId &&
			parsedMessage.Mti != nil && message.Mti != nil &&
			parsedMessage.Mti.Class == message.Mti.Class &&
			parsedMessage.Mti.Function == message.Mti.Function &&
			parsedMessage.Mti.Origin == message.Mti.Origin &&
			parsedMessage.Mti.Version == message.Mti.Version &&
			parsedMessage.Fields["037"] != nil && message.Fields["037"] != nil &&
			parsedMessage.Fields["037"].Value == message.Fields["037"].Value &&
			parsedMessage.Bitmap == message.Bitmap {
			// DO I REALLY NEED TO CHECK THE TIMESTAMP?
			// DOING SO MIGHT FLAG A MESSAGE AS DUPLICATE WHEN IT IS NOT
			/*
				difference, err := diff(parsedMessage.Timestamp, message.Timestamp)
				if err != nil {
					return nil, 0
				}
				if difference < 60 {
					return message, index
				}
			*/
			return message, index
		}
	}
	return nil, 0
}

func diff(a, b string) uint64 {
	x, err := strconv.Atoi(a)
	if err != nil {
		return 0
	}
	y, err := strconv.Atoi(b)
	if err != nil {
		return 0
	}
	return uint64(math.Abs(float64(x) - float64(y)))
}

func mergeMessages(a, b *protocolBuffer.Message) *protocolBuffer.Message {
	// the mti, thread id, and fld 37 are the same
	// we need to merge the fields, and the raw
	for key, value := range b.Fields {
		if a.Fields[key] == nil {
			a.Fields[key] = value
		}
	}
	if a.Raw == "" {
		a.Raw = b.Raw
	}
	if a.Raw != "" && b.Raw != "" && a.Raw != b.Raw {
		panic("Both raws are not empty")
	}
	return a
}

type Parser = parser
