package shortenexpression_test

import (
	"testing"

	"github.com/HarukiIdo/shortenexpression"
	"github.com/gostaticanalysis/testutil"
	"golang.org/x/tools/go/analysis/analysistest"
)

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, shortenexpression.Analyzer, "a")
}
