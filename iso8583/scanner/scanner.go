package scanner

import (
	"os"
)

type scanner struct {
	File           *os.File
	fld37          []string
	dumpPostilions []string
}

type Scanner = scanner
