package parser

import (
	"fmt"
	"github.com/xenedium/hps_logs_parser/iso8583/scanner"
	"github.com/xenedium/hps_logs_parser/iso8583/types"
	"os"
	"sync"
)

type parser struct {
	scanners []*scanner.Scanner
	messages []types.Message
	Files    []*os.File
}

func (p *parser) Parse() []types.Message {
	waitGroup := sync.WaitGroup{}
	for _, file := range p.Files {
		waitGroup.Add(1)
		fileScanner := new(scanner.Scanner)
		fileScanner.File = file
		go func() {
			defer waitGroup.Done()
			fileScanner.Scan()
		}()
		p.scanners = append(p.scanners, fileScanner)
	}
	waitGroup.Wait()
	for _, fileScanner := range p.scanners {
		for _, postilionDump := range fileScanner.DumpPostilions {
			// TODO: parse postilion dump
			fmt.Println(postilionDump)
		}
		for _, xmlDump := range fileScanner.DumpXmls {
			// TODO: parse xml dump
			fmt.Println(xmlDump)
		}
		for _, isoDump := range fileScanner.DumpIsos {
			// TODO: parse iso dump
			fmt.Println(isoDump)
		}
		for _, tlvBufferDump := range fileScanner.DumpTlvBuffers {
			// TODO: parse tlv buffer dump
			fmt.Println(tlvBufferDump)
		}
	}

	return p.messages
}

type Parser = parser
