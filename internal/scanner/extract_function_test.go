package scanner

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_extractFunctionTypes(t *testing.T) {
	type args struct {
		in []*ast.Field
	}
	tests := []struct {
		name string
		args args
		want []*functionSpecification
	}{
		{
			name: "",
			args: args{
				in: []*ast.Field{
					{
						Names: []*ast.Ident{ast.NewIdent("Add")},
						Type: &ast.FuncType{
							Params: &ast.FieldList{
								List: []*ast.Field{
									{
										Names: []*ast.Ident{ast.NewIdent("a")},
										Type:  ast.NewIdent("int"),
									},
								},
							},
							Results: &ast.FieldList{
								List: []*ast.Field{
									{Type: ast.NewIdent("int")},
								},
							},
						},
					},
				},
			},
			want: []*functionSpecification{
				{
					name:   "Add",
					input:  []*parameterSpecification{{names: []string{"a"}, source: "int"}},
					output: []*parameterSpecification{{names: []string{}, source: "int"}},
					tags:   &tagSpecification{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, extractFunctionTypes(tt.args.in))
		})
	}
}
