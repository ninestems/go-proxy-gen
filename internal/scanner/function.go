package scanner

import (
	"go/ast"

	"github.com/ninestems/go-proxy-gen/entity"
)

// function extracts function from ast tree.
func function(in *ast.Field) *entity.Function {
	funcType, ok := in.Type.(*ast.FuncType)
	if !ok {
		return nil
	}

	var (
		input  []*entity.Parameter
		output []*entity.Parameter
	)

	if funcType.Params != nil && len(funcType.Params.List) > 0 {
		input = parameters(funcType.Params.List, entity.ParameterTypeInput)
	}

	if funcType.Results != nil && len(funcType.Results.List) > 0 {
		output = parameters(funcType.Results.List, entity.ParameterTypeOutput)
	}

	return entity.NewFunction(in.Names[0].Name, input, output, tags(in.Doc))
}

// functions extracts functions from ast tree.
func functions(in []*ast.Field) []*entity.Function {
	var out = make([]*entity.Function, 0, len(in))
	for _, field := range in {
		if f := function(field); f != nil {
			out = append(out, f)
		}
	}
	return out
}
