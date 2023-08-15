package scanner

import (
	"bufio"
	"regexp"
	"strings"
)

// ExtractDumpPostilions should extract all dump_postilion from log file
func (s *Scanner) extractDumpPostilions() []string {
	_, err := s.File.Seek(0, 0)
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(s.File)

	headerMatcher := regexp.MustCompile(startDumpPostilionRegex)

	var dumpPostilions []string

	for scanner.Scan() {
		if headerMatcher.MatchString(scanner.Text()) {
			dumpPostilions = append(dumpPostilions, readDumpPostilion(scanner))
		}
	}

	return dumpPostilions
}

func readDumpPostilion(scanner *bufio.Scanner) string {
	dumpPostilionHeader := regexp.MustCompile(dataDumpPostilionRegex)
	dumpPostilionStr := strings.Builder{}
	for scanner.Scan() {
		if dumpPostilionHeader.MatchString(scanner.Text()) {
			dumpPostilionStr.WriteString(scanner.Text() + "\n")
			continue
		}
		break
	}
	return dumpPostilionStr.String()
}

func (s *Scanner) GetPostilionDumps() []string {
	if s.dumpPostilions == nil {
		s.dumpPostilions = s.extractDumpPostilions()
	}
	return s.dumpPostilions
}
