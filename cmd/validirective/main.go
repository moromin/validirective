package main

import (
	"validirective"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(validirective.Analyzer) }
