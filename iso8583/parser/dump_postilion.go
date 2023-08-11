package parser

import (
	"bufio"
	"os"
)

func ExtractDumpPostilions(f *os.File) []string {
	f.Seek(0, 0)

	scanner := bufio.NewScanner(f)
	scanner.Split(split_by_regex(data_dump_postilion_regex))

	// header_matcher := regexp.MustCompile(header_dump_postilion_regex)

	var dump_postilions []string

	for scanner.Scan() {
		dump_postilions = append(dump_postilions, scanner.Text())
	}

	return dump_postilions
}
