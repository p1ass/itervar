// This file can build as a plugin for golangci-lint by below command.
//    go build -buildmode=plugin -o path_to_plugin_dir github.com/p1ass/itervar/plugin/itervar
// See: https://golangci-lint.run/contributing/new-linters/#how-to-add-a-private-linter-to-golangci-lint

package main

import (
	"strings"

	"github.com/p1ass/itervar"
	"golang.org/x/tools/go/analysis"
)

// flags for Analyzer.Flag.
// If you would like to specify flags for your plugin, you can put them via 'ldflags' as below.
//     $ go build -buildmode=plugin -ldflags "-X 'main.flags=-opt val'" github.com/p1ass/itervar/plugin/itervar
var flags string

// AnalyzerPlugin provides analyzers as a plugin.
// It follows golangci-lint style plugin.
var AnalyzerPlugin analyzerPlugin //nolint

type analyzerPlugin struct{}

func (analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	if flags != "" {
		flagset := itervar.Analyzer.Flags
		if err := flagset.Parse(strings.Split(flags, " ")); err != nil {
			panic("cannot parse flags of itervar: " + err.Error())
		}
	}
	return []*analysis.Analyzer{
		itervar.Analyzer,
	}
}
