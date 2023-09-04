package parser

import (
	"strconv"
	"strings"

	protocolBuffer "github.com/xenedium/hps_logs_parser/gRPC"
)

func (p *parser) parseDumpPostilion(dumpPostilion *string, fileName string, lineNumber int) {
	parsedMessage := &protocolBuffer.Message{
		LogFileName: fileName,
		LineNumber:  strconv.Itoa(lineNumber),
	}

	dumpPostilionLines := strings.Split(*dumpPostilion, "\n")
	parsedMessage.Bitmap = strings.ReplaceAll(dumpPostilionBitMapRegexMatcher.FindStringSubmatch(dumpPostilionLines[1])[1], " ", "")
	parsedMessage.Mti = &protocolBuffer.MTI{}
	parsedMessage.Fields = make(map[string]*protocolBuffer.Field)
	//parsedMessage.SetMTI(dumpPostilionMTIRegexMatcher.FindStringSubmatch(dumpPostilionLines[3])[1])
	mtiStr := dumpPostilionMTIRegexMatcher.FindStringSubmatch(dumpPostilionLines[3])[1]
	parsedMessage.Mti.Version = uint32(mtiStr[0] - '0')
	parsedMessage.Mti.Class = uint32(mtiStr[1] - '0')
	parsedMessage.Mti.Function = uint32(mtiStr[2] - '0')
	parsedMessage.Mti.Origin = uint32(mtiStr[3] - '0')
	//parsedMessage.Fields = make(map[string]types.Field)

	for _, line := range dumpPostilionLines[6:] {
		if len(line) == 0 || len(dumpPostilionFieldRegexMatcher.FindStringSubmatch(line)) != 3 {
			continue
		}
		fieldNumber := dumpPostilionFieldRegexMatcher.FindStringSubmatch(line)[1]
		fieldValue := dumpPostilionFieldRegexMatcher.FindStringSubmatch(line)[2]
		parsedMessage.Fields[fieldNumber] = &protocolBuffer.Field{
			Value: strings.Trim(strings.ReplaceAll(fieldValue, "]", ""), " "),
		}
	}

	p.ParsedDumpPostilions = append(p.ParsedDumpPostilions, parsedMessage)
}

func (p *parser) parseDumpXml(dumpXml *string, fileName string, lineNumber int) {
	parsedMessage := &protocolBuffer.Message{
		LogFileName: fileName,
		LineNumber:  strconv.Itoa(lineNumber),
	}

	dumpXmlLines := strings.Split(*dumpXml, "\n")
	parsedMessage.Bitmap = strings.ReplaceAll(dumpXmlBitMapRegexMatcher.FindStringSubmatch(dumpXmlLines[0])[1], " ", "")
	//parsedMessage.SetMTI(dumpXmlMTIRegexMatcher.FindStringSubmatch(dumpXmlLines[2])[1])
	parsedMessage.Mti = &protocolBuffer.MTI{}
	mtiStr := dumpXmlMTIRegexMatcher.FindStringSubmatch(dumpXmlLines[2])[1]
	parsedMessage.Mti.Version = uint32(mtiStr[0] - '0')
	parsedMessage.Mti.Class = uint32(mtiStr[1] - '0')
	parsedMessage.Mti.Function = uint32(mtiStr[2] - '0')
	parsedMessage.Mti.Origin = uint32(mtiStr[3] - '0')
	//parsedMessage.Fields = make(map[string]types.Field)
	parsedMessage.Fields = make(map[string]*protocolBuffer.Field)
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
		parsedMessage.Fields[fieldNumber] = &protocolBuffer.Field{
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
			parsedMessage.Fields[currentFieldNumber] = &protocolBuffer.Field{
				Value: dumpXmlMultiLineFieldHeaderRegexMatcher.FindStringSubmatch(line)[2],
			}
			continue
		}
		if len(dumpXmlMultiLineFieldClosingTagRegexMatcher.FindStringSubmatch(line)) != 0 {
			parsedMessage.Fields[currentFieldNumber] = &protocolBuffer.Field{
				Value: parsedMessage.Fields[currentFieldNumber].Value + dumpXmlMultiLineFieldClosingTagRegexMatcher.FindStringSubmatch(line)[1],
			}
			continue
		}
		if len(dumpXmlMultiLineFieldRegexMatcher.FindStringSubmatch(line)) != 0 {
			parsedMessage.Fields[currentFieldNumber] = &protocolBuffer.Field{
				Value: parsedMessage.Fields[currentFieldNumber].Value + dumpXmlMultiLineFieldRegexMatcher.FindStringSubmatch(line)[1],
			}
			continue
		}
	}

	p.ParsedDumpXmls = append(p.ParsedDumpXmls, parsedMessage)
}

func (p *parser) parseDumpIso(dumpIso *string, fileName string, lineNumber int) {
	parsedMessage := &protocolBuffer.Message{
		LogFileName: fileName,
		LineNumber:  strconv.Itoa(lineNumber),
	}
	dumpIsoLines := strings.Split(*dumpIso, "\n")
	parsedMessage.Mti = &protocolBuffer.MTI{}
	parsedMessage.Fields = make(map[string]*protocolBuffer.Field)
	//parsedMessage.Fields = make(map[string]types.Field)
	for _, line := range dumpIsoLines {
		if len(line) == 0 {
			continue
		}
		if len(dumpIsoBitMapRegexMatcher.FindStringSubmatch(line)) != 0 {
			parsedMessage.Bitmap = strings.ReplaceAll(dumpIsoBitMapRegexMatcher.FindStringSubmatch(line)[1], " ", "")
			continue
		}
		if len(dumpIsoMTIRegexMatcher.FindStringSubmatch(line)) != 0 {
			//parsedMessage.SetMTI(dumpIsoMTIRegexMatcher.FindStringSubmatch(line)[1])
			mtiStr := dumpIsoMTIRegexMatcher.FindStringSubmatch(line)[1]
			parsedMessage.Mti.Version = uint32(mtiStr[0] - '0')
			parsedMessage.Mti.Class = uint32(mtiStr[1] - '0')
			parsedMessage.Mti.Function = uint32(mtiStr[2] - '0')
			parsedMessage.Mti.Origin = uint32(mtiStr[3] - '0')
			continue
		}
		if len(dumpIsoFieldRegexMatcher.FindStringSubmatch(line)) != 3 {
			continue
		}
		fieldNumber := dumpIsoFieldRegexMatcher.FindStringSubmatch(line)[1]
		fieldValue := dumpIsoFieldRegexMatcher.FindStringSubmatch(line)[2]
		parsedMessage.Fields[fieldNumber] = &protocolBuffer.Field{
			Value: strings.Trim(strings.ReplaceAll(fieldValue, "]", ""), " "),
		}
	}
	p.ParsedDumpIsos = append(p.ParsedDumpIsos, parsedMessage)
}

func (p *parser) parseDumpTlvBuffer(dumpTlvBuffer *string, fileName string, lineNumber int) {
	// parsedMessage := &types.Message{
	// 	 LogFileName: fileName,
	// }

	// p.ParsedDumpTlvBuffers = append(p.ParsedDumpTlvBuffers, parsedMessage)
}
