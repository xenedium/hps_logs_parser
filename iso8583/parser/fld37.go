package parser

import (
	"bufio"
	"os"
	"regexp"
)

func ExtractFLD37(f *os.File) []string {
	f.Seek(0, 0)

	scanner := bufio.NewScanner(f)
	fld37_regexp := regexp.MustCompile(`.*FLD \(037\).*\[(.*?)\]`)

	var fld37 []string

	for scanner.Scan() {
		var fld37s_found = fld37_regexp.FindStringSubmatch(scanner.Text())
		if len(fld37s_found) != 0 && !isElementExist(fld37, fld37s_found[1]) {
			fld37 = append(fld37, fld37s_found[1])
		}
	}

	return fld37
}
func isElementExist(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
