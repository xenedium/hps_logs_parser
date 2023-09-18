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
	DumpPostilions map[int]string
	DumpXmls       map[int]string
	DumpIsos       map[int]string
	DumpTlvBuffers map[int]string
	DumpBuffers    map[int]string
}
type matcherHandlerArray struct {
	Matcher *regexp.Regexp
	Handler func(*bufio.Scanner, *int) string
	Array   *map[int]string
}

func (s *scanner) Scan() {
	_, err := s.File.Seek(0, io.SeekStart)
	if err != nil {
		return
	}

	scanner := bufio.NewScanner(s.File)

	// TODO: add more matchers and handlers
	// ALL THE HANDLERS MUST RETURN A string AND RECEIVE A *bufio.Scanner
	mhaArray := []matcherHandlerArray{
		{
			Matcher: regexp.MustCompile(startDumpPostilionRegex),
			Handler: getGenericHandler(endDumpPostilionRegex),
			Array:   &s.DumpPostilions,
		},
		{
			Matcher: regexp.MustCompile(startXmlDumpRegex),
			Handler: getGenericHandler(endXmlDumpRegex),
			Array:   &s.DumpXmls,
		},
		{
			Matcher: regexp.MustCompile(startDumpIso),
			Handler: getGenericHandler(endDumpIso),
			Array:   &s.DumpIsos,
		},
		{
			Matcher: regexp.MustCompile(startDumpTlvBuffer),
			Handler: getGenericHandler(endDumpTlvBuffer),
			Array:   &s.DumpTlvBuffers,
		},
		{
			Matcher: regexp.MustCompile(startDumpBuffer),
			Handler: getGenericHandler(endDumpBuffer),
			Array:   &s.DumpBuffers,
		},
	}

	s.DumpIsos = make(map[int]string)
	s.DumpPostilions = make(map[int]string)
	s.DumpTlvBuffers = make(map[int]string)
	s.DumpXmls = make(map[int]string)
	s.DumpBuffers = make(map[int]string)

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		for _, mha := range mhaArray {
			if mha.Matcher.MatchString(scanner.Text()) {
				(*mha.Array)[lineNumber] = mha.Handler(scanner, &lineNumber)
			}
		}
	}
}

type Scanner = scanner
