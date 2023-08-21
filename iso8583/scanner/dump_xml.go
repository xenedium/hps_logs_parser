package scanner

import (
	"bufio"
	"regexp"
	"strings"
)

func readDumpXml(scanner *bufio.Scanner) string {
	endDumpXmlMatcher := regexp.MustCompile(endXmlDumpRegex)
	dumpXmlStr := strings.Builder{}
	for scanner.Scan() {
		if endDumpXmlMatcher.MatchString(scanner.Text()) {
			break
		}
		dumpXmlStr.WriteString(scanner.Text() + "\n")
	}
	return dumpXmlStr.String()
}
