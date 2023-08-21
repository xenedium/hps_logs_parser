package scanner

import (
	"bufio"
	"regexp"
	"strings"
)

func readDumpPostilion(scanner *bufio.Scanner) string {
	endDumpPostilionMatcher := regexp.MustCompile(endDumpPostilionRegex)
	dumpPostilionStr := strings.Builder{}
	for scanner.Scan() {
		if endDumpPostilionMatcher.MatchString(scanner.Text()) {
			break
		}
		dumpPostilionStr.WriteString(scanner.Text() + "\n")
	}
	return dumpPostilionStr.String()
}
