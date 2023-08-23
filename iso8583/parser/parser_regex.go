package parser

import "regexp"

var dumpPostilionBitMapRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - BIT MAP\s* (.*) \s*\.`)
var dumpPostilionMTIRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - M.T.I\s*:\s* \[?(\d*)]? \s*\.`)
var dumpPostilionFieldRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - FLD \((\d*\.?\d*)\)\s*\(\d*\)\s*\[(.*) \s*\.`)
