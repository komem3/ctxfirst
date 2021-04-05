package ctxfirst_test

import (
	"testing"

	"github.com/komem3/ctxfirst"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, ctxfirst.Analyzer, "a")
}
