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
	var (
		proxyType     entity.ProxyType
		parameterType entity.TagType
		alias         string
		name          string
		path          string
		key           string
	)

	switch in[0] {
	case "log":
		proxyType = entity.ProxyTypeLogger
	case "trace":
		proxyType = entity.ProxyTypeTracer
	case "retry":
		proxyType = entity.ProxyTypeRetrier
	}

	data := strings.Split(in[1], "::")
	switch data[0] {
	case "ctx":
		parameterType = entity.TagTypeContext
		alias, key = extractCtx(data)
	case "input":
		parameterType = entity.TagTypeInput
		alias, name, path, key = extractIO(data)
	case "output":
		parameterType = entity.TagTypeOutput
		alias, name, path, key = extractIO(data)
	}

	out := entity.NewTag(
		proxyType,
		parameterType,
		alias,
		name,
		path,
		key,
	)

	return out
}

// extractCtx extracts data for context tags.
func extractCtx(in []string) (alias, key string) {
	switch len(in) {
	case 2:
		alias = in[1]
		key = in[1]
	case 3:
		alias = in[1]
		key = in[2]
	}

	return
}

func extractIO(in []string) (alias, name, path, key string) {
	switch len(in) {
	case 3:
		alias = in[2]
		name, path = extractPath(in[1])
		key = in[2]
	case 4:
		alias = in[1]
		name, path = extractPath(in[2])
		key = in[3]
	}

	return
}

func extractPath(in string) (name, path string) {
	data := strings.Split(in, ":")
	path = data[0]
	if len(data) > 1 {
		name, path = data[0], data[1]
	}

	return
}

// tags extracts custom tags from all lines of comments.
func tags(in *ast.CommentGroup) []*entity.Tag {
	if in == nil {
		return nil
	}

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

// exprToParameter converts ast expression to target string.
func exprToParameter(in ast.Expr) string {
	switch expr := in.(type) {
	case *ast.Ident:
		return expr.Name
	case *ast.StarExpr:
		return "*" + exprToParameter(expr.X)
	case *ast.SelectorExpr:
		return exprToParameter(expr.X) + "." + expr.Sel.Name
	case *ast.ArrayType:
		return "[]" + exprToParameter(expr.Elt)
	case *ast.MapType:
		return "map[" + exprToParameter(expr.Key) + "]" + exprToParameter(expr.Value)
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.FuncType:
		return "func"
	case *ast.Ellipsis:
		return "..." + exprToParameter(expr.Elt)
	case *ast.IndexExpr: // generic type, e.g. Foo[T]
		return exprToParameter(expr.X) + "[" + exprToParameter(expr.Index) + "]"
	case *ast.IndexListExpr: // generic with multiple parameters, e.g. Foo[T, U]
		parts := make([]string, 0, len(expr.Indices))
		for _, idx := range expr.Indices {
			parts = append(parts, exprToParameter(idx))
		}
		return exprToParameter(expr.X) + "[" + strings.Join(parts, ", ") + "]"
	default:
		return fmt.Sprintf("%T", in) // for debugging purpose.
	}
}

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

// iface extracts interface from ast tree.
func iface(in *ast.TypeSpec) *entity.Interface {
	ifa, ok := in.Type.(*ast.InterfaceType)
	if !ok {
		return nil
	}

	return entity.NewInterface(in.Name.Name, functions(ifa.Methods.List))
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
		alias := ""
		if imp.Name != nil {
			alias = imp.Name.Name
		}

		out = append(out, entity.NewImport(alias, imp.Path.Value))
	}

	return out
}
