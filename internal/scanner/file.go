package scanner

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// file parses source to *ast.File.
func file(filename string) (*ast.File, error) {
	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return f, nil
}
