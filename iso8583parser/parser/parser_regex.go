package parser

import "regexp"

var dumpPostilionBitMapRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - BIT MAP\s* (.*)\s*\.`)
var dumpPostilionMTIRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - M.T.I\s*:\s* \[?(\d*)]?\s*\.`)
var dumpPostilionFieldRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - FLD \((\d*\.?\d*)\)\s*\(\d*\)\s*\[(.*)\s*\.`)

var dumpXmlBitMapRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s* (.*) \s*\.`)
var dumpXmlMTIRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s*MESSAGE_ISO_XML_FORMAT_START\[(\d*)]\s*\.`)
var dumpXmlFieldRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s* <Field Number="(\d{3})" Value="(.*)"/>\s*\.`)
var dumpXmlClosingMessageInfoTagRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s*</MessageInfo>\s*\.`)
var dumpXmlMultiLineFieldHeaderRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s*<Field Number="(\d{3})" Value="(.*)\s*\.`)
var dumpXmlMultiLineFieldRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s* (.*)\s*\.`)
var dumpXmlMultiLineFieldClosingTagRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\|\s*(.*)"/>\s*\.`)

var dumpIsoBitMapRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - BIT MAP\s*:\s*(.*)\s*\.`)
var dumpIsoMTIRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - M.T.I\s*:\s* \[?(\d*)]?\s*\.`)
var dumpIsoFieldRegexMatcher = regexp.MustCompile(`\d+ \d+ \d+ \d+\|\w\| - FLD \((\d*)\)\s*:\s*\(\d*\)\s*:\s*\[(.*)\s*\.`)
