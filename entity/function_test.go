package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewFunction(t *testing.T) {
	type args struct {
		name      string
		input     []*Parameter
		output    []*Parameter
		ctxTag    []*ContextIO
		inputTag  []*InputIO
		outputTag []*OutputIO
	}
	tests := []struct {
		name string
		args args
		want *Function
	}{
		{
			name: "empty input&output",
			args: args{
				name:      "fn",
				input:     nil,
				output:    nil,
				ctxTag:    nil,
				inputTag:  nil,
				outputTag: nil,
			},
		},
		{
			name: "full set of params and tags",
			args: args{
				name: "fn",
				input: []*Parameter{
					NewInputParameter("ctx", "context.Context"),
					NewInputParameter("id", "int")},
				output: []*Parameter{NewInputParameter("out0", "error")},
				ctxTag: []*ContextIO{
					NewIOContextTag("log_traceID", "context.Context", "trace_id", ProxyTypeLogger),
				},
				inputTag: []*InputIO{
					NewIOInputTag("object_id", "id", "int", "", ProxyTypeLogger),
				},
				outputTag: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var tags Tags
			tags.AddContext(tt.args.ctxTag...)
			tags.AddInput(tt.args.inputTag...)
			tags.AddOutput(tt.args.outputTag...)

			got := NewFunction(tt.args.name, tt.args.input, tt.args.output, &tags)
			require.Equal(t, tt.args.name, got.Name())
			require.Equal(t, tt.args.input, got.Input())
			require.Equal(t, tt.args.output, got.Output())
			require.Equal(t, &tags, got.Tags())

		})
	}
}
