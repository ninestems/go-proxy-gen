package scanner

import (
	"github.com/ninestems/go-proxy-gen/pkg/entity"
)

type parameterSpecification struct {
	names   []string // list names of parameter
	source  string   // source path of parameters, e.x. base type or custom struct from other package
	pointer bool     // sets in true if parameter is a pointer on type
}

func (s *parameterSpecification) BuildInput() *entity.Parameter {
	return entity.NewInputParameter(s.names, s.source, s.pointer)
}

func (s *parameterSpecification) BuildOutput() *entity.Parameter {
	return entity.NewOutputParameter(s.names, s.source, s.pointer)
}
