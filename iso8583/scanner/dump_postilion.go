package scanner

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

// ExtractDumpPostilions should extract all dump_postilion from log file
func ExtractDumpPostilions(f *os.File) []string {
	_, err := f.Seek(0, 0)
	if err != nil {
		return nil
	}

	scanner := bufio.NewScanner(f)

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
