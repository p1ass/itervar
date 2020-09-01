package itervar_test

import (
	"testing"

	"github.com/p1ass/itervar"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, itervar.Analyzer, "ref_to_loop_iter_var")
}
