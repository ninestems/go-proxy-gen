package scanner

import (
	"go/ast"
)

func extractFunctionType(in *ast.Field) *functionSpecification {
	processed, ok := in.Type.(*ast.FuncType)
	if !ok {
		return nil
	}

	tags := &tagSpecification{}
	if in.Doc != nil {
		tags = extractFunctionTags(in.Doc)
	}

	return &functionSpecification{
		name:   in.Names[0].Name,
		input:  extractFunctionParameters(processed.Params.List),
		output: extractFunctionParameters(processed.Results.List),
		tags:   tags,
	}
}

func extractFunctionTypes(in []*ast.Field) []*functionSpecification {
	var out = make([]*functionSpecification, 0, len(in))
	for _, spec := range in {
		processed := extractFunctionType(spec)
		if processed != nil {
			out = append(out, processed)
		}
	}

	return out
}
