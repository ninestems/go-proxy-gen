package scanner

import (
	"go/ast"
	"testing"
)

func Test_exprToParameter(t *testing.T) {
	type args struct {
		in ast.Expr
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Ident",
			args: args{
				in: &ast.Ident{Name: "param"},
			},
			want: "param",
		},
		{
			name: "StarExpr",
			args: args{
				in: &ast.StarExpr{
					X: &ast.Ident{Name: "param"},
				},
			},
			want: "*param",
		},
		{
			name: "SelectorExpr",
			args: args{
				in: &ast.SelectorExpr{
					X:   &ast.Ident{Name: "path"},
					Sel: &ast.Ident{Name: "param"},
				},
			},
			want: "path.param",
		},
		{
			name: "ArrayType",
			args: args{
				in: &ast.ArrayType{
					Elt: &ast.Ident{Name: "int"},
				},
			},
			want: "[]int",
		},
		{
			name: "MapType",
			args: args{
				in: &ast.MapType{
					Key:   &ast.Ident{Name: "string"},
					Value: &ast.Ident{Name: "int"},
				},
			},
			want: "map[string]int",
		},
		{
			name: "InterfaceType",
			args: args{
				in: &ast.InterfaceType{},
			},
			want: "interface{}",
		},
		{
			name: "FuncType",
			args: args{
				in: &ast.FuncType{},
			},
			want: "func",
		},
		{
			name: "Ellipsis",
			args: args{
				in: &ast.Ellipsis{
					Elt: &ast.Ident{Name: "string"},
				},
			},
			want: "...string",
		},
		{
			name: "IndexExpr",
			args: args{
				in: &ast.IndexExpr{
					X:     &ast.Ident{Name: "List"},
					Index: &ast.Ident{Name: "int"},
				},
			},
			want: "List[int]",
		},
		{
			name: "IndexListExpr",
			args: args{
				in: &ast.IndexListExpr{
					X: &ast.Ident{Name: "Foo"},
					Indices: []ast.Expr{
						&ast.Ident{Name: "string"},
						&ast.Ident{Name: "int"},
					},
				},
			},
			want: "Foo[string, int]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exprToParameter(tt.args.in); got != tt.want {
				t.Errorf("exprToParameter() = %v, want %v", got, tt.want)
			}
		})
	}
}
