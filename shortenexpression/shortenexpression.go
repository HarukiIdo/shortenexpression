package shortenexpression

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "shortenexpression is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "shortenexpression",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (any, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.AssignStmt:
			var count int
			for _, r := range n.Rhs {
				count += countNode(r)
			}
			if count > 5 {
				pass.Reportf(n.Pos(), "expression is too long (%d)", count)
			}
		}
	})

	return nil, nil
}


func countNode(n ast.Node) int {
	var count int

	ast.Inspect(n, func(n ast.Node) bool {
		count++
		return true
	})

	return count
}

