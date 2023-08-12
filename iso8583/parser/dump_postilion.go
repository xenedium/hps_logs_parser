package parser

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ExtractDumpPostilions(f *os.File) []string {
	_, err := f.Seek(0, 0)
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(f)

	headerMatcher := regexp.MustCompile(start_dump_postilion_regex)

	var dumpPostilions []string

	for scanner.Scan() {
		if headerMatcher.MatchString(scanner.Text()) {
			dumpPostilions = append(dumpPostilions, readDumpPostilion(scanner))
		}
	}

	return dumpPostilions
}

func readDumpPostilion(scanner *bufio.Scanner) string {
	dumpPostilionHeader := regexp.MustCompile(data_dump_postilion_regex)
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
