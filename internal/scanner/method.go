package scanner

import (
	"slices"

	"go-proxy-gen/entity"
)

// Scan parses the specified path and returns a list of interfaces
// filtered by optional names (if provided).
func (s *Scanner) Scan(in string, names ...string) (*entity.Package, error) {
	var out entity.Package
	node, err := file(in)
	if err != nil {
		return nil, err
	}

	out.SetName(node.Name.Name)
	out.SetInterfaces(slices.DeleteFunc(declarations(node.Decls), func(in *entity.Interface) bool {
		return !slices.Contains(names, in.Name()) && len(names) > 0
	}))
	out.SetImports(imports(node.Imports))

	return &out, nil
}
