package itervar

import (
	"go/ast"
	"go/token"

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
		switch n := n.(type) {
		case *ast.ForStmt:
			checkForStmt(pass, n)
		}
	})

	return nil, nil
}

func checkForStmt(pass *analysis.Pass, forStmt *ast.ForStmt) {
	iterVar := extractIteratorVariable(forStmt)
	if iterVar == "" {
		return
	}

	if forStmt.Body == nil {
		return
	}

	for _, stmt := range forStmt.Body.List {
		traverseStmt(pass, stmt, iterVar)
	}
}

func extractIteratorVariable(forStmt *ast.ForStmt) string {
	iterVar := ""
	switch init := forStmt.Init.(type) {
	case *ast.AssignStmt:
		if len(init.Lhs) == 0 {
			break
		}
		// TODO １つ以上の場合も考える
		switch lhs := init.Lhs[0].(type) {
		case *ast.Ident:
			// TODO これがちゃんとイテレータになってるか確認する (インクリメントされてるとか）
			if lhs.Obj.Kind == ast.Var {
				iterVar = lhs.Name
			}
		}
	}
	return iterVar
}

func traverseStmt(pass *analysis.Pass, stmt ast.Stmt, iterVar string) {
	ast.Inspect(stmt, func(n ast.Node) bool {
		if n == nil {
			return true
		}
		switch n := n.(type) {
		case *ast.UnaryExpr:
			x, ok := n.X.(*ast.Ident)
			if !ok {
				return true
			}
			if n.Op == token.AND && x.Obj.Kind == ast.Var && x.Obj.Name == iterVar {
				pass.Reportf(n.Pos(), "using reference to loop iterator variable")
			}
		}

		return true
	})
}
