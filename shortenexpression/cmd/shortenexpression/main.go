package main

import (
	"github.com/HarukiIdo/shortenexpression"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(shortenexpression.Analyzer) }
