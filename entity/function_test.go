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
					NewInputParameter("id", "int"),
				},
				output: []*Parameter{
					NewInputParameter("out0", "error"),
				},
				ctxTag: []*ContextIO{
					NewIOContextTag("log_traceID", "context.Context", "trace_id", ProxyTypeLogger),
				},
				inputTag: []*InputIO{
					NewIOInputTag("object_id", "id", "int", "id", ProxyTypeLogger),
				},
				outputTag: []*OutputIO{
					NewIOOutputTag("out0", "out0", "error", "out0", ProxyTypeLogger),
				},
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

func TestFunction_Prepare(t *testing.T) {

	type fields struct {
		input  []*Parameter
		output []*Parameter
	}
	tests := []struct {
		name   string
		fields fields
		want   fields
	}{
		{
			name: "correct fill parameter names",
			fields: fields{
				input: []*Parameter{
					NewInputParameter("", "context.Context"),
					NewInputParameter("", "int"),
				},
				output: []*Parameter{
					NewOutputParameter("", "error"),
				},
			},
			want: fields{
				input: []*Parameter{
					NewInputParameter("in0", "context.Context"),
					NewInputParameter("in1", "int"),
				},
				output: []*Parameter{
					NewOutputParameter("out0", "error"),
				},
			},
		},
		{
			name: "part fill parameter names",
			fields: fields{
				input: []*Parameter{
					NewInputParameter("ctx", "context.Context"),
					NewInputParameter("id", "int"),
				},
				output: []*Parameter{
					NewOutputParameter("", "error"),
				},
			},
			want: fields{
				input: []*Parameter{
					NewInputParameter("ctx", "context.Context"),
					NewInputParameter("id", "int"),
				},
				output: []*Parameter{
					NewOutputParameter("out0", "error"),
				},
			},
		},
		{
			name: "no change parameter names",
			fields: fields{
				input: []*Parameter{
					NewInputParameter("ctx", "context.Context"),
					NewInputParameter("id", "int"),
				},
				output: []*Parameter{
					NewOutputParameter("err", "error"),
				},
			},
			want: fields{
				input: []*Parameter{
					NewInputParameter("ctx", "context.Context"),
					NewInputParameter("id", "int"),
				},
				output: []*Parameter{
					NewOutputParameter("err", "error"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				input:  tt.fields.input,
				output: tt.fields.output,
			}
			f.Prepare()
			require.Equal(t, tt.want.input, f.Input())
			require.Equal(t, tt.want.output, f.Output())
		})
	}
}

func TestFunction_LinkParameters(t *testing.T) {
	type fields struct {
		input  []*Parameter
		output []*Parameter
		tags   *Tags
	}
	type want struct {
		tags *Tags
	}
	tests := []struct {
		name   string
		fields fields
		want   want
	}{
		{
			name: "correct link parameters",
			fields: fields{
				input: []*Parameter{
					NewInputParameter("ctx", "context.Context"),
					NewInputParameter("id", "int"),
				},
				output: []*Parameter{
					NewOutputParameter("", "error"),
				},
				tags: &Tags{
					context: []*ContextIO{
						NewIOContextTag("trace_id", "context.Context", "trace_id", ProxyTypeLogger),
					},
					input: []*InputIO{
						NewIOInputTag("object_id", "", "int", "id", ProxyTypeLogger),
					},
					output: []*OutputIO{
						NewIOOutputTag("custom error", "", "error", "", ProxyTypeLogger),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &Function{
				input:  tt.fields.input,
				output: tt.fields.output,
				tags:   tt.fields.tags,
			}
			f.LinkParameters()
			for _, tag := range f.LogContextTags() {
				require.True(t, !tag.IsEmptyParameter())
			}
			for _, tag := range f.LogInputTags() {
				require.True(t, !tag.IsEmptyParameter())
			}
			for _, tag := range f.LogOutputTags() {
				require.True(t, !tag.IsEmptyParameter())
			}
		})
	}
}
