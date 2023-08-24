package parser

import (
	"github.com/xenedium/hps_logs_parser/iso8583/types"
	"strings"
)

func (p *parser) parseDumpPostilion(dumpPostilion *string, fileName string) {
	parsedMessage := &types.Message{
		LogFileName: fileName,
	}

	dumpPostilionLines := strings.Split(*dumpPostilion, "\n")
	parsedMessage.Bitmap = strings.ReplaceAll(dumpPostilionBitMapRegexMatcher.FindStringSubmatch(dumpPostilionLines[1])[1], " ", "")
	mtiStr := dumpPostilionMTIRegexMatcher.FindStringSubmatch(dumpPostilionLines[3])[1]
	parsedMessage.MTI = types.MTI{
		Version:  mtiStr[0] - '0',
		Class:    mtiStr[1] - '0',
		Function: mtiStr[2] - '0',
		Origin:   mtiStr[3] - '0',
	}

	parsedMessage.Fields = make(map[string]types.Field)

	for _, line := range dumpPostilionLines[6:] {
		if len(line) == 0 || len(dumpPostilionFieldRegexMatcher.FindStringSubmatch(line)) != 3 {
			continue
		}
		fieldNumber := dumpPostilionFieldRegexMatcher.FindStringSubmatch(line)[1]
		fieldValue := dumpPostilionFieldRegexMatcher.FindStringSubmatch(line)[2]
		parsedMessage.Fields[fieldNumber] = types.Field{
			Value: strings.ReplaceAll(fieldValue, "]", ""),
		}
	}

	p.ParsedDumpPostilions = append(p.ParsedDumpPostilions, parsedMessage)
}

func (p *parser) parseDumpXml(dumpXml *string, fileName string) {
	parsedMessage := &types.Message{
		LogFileName: fileName,
	}

	dumpXmlLines := strings.Split(*dumpXml, "\n")
	parsedMessage.Bitmap = strings.ReplaceAll(dumpXmlBitMapRegexMatcher.FindStringSubmatch(dumpXmlLines[0])[1], " ", "")
	mtiStr := dumpXmlMTIRegexMatcher.FindStringSubmatch(dumpXmlLines[2])[1]
	parsedMessage.MTI = types.MTI{
		Version:  mtiStr[0] - '0',
		Class:    mtiStr[1] - '0',
		Function: mtiStr[2] - '0',
		Origin:   mtiStr[3] - '0',
	}

	parsedMessage.Fields = make(map[string]types.Field)
	multilineFields := make([]string, 0)

	for _, line := range dumpXmlLines[6:] {
		if len(line) == 0 {
			continue
		}
		if len(dumpXmlClosingMessageInfoTagRegexMatcher.FindStringSubmatch(line)) != 0 {
			break
		}
		if len(dumpXmlFieldRegexMatcher.FindStringSubmatch(line)) != 3 {
			multilineFields = append(multilineFields, line)
			continue
		}
		fieldNumber := dumpXmlFieldRegexMatcher.FindStringSubmatch(line)[1]
		fieldValue := dumpXmlFieldRegexMatcher.FindStringSubmatch(line)[2]
		parsedMessage.Fields[fieldNumber] = types.Field{
			Value: strings.ReplaceAll(fieldValue, "]", ""),
		}
	}

	// some fields are multiline, and we need to parse them
	currentFieldNumber := ""
	for _, line := range multilineFields {
		if len(line) == 0 {
			continue
		}
		if len(dumpXmlMultiLineFieldHeaderRegexMatcher.FindStringSubmatch(line)) != 0 {
			currentFieldNumber = dumpXmlMultiLineFieldHeaderRegexMatcher.FindStringSubmatch(line)[1]
			parsedMessage.Fields[currentFieldNumber] = types.Field{
				Value: dumpXmlMultiLineFieldHeaderRegexMatcher.FindStringSubmatch(line)[2],
			}
			continue
		}
		if len(dumpXmlMultiLineFieldClosingTagRegexMatcher.FindStringSubmatch(line)) != 0 {
			parsedMessage.Fields[currentFieldNumber] = types.Field{
				Value: parsedMessage.Fields[currentFieldNumber].Value + dumpXmlMultiLineFieldClosingTagRegexMatcher.FindStringSubmatch(line)[1],
			}
			continue
		}
		if len(dumpXmlMultiLineFieldRegexMatcher.FindStringSubmatch(line)) != 0 {
			parsedMessage.Fields[currentFieldNumber] = types.Field{
				Value: parsedMessage.Fields[currentFieldNumber].Value + dumpXmlMultiLineFieldRegexMatcher.FindStringSubmatch(line)[1],
			}
			continue
		}
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
