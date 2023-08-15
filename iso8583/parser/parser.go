package parser

import (
	"os"

	"github.com/xenedium/hps_logs_parser/iso8583/scanner"
	"github.com/xenedium/hps_logs_parser/iso8583/types"
)

type parser struct {
	scanner  scanner.Scanner
	messages []types.Message
}

func (p *parser) Parse(f *os.File) []types.Message {
	p.scanner = scanner.Scanner{File: f}
	fld37s := p.scanner.GetFLD37()                  // extract all the transaction IDs from different types of dumps
	p.messages = make([]types.Message, len(fld37s)) // allocate an array depending on the number of transaction IDs found

	for i, fld37 := range fld37s {
		p.messages[i] = types.Message{
			Fields: map[int]types.Field{
				37: {
					Length: len(fld37),
					Value:  fld37,
					Raw:    []byte(fld37),
				},
			},
		}
	}

	return p.messages
}

type Parser = parser
