package scanner

import (
	"go/ast"
	"go/token"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_extractPackage(t *testing.T) {
	type args struct {
		in       *ast.File
		relative string
		names    []string
	}
	tests := []struct {
		name string
		args args
		want *packageSpecification
	}{
		{
			name: "convert ast file to package",
			args: args{
				in: &ast.File{
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
				relative: "relative",
			},
			want: &packageSpecification{
				name:     "sample",
				relative: "relative",
				imports: []*importSpecification{
					{
						alias: "add",
						path:  "some/path",
					},
				},
				ifasec: []*interfaceSpecification{
					{
						name: "Calculator",
						functions: []*functionSpecification{
							{
								name: "Add",
								input: []*parameterSpecification{
									{
										names:   []string{"a"},
										source:  "int",
										pointer: false,
									},
									{
										names:   []string{"b"},
										source:  "int",
										pointer: false,
									},
								},
								output: []*parameterSpecification{
									{
										names:   []string{},
										source:  "int",
										pointer: false,
									},
								},
								tags: &tagSpecification{},
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, extractPackage(tt.args.in, tt.args.relative, tt.args.names...))
		})
	}
}
