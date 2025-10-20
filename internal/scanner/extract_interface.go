package scanner

import (
	"go/ast"
	"go/token"
	"slices"
)

func extractTypeSpecification(in []ast.Decl) []*ast.TypeSpec {
	var genDecls = make([]*ast.GenDecl, 0, len(in))
	for _, spec := range in {
		genDecl, ok := spec.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		genDecls = append(genDecls, genDecl)
	}

	var specs = make([]ast.Spec, 0, len(genDecls))
	for _, genSpec := range genDecls {
		specs = append(specs, genSpec.Specs...)
	}

	var typeSpecs = make([]*ast.TypeSpec, 0, len(specs))
	for _, spec := range specs {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}
		typeSpecs = append(typeSpecs, typeSpec)
	}

	return typeSpecs
}

func extractInterfaceTypeSpecifications(in []*ast.TypeSpec, names ...string) []*interfaceSpecification {
	var interfaceTypes = make([]*interfaceSpecification, 0, len(in))

	for _, spec := range in {
		processed, ok := spec.Type.(*ast.InterfaceType)
		if !ok {
			continue
		}
		isFiltered := !slices.Contains(names, spec.Name.Name) && len(names) > 0
		isEmptyMethod := processed.Methods == nil || processed.Methods != nil && len(processed.Methods.List) == 0
		if isFiltered || isEmptyMethod {
			continue
		}

		interfaceTypes = append(
			interfaceTypes,
			&interfaceSpecification{
				name:      spec.Name.Name,
				functions: extractFunctionTypes(processed.Methods.List),
			},
		)
	}

	return interfaceTypes
}
