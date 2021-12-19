package main

import (
	"path"

	"github.com/sawadyecma/muconst"
)

var rootPath = "/"

func main() {
	muconst.Exec(
		path.Join(
			rootPath,
			"./testdata/src/a/a.go",
		),
	)
}
