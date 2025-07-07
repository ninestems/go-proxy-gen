package scanner

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"

	"go-proxy-gen/entity"
)

// file parses source to *ast.File.
func file(filename string) (*ast.File, error) {
	f, err := parser.ParseFile(token.NewFileSet(), filename, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	return f, nil
}

// tag extracts custom tag from one line of comment.
func tag(in []string) *entity.Tag {
	var out entity.Tag
	switch in[0] {
	case "log":
		out.Type = entity.ProxyTypeLogger
	case "trace":
		out.Type = entity.ProxyTypeTracer
	case "retry":
		out.Type = entity.ProxyTypeRetrier
	}

	data := strings.Split(in[1], "::")
	switch len(data) {
	case 2:
		out.Name = data[0]
		out.Alias = data[1]
		out.Path = data[1]
	case 3:
		out.Name = data[0]
		out.Alias = data[1]
		out.Path = data[2]
	}

	return &out
}

// tags extracts custom tags from all lines of comments.
func tags(in *ast.CommentGroup) []*entity.Tag {
	var (
		out   = make([]*entity.Tag, 0, len(in.List))
		found = false
	)

	for _, c := range in.List {
		line := strings.TrimSpace(strings.TrimPrefix(c.Text, "//"))
		switch {
		case strings.HasPrefix(line, "goproxygen:"):
			found = true
			continue

		case !found || line == "":
			continue

		default:
			out = append(out, tag(strings.Fields(line)))

		}
	}

	return out
}

// parameter extracts parameter from input/output list of function.
func parameter(in *ast.Field) *entity.Parameter {
	var out = entity.Parameter{Name: in.Names[0].Name}
	switch exp := in.Type.(type) {
	case *ast.Ident:
		out.Source = exp.Name
	case *ast.SelectorExpr:
		out.Source = fmt.Sprintf("%s.%s", exp.X.(*ast.Ident), exp.Sel.Name)
	case *ast.StarExpr:
		selector := exp.X.(*ast.SelectorExpr)
		out.Source = fmt.Sprintf("%s.%s", selector.X.(*ast.Ident), selector.Sel.Name)
	}

	return &out
}

// parameter extracts all parameter from input/output list of function.
func parameters(in []*ast.Field) []*entity.Parameter {
	var out = make([]*entity.Parameter, 0, len(in))
	for _, field := range in {
		out = append(out, parameter(field))
	}

	return out
}

// function extracts function from ast tree.
func function(in *ast.Field) *entity.Function {
	funcType, ok := in.Type.(*ast.FuncType)
	if !ok {
		return nil
	}

	var out = entity.Function{
		Name:   in.Names[0].Name,
		Input:  parameters(funcType.Params.List),
		Output: parameters(funcType.Results.List),
		Tags:   tags(in.Doc),
	}

	return &out
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

// iface extracts interface from ast tree.
func iface(in *ast.TypeSpec) *entity.Interface {
	var out = entity.Interface{Name: in.Name.Name}
	ifa, ok := in.Type.(*ast.InterfaceType)
	if !ok {
		return nil
	}

	out.Functions = functions(ifa.Methods.List)

	return &out
}

// ifaces extracts interfaces from ast tree.
func ifaces(in []ast.Spec) []*entity.Interface {
	var out = make([]*entity.Interface, 0, len(in))

	for _, spec := range in {
		typeSpec, ok := spec.(*ast.TypeSpec)
		if !ok {
			continue
		}

		if ifa := iface(typeSpec); ifa != nil {
			out = append(out, ifa)
		}
	}

	return out
}

// declaration extracts interfaces from ast tree.
func declaration(in ast.Decl) []*entity.Interface {
	genDecl, ok := in.(*ast.GenDecl)
	if !ok || genDecl.Tok != token.TYPE {
		return nil
	}

	return ifaces(genDecl.Specs)
}

// declarations extracts interfaces from ast tree.
func declarations(in []ast.Decl) []*entity.Interface {
	var out = make([]*entity.Interface, 0, len(in))
	for _, spec := range in {
		if ifas := declaration(spec); len(ifas) > 0 {
			out = append(out, ifas...)
		}
	}

	return out
}

// imports extracts imports from ast tree.
func imports(in []*ast.ImportSpec) []*entity.Import {
	var out = make([]*entity.Import, 0, len(in))
	for _, imp := range in {
		var el entity.Import
		if imp.Name != nil {
			el.Alias = imp.Name.Name
		}

		el.Source = imp.Path.Value

		out = append(out, &el)
	}

	return out
}
