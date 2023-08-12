package scanner

import (
	"bufio"
	"regexp"
)

var splitByRegex = func(regex string) bufio.SplitFunc {
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		reg := regexp.MustCompile(regex)
		loc := reg.FindIndex(data)
		if loc == nil {
			return 0, nil, nil
		}
		return loc[1], data[loc[0]:loc[1]], nil
	}
}
