package muconst

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func Exec(filepath string) {
	// ファイルごとのトークンの位置を記録するFileSetを作成する
	fset := token.NewFileSet()

	// ファイル単位で構文解析を行う
	f, err := parser.ParseFile(fset, filepath, nil, 0)
	if err != nil {
		log.Fatal("Error:", err)
	}

	ret := checkHasConstructor(f, fset)
	fmt.Println(ret)
}

func checkHasConstructor(
	f *ast.File,
	fset *token.FileSet,
) map[string]bool {
	structs := parseStructNames(f, fset)
	constructors := hasConstructor(f, fset, structs)
	return constructors
}

func parseStructNames(f *ast.File, fset *token.FileSet) map[string]*ast.TypeSpec {
	structs := map[string]*ast.TypeSpec{}

	ast.Inspect(f, func(n ast.Node) bool {
		switch a := n.(type) {
		case *ast.TypeSpec:
			structs[a.Name.String()] = a
		}
		return true
	})

	return structs
}

func hasConstructor(
	f *ast.File,
	fset *token.FileSet,
	structs map[string]*ast.TypeSpec,
) map[string]bool {

	constructors := map[string]bool{}
	for k := range structs {
		constructors[k] = false
	}

	ast.Inspect(f, func(n ast.Node) bool {
		switch a := n.(type) {
		case *ast.FuncDecl:
			funcName := a.Name.String()
			if funcName[0:3] == "New" {
				model := funcName[3:]
				_, ok := structs[model]
				if ok {
					constructors[model] = true
				}

				if a.Type.Results == nil {
					return true
				}
				rets := a.Type.Results.List
				for i := range rets {
					v, ok := rets[i].Type.(*ast.Ident)
					if !ok {
						return true
					}
					_, ok = structs[v.Name]
					if ok {
						constructors[model] = true
					}
				}
			}
		}
		return true
	})

	return constructors
}
