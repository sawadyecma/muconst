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

	structs := parseStructNames(f, fset)
	fmt.Println(structs)
}

func parseStructNames(f *ast.File, fset *token.FileSet) map[*ast.Ident]*ast.TypeSpec {
	structs := map[*ast.Ident]*ast.TypeSpec{}

	ast.Inspect(f, func(n ast.Node) bool {
		switch a := n.(type) {
		case *ast.TypeSpec:
			structs[a.Name] = a
		}
		return true
	})

	return structs
}
