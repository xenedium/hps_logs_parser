package parser

import (
	"github.com/xenedium/hps_logs_parser/iso8583/types"
)

func (p *parser) parseDumpPostilion(dumpPostilion *string, fileName string) {
	parsedMessage := &types.Message{
		LogFileName: fileName,
	}

	p.ParsedDumpPostilions = append(p.ParsedDumpPostilions, parsedMessage)

}

func (p *parser) parseDumpXml(dumpXml *string, fileName string) {
	parsedMessage := &types.Message{
		LogFileName: fileName,
	}

	p.ParsedDumpXmls = append(p.ParsedDumpXmls, parsedMessage)
}

func (p *parser) parseDumpIso(dumpIso *string, fileName string) {
	parsedMessage := &types.Message{
		LogFileName: fileName,
	}

	p.ParsedDumpIsos = append(p.ParsedDumpIsos, parsedMessage)
}

func (p *parser) parseDumpTlvBuffer(dumpTlvBuffer *string, fileName string) {
	parsedMessage := &types.Message{
		LogFileName: fileName,
	}

	p.ParsedDumpTlvBuffers = append(p.ParsedDumpTlvBuffers, parsedMessage)
}
