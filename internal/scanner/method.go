package scanner

import (
	"github.com/ninestems/go-proxy-gen/entity"
	"github.com/ninestems/go-proxy-gen/pkg/log"
)

// Scan scans the provided path and returns interface descriptions
// with parsed metadata such as proxy tags.
func (s *Scanner) Scan(in string) ([]*entity.Interface, error) {
	log.Debugf("scan file in %s: start", in)

	node, err := file(in)
	if err != nil {
		return nil, err
	}

	log.Debugf("scan file in %s: success", in)

	return extractPackage(node, s.relative, s.names...).Interfaces()
}
