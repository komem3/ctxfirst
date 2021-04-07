// Package ctxfirst is analyzer that points out function arguments
// does not come context.Context at the beginning of the arguments.
package ctxfirst

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Doc = `check for context of argument.
this analayzer checks function args whether context.Context is first argument.
`

var Analyzer = &analysis.Analyzer{
	Name:     "ctxfirst",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}
	inspector.Preorder(nodeFilter, func(n ast.Node) {
		var ftype *ast.FuncType
		switch n := n.(type) {
		case *ast.FuncDecl:
			ftype = n.Type
		case *ast.FuncLit:
			ftype = n.Type
		}

		for i, t := range ftype.Params.List {
			if 0 < i &&
				pass.TypesInfo.Types[t.Type].Type.String() == "context.Context" {
				pass.Reportf(n.Pos(), "arguments contain context.Context but is not at the beginning")
			}
		}
	})
	return nil, nil
}
