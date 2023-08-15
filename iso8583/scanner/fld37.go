package scanner

import (
	"bufio"
	"os"
	"regexp"
)

// ExtractFLD37 should extract all fld37 from all different types of dumps not only from dump_postilion
func ExtractFLD37(f *os.File) []string {
	_, err := f.Seek(0, 0)
	if err != nil {
		return nil
	}
	scanner := bufio.NewScanner(f)
	fld37PostRegexp := regexp.MustCompile(fld37DumpPostilionRegex)
	fld37XmlRegexp := regexp.MustCompile(fld37XmlDumpRegex)
	fld37BuffRegexp := regexp.MustCompile(fld37DumpBufferRegex)

	var fld37 []string

	for scanner.Scan() {
		var fld37sFound string
		var matchArray = fld37PostRegexp.FindStringSubmatch(scanner.Text())
		matchArray = append(matchArray, fld37XmlRegexp.FindStringSubmatch(scanner.Text())...)
		matchArray = append(matchArray, fld37BuffRegexp.FindStringSubmatch(scanner.Text())...)
		if len(matchArray) > 0 {
			fld37sFound = matchArray[1]
			goto append
		}

		continue

	append:
		if fld37sFound != "" && !isElementExist(fld37, fld37sFound) {
			fld37 = append(fld37, fld37sFound)
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
