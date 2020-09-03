package itervar

import (
	"go/ast"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "itervar is a static analysis tool which detects references to loop iterator variable."

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
	assignStmt, ok := forStmt.Init.(*ast.AssignStmt)
	if !ok {
		return
	}

	iterVar := searchIteratorVariableIdent(assignStmt)
	reportUsingIterVarRef(pass, forStmt.Body, iterVar)

}

func checkRangeStmt(pass *analysis.Pass, rangeStmt *ast.RangeStmt) {
	ident, ok := rangeStmt.Key.(*ast.Ident)
	if !ok {
		return
	}
	assignStmt, ok := ident.Obj.Decl.(*ast.AssignStmt)
	if !ok {
		return
	}

	iterVar := searchIteratorVariableIdent(assignStmt)
	reportUsingIterVarRef(pass, rangeStmt.Body, iterVar)
}

func searchIteratorVariableIdent(stmt *ast.AssignStmt) *ast.Ident {
	if len(stmt.Lhs) == 0 {
		return nil
	}

	iterVar := stmt.Lhs[0]
	if len(stmt.Lhs) > 1 {
		iterVar = stmt.Lhs[1]
	}

	switch iterVar := iterVar.(type) {
	case *ast.Ident:
		if iterVar.Obj.Kind == ast.Var {
			return iterVar
		}
	}
	return nil
}

func reportUsingIterVarRef(pass *analysis.Pass, stmt ast.Stmt, iterVar *ast.Ident) {
	typ := pass.TypesInfo.TypeOf(iterVar)

	ast.Inspect(stmt, func(n ast.Node) bool {
		if n == nil {
			return false
		}

		// スコープ内で変数を再定義していたら探索をスキップする
		s := pass.TypesInfo.Scopes[n]
		if s != nil {
			if s.Lookup(iterVar.Obj.Name) != nil {
				return false
			}
		}

		switch n := n.(type) {
		// &i を検出
		case *ast.UnaryExpr:
			x, ok := n.X.(*ast.Ident)
			if !ok {
				return true
			}
			if n.Op == token.AND && x.Obj == iterVar.Obj {
				pass.Reportf(n.Pos(), "using reference to loop iterator variable")
			}
		// i[:]を検出
		case *ast.SliceExpr:
			x, ok := n.X.(*ast.Ident)
			if !ok {
				return true
			}

			if x.Obj == iterVar.Obj && isArray(typ) {
				pass.Reportf(n.Pos(), "using reference to loop iterator variable")
			}
		}

		return true
	})
}

func isArray(typ types.Type) bool {
	switch typ.(type) {
	case *types.Array:
		return true
	}
	return false
}
