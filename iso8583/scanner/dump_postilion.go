package scanner

import (
	"bufio"
	"regexp"
	"strings"
)

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
