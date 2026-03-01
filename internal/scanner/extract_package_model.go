package scanner

import (
	"github.com/ninestems/go-proxy-gen/pkg/entity"
)

type packageSpecification struct {
	name     string
	relative string
	imports  []*importSpecification
	ifasec   []*interfaceSpecification
}

func (s *packageSpecification) Interfaces() ([]*entity.Interface, error) {
	var (
		interfaces = make([]*entity.Interface, 0, len(s.ifasec))
		imprts     = make([]*entity.Import, 0, len(s.imports))
	)

	for _, imprt := range s.imports {
		imprts = append(imprts, imprt.Build())
	}

	for _, ifa := range s.ifasec {
		interfaces = append(interfaces, ifa.Build(s.name, s.relative, imprts))
	}

	return interfaces, nil
}
