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
	fld37Regexp := regexp.MustCompile(fld37DumpPostilionRegex)

	var fld37 []string

	for scanner.Scan() {
		var fld37sFound = fld37Regexp.FindStringSubmatch(scanner.Text())
		if len(fld37sFound) != 0 && !isElementExist(fld37, fld37sFound[1]) {
			fld37 = append(fld37, fld37sFound[1])
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
