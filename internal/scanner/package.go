package scanner

import (
	"go/ast"
	"slices"

	entity "github.com/ninestems/go-proxy-gen/entity"
)

// packag returns new entity.Package build from ast.File.
func packag(node *ast.File, names ...string) *entity.Package {
	return entity.NewPackage(
		node.Name.Name,
		"",
		imports(node.Imports),
		slices.DeleteFunc(declarations(node.Decls), func(in *entity.Interface) bool {
			return !slices.Contains(names, in.Name()) && len(names) > 0
		}),
	)
}
