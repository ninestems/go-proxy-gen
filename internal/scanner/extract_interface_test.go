package scanner

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_extractInterfaceTypeSpecifications(t *testing.T) {
	type args struct {
		in    []*ast.TypeSpec
		names []string
	}
	tests := []struct {
		name string
		args args
		want []*interfaceSpecification
	}{
		{
			name: "interface conversion",
			args: args{
				in: []*ast.TypeSpec{
					{
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
				names: nil,
			},
			want: []*interfaceSpecification{
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
		{
			name: "filtered conversion",
			args: args{
				in: []*ast.TypeSpec{
					{
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
					{
						Name: ast.NewIdent("Filtered"),
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
				names: []string{"Calculator"},
			},
			want: []*interfaceSpecification{
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
		{
			name: "empty conversion",
			args: args{
				in: []*ast.TypeSpec{
					{
						Name: ast.NewIdent("Empty"),
						Type: &ast.InterfaceType{
							Methods: &ast.FieldList{},
						},
					},
				},
				names: nil,
			},
			want: []*interfaceSpecification{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := extractInterfaceTypeSpecifications(tt.args.in, tt.args.names...)
			require.Equal(t, tt.want, got)
		})
	}
}
