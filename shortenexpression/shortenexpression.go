package shortenexpression

import (
	"fmt"
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
	stmtCount := 0
	inspect.Preorder(nil, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.AssignStmt:
			ast.Inspect(n, func(n ast.Node) bool {
				switch n := n.(type) {
				case *ast.BinaryExpr:	
					stmtCount++
					fmt.Println(stmtCount)
					if stmtCount > 5 {
						pass.Reportf(n.Pos(), "expression is too long")
						return false
					}
					
				}
				return true
			})
		}
	})

	return nil, nil
}
