package main

import (
	"github.com/p1ass/itervar"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(itervar.Analyzer) }
