package validirective

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "validirective is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "validirective",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.CommentGroup)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.CommentGroup:
			for _, comment := range n.List {
				index := strings.Index(comment.Text, "go:")
				if index != -1 && index != 2 {
					pass.Reportf(n.Pos(), "this directive has become a comment...")
				}
			}
		}
	})

	return nil, nil
}
