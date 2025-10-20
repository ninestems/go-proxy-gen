package scanner

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_extractImports(t *testing.T) {
	type args struct {
		in []*ast.ImportSpec
	}
	tests := []struct {
		name string
		args args
		want []*importSpecification
	}{
		{
			name: "import conversion",
			args: args{
				in: []*ast.ImportSpec{
					{
						Name: ast.NewIdent("add"),
						Path: &ast.BasicLit{
							Value: "some/path",
						},
					},
				},
			},
			want: []*importSpecification{
				{
					alias: "add",
					path:  "some/path",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, extractImports(tt.args.in))
		})
	}
}
