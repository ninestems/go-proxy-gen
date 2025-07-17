package scanner

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ninestems/go-proxy-gen/entity"
)

func Test_packag(t *testing.T) {
	type args struct {
		node  *ast.File
		names []string
	}
	tests := []struct {
		name string
		args args
		want *entity.Package
	}{
		{
			name: "convert file to package",
			args: args{
				node: &ast.File{
					Name: ast.NewIdent("sample"),
					Decls: []ast.Decl{
						&ast.GenDecl{
							Doc: &ast.CommentGroup{
								List: []*ast.Comment{
									{Text: "// Calculator интерфейс для арифметических операций"},
								},
							},
							Tok: token.TYPE,
							Specs: []ast.Spec{
								&ast.TypeSpec{
									Name: ast.NewIdent("Calculator"),
									Type: &ast.InterfaceType{
										Methods: &ast.FieldList{
											List: []*ast.Field{
												{
													Doc: &ast.CommentGroup{
														List: []*ast.Comment{
															{Text: "// Add складывает два числа"},
															{Text: "// Возвращает сумму"},
														},
													},
													Names: []*ast.Ident{ast.NewIdent("Add")},
													Type: &ast.FuncType{
														Params: &ast.FieldList{
															List: []*ast.Field{
																{
																	Names: []*ast.Ident{ast.NewIdent("a")},
																	Type:  ast.NewIdent("int"),
																},
																{
																	Names: []*ast.Ident{ast.NewIdent("b")},
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
									},
								},
							},
						},
					},
					Imports: []*ast.ImportSpec{
						{
							Name: ast.NewIdent("add"),
							Path: &ast.BasicLit{
								Value: "some/path",
							},
						},
					},
				},
				names: nil,
			},
			want: entity.NewPackage(
				"sample",
				"",
				[]*entity.Import{entity.NewImport("add", "some/path")},
				[]*entity.Interface{
					entity.NewInterface(
						"Calculator",
						[]*entity.Function{
							entity.NewFunction(
								"Add",
								[]*entity.Parameter{
									entity.NewInputParameter("a", "int"),
									entity.NewInputParameter("b", "int"),
								},
								[]*entity.Parameter{
									entity.NewOutputParameter("", "int"),
								},
								&entity.Tags{},
							),
						},
					),
				},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, packag(tt.args.node, tt.args.names...))
		})
	}
}
