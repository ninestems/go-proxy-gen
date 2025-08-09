package scanner

import (
	"go/ast"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ninestems/go-proxy-gen/entity"
)

func Test_tags(t *testing.T) {
	type args struct {
		in *ast.CommentGroup
	}
	tests := []struct {
		name       string
		args       args
		wantCtx    []*entity.ContextIO
		wantInput  []*entity.InputIO
		wantOutput []*entity.OutputIO
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
			wantCtx: []*entity.ContextIO{
				entity.NewIOContextTag("log_traceID", "context.Context", "trace_id", entity.ProxyTypeLogger),
			},
			wantInput: []*entity.InputIO{
				entity.NewIOInputTag("log_some_input", "in", "entity.Example", "Field", entity.ProxyTypeLogger),
			},
			wantOutput: []*entity.OutputIO{
				entity.NewIOOutputTag("log_some_output", "", "entity.Example", "Field", entity.ProxyTypeLogger),
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
			wantCtx: []*entity.ContextIO{
				entity.NewIOContextTag("trace_traceID", "context.Context", "trace_id", entity.ProxyTypeTracer),
			},
			wantInput: []*entity.InputIO{
				entity.NewIOInputTag("trace_some_input", "in", "entity.Example", "Field", entity.ProxyTypeTracer),
			},
			wantOutput: []*entity.OutputIO{
				entity.NewIOOutputTag("trace_some_output", "", "entity.Example", "Field", entity.ProxyTypeTracer),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := tags(tt.args.in)
			require.Equal(t, tt.wantCtx, out.Context())
			require.Equal(t, tt.wantInput, out.Input())
			require.Equal(t, tt.wantOutput, out.Output())
		})
	}
}
