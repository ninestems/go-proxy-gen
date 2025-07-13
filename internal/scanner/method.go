package scanner

import (
	"slices"

	"github.com/ninestems/go-proxy-gen/entity"
	"github.com/ninestems/go-proxy-gen/pkg/log"
)

// Scan parses the specified path and returns a list of interfaces
// filtered by optional names (if provided).
func (s *Scanner) Scan(in string, names ...string) (*entity.Package, error) {
	log.Debugf("scan file in %s start", in)

	node, err := file(in)
	if err != nil {
		return nil, err
	}

	log.Debugf("scan file in %s start", in)

	out := entity.NewPackage(
		node.Name.Name,
		"",
		imports(node.Imports),
		slices.DeleteFunc(declarations(node.Decls), func(in *entity.Interface) bool {
			return !slices.Contains(names, in.Name()) && len(names) > 0
		}),
	)

	return out, nil
}
