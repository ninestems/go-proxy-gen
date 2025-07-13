package scanner

import (
	"go/ast"

	"github.com/ninestems/go-proxy-gen/entity"
)

// parameter extracts parameter from input/output list of function.
func parameter(in *ast.Field, typ entity.ParameterType) *entity.Parameter {
	var (
		out  *entity.Parameter
		name = ""
	)
	if in.Names != nil {
		name = in.Names[0].Name
	}

	switch typ {
	case entity.ParameterTypeInput:
		out = entity.NewInputParameter(name, exprToParameter(in.Type))
	case entity.ParameterTypeOutput:
		out = entity.NewOutputParameter(name, exprToParameter(in.Type))
	}

	return out
}

// parameter extracts all parameter from input/output list of function.
func parameters(in []*ast.Field, typ entity.ParameterType) []*entity.Parameter {
	var out = make([]*entity.Parameter, 0, len(in))
	for _, field := range in {
		out = append(out, parameter(field, typ))
	}

	return out
}
