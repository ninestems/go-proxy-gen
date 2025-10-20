package scanner

import (
	"go/ast"
)

func extractPackage(in *ast.File, relative string, names ...string) *packageSpecification {
	imptrs := extractImports(in.Imports)
	specifications := extractTypeSpecification(in.Decls)
	ifasec := extractInterfaceTypeSpecifications(specifications, names...)

	return &packageSpecification{
		name:     in.Name.Name,
		relative: relative,
		imports:  imptrs,
		ifasec:   ifasec,
	}
}
