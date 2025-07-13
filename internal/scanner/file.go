package scanner

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

// file parses source to *ast.File.
func file(path string) (*ast.File, error) {
	f, err := parser.ParseFile(token.NewFileSet(), path, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("parsing file in path %s as ast.File with error: %w", path, err)
	}

	return f, nil
}
