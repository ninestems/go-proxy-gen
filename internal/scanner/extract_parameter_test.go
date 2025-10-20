package scanner

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_extractFunctionParameters(t *testing.T) {
	type args struct {
		in []*ast.Field
	}
	tests := []struct {
		name string
		args args
		want []*parameterSpecification
	}{
		{
			name: "extract function input parameters",
			args: args{
				in: []*ast.Field{
					{
						Names: []*ast.Ident{ast.NewIdent("a1"), ast.NewIdent("a2")},
						Type:  ast.NewIdent("int"),
					},
					{
						Names: []*ast.Ident{ast.NewIdent("b")},
						Type:  ast.NewIdent("Example"),
					},
					{
						Names: []*ast.Ident{ast.NewIdent("c")},
						Type:  ast.NewIdent("*source.Example"),
					},
				},
			},
			want: []*parameterSpecification{
				{
					names:   []string{"a1", "a2"},
					source:  "int",
					pointer: false,
				},
				{
					names:   []string{"b"},
					source:  "Example",
					pointer: false,
				},
				{
					names:   []string{"c"},
					source:  "*source.Example",
					pointer: true,
				},
			},
		},
		{
			name: "extract function output parameters",
			args: args{
				in: []*ast.Field{
					{
						Type: ast.NewIdent("Example"),
					},
					{
						Type: ast.NewIdent("error"),
					},
				},
			},
			want: []*parameterSpecification{
				{
					names:  []string{},
					source: "Example",
				},
				{
					names:  []string{},
					source: "error",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, extractFunctionParameters(tt.args.in))
		})
	}
}
