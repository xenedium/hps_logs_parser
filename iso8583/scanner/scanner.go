package scanner

import (
	"bufio"
	"io"
	"os"
	"regexp"
)

type scanner struct {
	File           *os.File
	Fld37          []string
	DumpPostilions []string
	DumpXmls       []string
}
type matcherHandlerArray struct {
	Matcher *regexp.Regexp
	Handler func(*bufio.Scanner) string
	Array   *[]string
}

func (s *scanner) Scan() {
	_, err := s.File.Seek(0, io.SeekStart)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(s.File)

	// TODO: add more matchers and handlers
	// ALL THE HANDLERS MUST RETURN A STRING
	mhaArray := []matcherHandlerArray{
		{
			Matcher: regexp.MustCompile(startDumpPostilionRegex),
			Handler: readDumpPostilion,
			Array:   &s.DumpPostilions,
		},
		{
			Matcher: regexp.MustCompile(startXmlDumpRegex),
			Handler: readDumpXml,
			Array:   &s.DumpXmls,
		},
	}

	for scanner.Scan() {
		for _, mha := range mhaArray {
			if mha.Matcher.MatchString(scanner.Text()) {
				*mha.Array = append(*mha.Array, mha.Handler(scanner))
			}
		}
	}
}

type Scanner = scanner
