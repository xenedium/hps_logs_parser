package parser

import (
	"github.com/xenedium/hps_logs_parser/iso8583/scanner"
	"github.com/xenedium/hps_logs_parser/iso8583/types"
	"os"
	"sync"
)

type parser struct {
	scanners             []*scanner.Scanner
	messages             []*types.Message
	Files                []*os.File
	ParsedDumpPostilions []*types.Message
	ParsedDumpXmls       []*types.Message
	ParsedDumpIsos       []*types.Message
	ParsedDumpTlvBuffers []*types.Message
}

func (p *parser) Parse() []*types.Message {
	waitGroup := sync.WaitGroup{}
	for _, file := range p.Files {
		waitGroup.Add(1)
		fileScanner := &scanner.Scanner{File: file}
		go func() {
			defer waitGroup.Done()
			fileScanner.Scan()
		}()
		p.scanners = append(p.scanners, fileScanner)
	}
	waitGroup.Wait()

	// each scanner will have a separate go routine
	// and each dump will have a separate go routine
	// this will make the parsing faster by a factor of the number of files

	for _, fileScanner := range p.scanners {
		waitGroup.Add(1)
		fileScanner := fileScanner
		go func() {
			defer waitGroup.Done()
			for _, postilionDump := range fileScanner.DumpPostilions {
				waitGroup.Add(1)
				postilionDump := postilionDump
				go func() {
					defer waitGroup.Done()
					p.parseDumpPostilion(postilionDump, fileScanner.File.Name())
				}()
			}
			for _, xmlDump := range fileScanner.DumpXmls {
				waitGroup.Add(1)
				xmlDump := xmlDump
				go func() {
					defer waitGroup.Done()
					p.parseDumpXml(xmlDump, fileScanner.File.Name())
				}()
			}
			for _, isoDump := range fileScanner.DumpIsos {
				waitGroup.Add(1)
				isoDump := isoDump
				go func() {
					defer waitGroup.Done()
					p.parseDumpIso(isoDump, fileScanner.File.Name())
				}()
			}
			for _, tlvBufferDump := range fileScanner.DumpTlvBuffers {
				waitGroup.Add(1)
				tlvBufferDump := tlvBufferDump
				go func() {
					defer waitGroup.Done()
					p.parseDumpTlvBuffer(tlvBufferDump, fileScanner.File.Name())
				}()
			}
		}()
	}

	waitGroup.Wait()

	return p.messages
}

type Parser = parser
