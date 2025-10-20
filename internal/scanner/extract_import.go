package scanner

import (
	"go/ast"
)

func extractImports(in []*ast.ImportSpec) []*importSpecification {
	var out = make([]*importSpecification, 0, len(in))
	for _, imp := range in {
		alias := ""
		if imp.Name != nil {
			alias = imp.Name.Name
		}
		out = append(out, &importSpecification{
			alias: alias,
			path:  imp.Path.Value,
		})
	}
	return out
}
