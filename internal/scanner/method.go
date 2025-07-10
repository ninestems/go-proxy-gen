package scanner

import (
	"slices"

	"go-proxy-gen/entity"
)

// Scan parses the specified path and returns a list of interfaces
// filtered by optional names (if provided).
func (s *Scanner) Scan(in string, names ...string) (*entity.Package, error) {
	node, err := file(in)
	if err != nil {
		return nil, err
	}

	out := entity.NewPackage(
		node.Name.Name,
		"",
		0,
		0,
		0,
		imports(node.Imports),
		slices.DeleteFunc(declarations(node.Decls), func(in *entity.Interface) bool {
			return !slices.Contains(names, in.Name()) && len(names) > 0
		}),
	)

	return out, nil
}
