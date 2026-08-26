package main

import (
	gochecksumtype "github.com/skalt/go-check-sumtype"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(gochecksumtype.Analyzer) }
