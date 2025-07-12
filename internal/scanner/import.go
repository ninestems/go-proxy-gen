package scanner

import (
	"go/ast"

	"go-proxy-gen/entity"
)

// imports extracts imports from ast tree.
func imports(in []*ast.ImportSpec) []*entity.Import {
	var out = make([]*entity.Import, 0, len(in))
	for _, imp := range in {
		alias := ""
		if imp.Name != nil {
			alias = imp.Name.Name
		}

		out = append(out, entity.NewImport(alias, imp.Path.Value))
	}

	return out
}
