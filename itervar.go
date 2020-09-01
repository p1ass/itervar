package itervar

import (
	"fmt"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "itervar is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "itervar",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.ForStmt)(nil),
	}

	inspector.Preorder(nodeFilter, func(n ast.Node) {
		ast.Print(pass.Fset, n)

		forStmt := n.(*ast.ForStmt)
		iterVar := extractIteratorVariable(forStmt)
		fmt.Println(iterVar)
	})

	return nil, nil
}

func extractIteratorVariable(forStmt *ast.ForStmt) string {
	iterVar := ""
	switch init := forStmt.Init.(type) {
	case *ast.AssignStmt:
		if len(init.Lhs) == 0 {
			break
		}
		switch lhs := init.Lhs[0].(type) {
		case *ast.Ident:
			if lhs.Obj.Kind == ast.Var {
				iterVar = lhs.Name
			}
		}
	}
	return iterVar
}
