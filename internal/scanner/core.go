package scanner

import (
	"fmt"
	"go/ast"
	"strings"
)

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
