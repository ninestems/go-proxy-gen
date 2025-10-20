package scanner

import (
	"go/ast"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_extractFunctionTags(t *testing.T) {
	type args struct {
		in *ast.CommentGroup
	}
	type want struct {
		ctx    []*tagContextSpecification
		input  []*tagInputSpecification
		output []*tagOutputSpecification
		retry  []*tagRetrySpecification
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "good parse tags logger",
			args: args{
				in: &ast.CommentGroup{
					List: []*ast.Comment{
						{
							Text: "// goproxygen:",
						},
						{
							Text: "// log ctx::trace_id::log_traceID",
						},
						{
							Text: "//  log input::in:entity.Example::Field::log_some_input",
						},
						{
							Text: "//  log output::entity.Example::Field::log_some_output",
						},
					},
				},
			},
			want: want{
				ctx: []*tagContextSpecification{{
					proxy:  "log",
					alias:  "log_traceID",
					key:    "trace_id",
					source: "context.Context",
				}},
				input: []*tagInputSpecification{{
					proxy:    "log",
					name:     "in",
					source:   "entity.Example",
					accessor: "Field",
					alias:    "log_some_input",
				}},
				output: []*tagOutputSpecification{{
					proxy:    "log",
					name:     "",
					source:   "entity.Example",
					accessor: "Field",
					alias:    "log_some_output",
				}},
			},
		},
		{
			name: "good parse tags tracer",
			args: args{
				in: &ast.CommentGroup{
					List: []*ast.Comment{
						{
							Text: "// goproxygen:",
						},
						{
							Text: "//  trace ctx::trace_id::trace_traceID",
						},
						{
							Text: "//  trace input::in:entity.Example::Field::trace_some_input",
						},
						{
							Text: "//  trace output::entity.Example::Field::trace_some_output",
						},
					},
				},
			},
			want: want{
				ctx: []*tagContextSpecification{{
					proxy:  "trace",
					alias:  "trace_traceID",
					key:    "trace_id",
					source: "context.Context",
				}},
				input: []*tagInputSpecification{{
					proxy:    "trace",
					name:     "in",
					source:   "entity.Example",
					accessor: "Field",
					alias:    "trace_some_input",
				}},
				output: []*tagOutputSpecification{{
					proxy:    "trace",
					name:     "",
					source:   "entity.Example",
					accessor: "Field",
					alias:    "trace_some_output",
				}},
			},
		},
		{
			name: "good parse with retry",
			args: args{
				in: &ast.CommentGroup{
					List: []*ast.Comment{
						{
							Text: "// goproxygen:",
						},
						{
							Text: "//  technical retry::1s::60s::1.2::4",
						},
					},
				},
			},
			want: want{
				retry: []*tagRetrySpecification{{
					proxy:      "technical",
					start:      time.Second * 1,
					end:        time.Minute * 1,
					multiplier: 1.2,
					attempts:   4,
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := extractFunctionTags(tt.args.in)
			require.Equal(t, tt.want.ctx, out.context)
			require.Equal(t, tt.want.input, out.input)
			require.Equal(t, tt.want.output, out.output)
			require.Equal(t, tt.want.retry, out.retry)
		})
	}
}
