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
		(*ast.RangeStmt)(nil),
	}

	inspector.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.ForStmt:
			checkForStmt(pass, n)
		case *ast.RangeStmt:
			checkRangeStmt(pass, n)
		}
	})

	return nil, nil
}

func checkForStmt(pass *analysis.Pass, forStmt *ast.ForStmt) {
	iterVar := extractIteratorVariableFromForStmt(forStmt)
	if iterVar == "" {
		return
	}

	if forStmt.Body == nil {
		return
	}

	for _, stmt := range forStmt.Body.List {
		findUsingIterVarRef(pass, stmt, iterVar)
	}
}

func checkRangeStmt(pass *analysis.Pass, forStmt *ast.RangeStmt) {
	iterVar := extractIteratorVariableFromRangeStmt(forStmt)
	if iterVar == "" {
		return
	}

	if forStmt.Body == nil {
		return
	}

	for _, stmt := range forStmt.Body.List {
		findUsingIterVarRef(pass, stmt, iterVar)
	}
}

func extractIteratorVariableFromForStmt(forStmt *ast.ForStmt) string {
	switch init := forStmt.Init.(type) {
	case *ast.AssignStmt:
		return extractIteratorVariableFromAssignStmt(init)
	}
	return ""
}

func extractIteratorVariableFromRangeStmt(rangeStmt *ast.RangeStmt) string {
	ident, ok := rangeStmt.Key.(*ast.Ident)
	if !ok {
		return ""
	}
	assignStmt, ok := ident.Obj.Decl.(*ast.AssignStmt)
	if !ok {
		return ""
	}

	return extractIteratorVariableFromAssignStmt(assignStmt)
}

func extractIteratorVariableFromAssignStmt(stmt *ast.AssignStmt) string {
	if len(stmt.Lhs) == 0 {
		return ""
	}

	iterVar := stmt.Lhs[0]
	if len(stmt.Lhs) > 1 {
		iterVar = stmt.Lhs[1]
	}

	switch iterVar := iterVar.(type) {
	case *ast.Ident:
		// TODO これがちゃんとイテレータになってるか確認する (インクリメントされてるとか）
		if iterVar.Obj.Kind == ast.Var {

			return iterVar.Name
		}
	}
	return ""
}

func findUsingIterVarRef(pass *analysis.Pass, stmt ast.Stmt, iterVar string) {
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
