package scanner

import (
	"bufio"
	"regexp"
	"strings"
)

func getGenericHandler(regex string) func(scanner *bufio.Scanner, lineNumber *int) string {
	return func(scanner *bufio.Scanner, lineNumber *int) string {
		endMatcher := regexp.MustCompile(regex)
		dumpPostilionStr := strings.Builder{}
		for scanner.Scan() {
			*lineNumber++
			if endMatcher.MatchString(scanner.Text()) {
				break
			}
			dumpPostilionStr.WriteString(scanner.Text() + "\n")
		}
		return dumpPostilionStr.String()
	}
}
