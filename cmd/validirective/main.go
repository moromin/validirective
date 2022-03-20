package main

import (
	"github.com/moromin/validirective"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(validirective.Analyzer) }
