package scanner

import (
	"github.com/ninestems/go-proxy-gen/entity"
)

type interfaceSpecification struct {
	name      string
	functions []*functionSpecification
}

func (s *interfaceSpecification) Build(pname, relative string, imprts []*entity.Import) *entity.Interface {
	var fns = make([]*entity.Function, 0, len(s.functions))
	for _, fn := range s.functions {
		fns = append(fns, fn.Build())
	}

	return entity.NewInterfaceV2(s.name, pname, relative, imprts, fns)
}
