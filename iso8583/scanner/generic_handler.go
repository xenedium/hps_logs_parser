package scanner

import (
	"bufio"
	"regexp"
	"strings"
)

func getGenericHandler(regex string) func(scanner *bufio.Scanner) string {
	return func(scanner *bufio.Scanner) string {
		endMatcher := regexp.MustCompile(regex)
		dumpPostilionStr := strings.Builder{}
		for scanner.Scan() {
			if endMatcher.MatchString(scanner.Text()) {
				break
			}
			dumpPostilionStr.WriteString(scanner.Text() + "\n")
		}
		return dumpPostilionStr.String()
	}
}
