package scanner

import (
	"github.com/ninestems/go-proxy-gen/entity"
)

type functionSpecification struct {
	name   string                    // name of function
	input  []*parameterSpecification // list of input parameters
	output []*parameterSpecification // list of output parameters
	tags   *tagSpecification         // collection of tags
}

func (s *functionSpecification) Build() *entity.Function {
	var (
		input  = make([]*entity.Parameter, 0, len(s.input))
		output = make([]*entity.Parameter, 0, len(s.output))
	)

	for _, param := range s.input {
		input = append(input, param.BuildInput())
	}

	for _, param := range s.output {
		output = append(output, param.BuildOutput())
	}

	tags := s.tags.Build()

	return entity.NewFunction(s.name, input, output, tags)
}
