package parser

import (
	"errors"
	"fmt"
	"github.com/xenedium/hps_logs_parser/iso8583parser/scanner"
	"math"
	"os"
	"path"
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

	for _, parsedMessage := range tempArray {
		isDupli, _ := isDuplicate(p.Messages, parsedMessage)
		if isDupli != nil {
			// TODO: MUST MERGE THE TWO MESSAGES

			continue
		}
		p.Messages = append(p.Messages, parsedMessage)
	}

	return p.Messages
}

// TODO: MUST MERGE RAW DUMPS WITH PARSED DUMPS BUT HOW ?????, THIS MUST BE DONE BEFORE REDUCING PHASE
func isDuplicate(messages []*protocolBuffer.Message, parsedMessage *protocolBuffer.Message) (*protocolBuffer.Message, int) {
	for index, message := range messages {
		if parsedMessage.ThreadId == message.ThreadId &&
			parsedMessage.Mti != nil && message.Mti != nil &&
			parsedMessage.Mti.Class == message.Mti.Class &&
			parsedMessage.Mti.Function == message.Mti.Function &&
			parsedMessage.Mti.Origin == message.Mti.Origin &&
			parsedMessage.Mti.Version == message.Mti.Version &&
			parsedMessage.Fields["037"] != nil && message.Fields["037"] != nil &&
			parsedMessage.Fields["037"].Value == message.Fields["037"].Value {
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

func diff(a, b string) (uint64, error) {
	x, err := strconv.Atoi(a)
	if err != nil {
		return 0, errors.New("error converting string to int")
	}
	y, err := strconv.Atoi(b)
	if err != nil {
		return 0, errors.New("error converting string to int")
	}
	return uint64(math.Abs(float64(x) - float64(y))), nil
}

type Parser = parser
