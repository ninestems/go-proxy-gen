package scanner

import (
	"github.com/ninestems/go-proxy-gen/entity"
)

type importSpecification struct {
	alias string
	path  string
}

// Build return *entity.Import instance.
func (s *importSpecification) Build() *entity.Import {
	return entity.NewImport(s.alias, s.path)
}
