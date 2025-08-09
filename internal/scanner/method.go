package scanner

import (
	"github.com/ninestems/go-proxy-gen/entity"
	"github.com/ninestems/go-proxy-gen/pkg/log"
)

// Scan parses the specified path and returns a list of interfaces
// filtered by optional names (if provided).
func (s *Scanner) Scan(in string, names ...string) (*entity.Package, error) {
	log.Debugf("scan file in %s: start", in)

	node, err := file(in)
	if err != nil {
		return nil, err
	}

	log.Debugf("scan file in %s: success", in)

	return packag(node, names...), nil
}
