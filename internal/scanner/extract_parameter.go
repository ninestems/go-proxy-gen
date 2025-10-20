package scanner

import (
	"go/ast"
	"strings"
)

func extractFunctionParameter(in *ast.Field) *parameterSpecification {
	names := make([]string, 0, len(in.Names))
	source := exprToParameter(in.Type)
	for _, spec := range in.Names {
		names = append(names, spec.Name)
	}

	return &parameterSpecification{
		names:   names,
		source:  source,
		pointer: strings.Contains(source, "*"),
	}
}

func extractFunctionParameters(in []*ast.Field) []*parameterSpecification {
	var out = make([]*parameterSpecification, 0, len(in))
	for _, spec := range in {
		out = append(out, extractFunctionParameter(spec))
	}

	return out
}
