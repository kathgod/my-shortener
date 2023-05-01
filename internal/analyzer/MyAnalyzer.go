package analyzer

import (
	"log"

	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// ExitInMainAnalyzer переменная типа nalysis.Analyzer, нужна для реализации анализатора вызова os.Exit в функции main() пакета main
var ExitInMainAnalyzer = &analysis.Analyzer{
	Name: "NoOsExitInMain",
	Doc:  "os.Exit() in main function",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		f := file.Name.Name
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok {
				continue
			}
			if fn.Name.Name != "main" || fn.Recv != nil || f != "main" {
				continue
			}
			ast.Inspect(fn.Body, func(node ast.Node) bool {
				call, ok := node.(*ast.CallExpr)
				if !ok {
					return true
				}
				ident, ok := call.Fun.(*ast.SelectorExpr)
				if !ok || ident.Sel.Name != "Exit" {
					return true
				}
				selector, ok := ident.X.(*ast.Ident)
				if !ok || selector.Name != "os" {
					return true
				}
				pass.Reportf(call.Pos(), "os.Exit should not be called directly in main function")
				return false
			})
		}
	}
	log.Println("PASS")
	return nil, nil
}
