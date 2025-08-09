package validator

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ninestems/go-proxy-gen/entity"
)

func Test_validatePackage(t *testing.T) {
	type args struct {
		in *entity.Package
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "got error",
			args: args{
				in: entity.NewPackage(
					"",
					"",
					nil,
					nil,
				),
			},
			err: entity.ErrEmptyPackageName,
		},
		{
			name: "no error",
			args: args{
				in: entity.NewPackage(
					"some name",
					"",
					nil,
					nil,
				),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.ErrorIs(t, validatePackage(tt.args.in), tt.err)
		})
	}
}

func Test_validateInterface(t *testing.T) {
	type args struct {
		in *entity.Interface
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "no name error",
			args: args{
				in: entity.NewInterface("", nil),
			},
			err: entity.ErrEmptyInterfaceName,
		},
		{
			name: "no functions",
			args: args{
				in: entity.NewInterface("some name", nil),
			},
			err: entity.ErrEmptyFunctionName,
		},
		{
			name: "no error",
			args: args{
				in: entity.NewInterface("some name", []*entity.Function{{}}),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.ErrorIs(t, validateInterface(tt.args.in), tt.err)
		})
	}
}

func Test_validateFunction(t *testing.T) {
	type args struct {
		in *entity.Function
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "empty name",
			args: args{
				in: entity.NewFunction("", nil, nil, nil),
			},
			err: entity.ErrEmptyFunctionType,
		},
		{
			name: "have no context",
			args: args{
				in: entity.NewFunction("some", nil, nil, nil),
			},
			err: entity.ErrEmptyContext,
		},
		{
			name: "no error",
			args: args{
				in: entity.NewFunction("some", []*entity.Parameter{entity.NewInputParameter("", "context.Context")}, nil, nil),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.ErrorIs(t, validateFunction(tt.args.in), tt.err)
		})
	}
}

func Test_validateTags(t *testing.T) {
	type args struct {
		in         entity.Tags
		ctx        []*entity.ContextIO
		input      []*entity.InputIO
		iparameter *entity.Parameter
		output     []*entity.OutputIO
		oparameter *entity.Parameter
	}
	tests := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "context io undefined",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("", "", "", entity.ProxyTypeUndefined),
				},
				input:  nil,
				output: nil,
			},
			err: entity.ErrInvalidContextTagProxyType,
		},
		{
			name: "context io empty alias",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("", "", "", entity.ProxyTypeLogger),
				},
				input:  nil,
				output: nil,
			},
			err: entity.ErrEmptyContextTagAlias,
		},
		{
			name: "context io empty key",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "", entity.ProxyTypeLogger),
				},
				input:  nil,
				output: nil,
			},
			err: entity.ErrEmptyContextTagKey,
		},
		{
			name: "context io success",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input:  nil,
				output: nil,
			},
			err: nil,
		},
		{
			name: "input io undefined",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("", "", "", "", entity.ProxyTypeUndefined),
				},
				output: nil,
			},
			err: entity.ErrInvalidInputTagProxyType,
		},
		{
			name: "input io empty alias",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("", "", "", "", entity.ProxyTypeLogger),
				},
				output: nil,
			},
			err: entity.ErrEmptyInputTagAlias,
		},
		{
			name: "input io empty key",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("some alias", "", "", "", entity.ProxyTypeLogger),
				},
				output: nil,
			},
			err: entity.ErrEmptyInputTagKey,
		},
		{
			name: "input io empty key",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("some alias", "", "", "some key", entity.ProxyTypeLogger),
				},
				output: nil,
			},
			err: entity.ErrEmptyInputTagName,
		},
		{
			name: "input io empty source",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("some alias", "some name", "", "some key", entity.ProxyTypeLogger),
				},
				output: nil,
			},
			err: entity.ErrEmptyInputTagSource,
		},
		{
			name: "input io empty parameter",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("some alias", "some name", "some source", "some key", entity.ProxyTypeLogger),
				},
				output: nil,
			},
			err: entity.ErrEmptyInputTagParameter,
		},
		{
			name: "input io success",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("some alias", "in0", "some source", "some key", entity.ProxyTypeLogger),
				},
				iparameter: entity.NewInputParameter("in0", "*some source"),
			},
			err: nil,
		},
		{
			name: "output io undefined",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("some alias", "in0", "some source", "some key", entity.ProxyTypeLogger),
				},
				iparameter: entity.NewInputParameter("in0", "*some source"),
				output: []*entity.OutputIO{
					entity.NewIOOutputTag("", "", "", "", entity.ProxyTypeUndefined),
				},
			},
			err: entity.ErrInvalidOutputTagProxyType,
		},
		{
			name: "output io empty alias",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("some alias", "in0", "some source", "some key", entity.ProxyTypeLogger),
				},
				iparameter: entity.NewInputParameter("in0", "*some source"),
				output: []*entity.OutputIO{
					entity.NewIOOutputTag("", "", "", "", entity.ProxyTypeLogger),
				},
			},
			err: entity.ErrEmptyOutputTagAlias,
		},
		{
			name: "output io empty key",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("some alias", "in0", "some source", "some key", entity.ProxyTypeLogger),
				},
				iparameter: entity.NewInputParameter("in0", "*some source"),
				output: []*entity.OutputIO{
					entity.NewIOOutputTag("some alias", "", "", "", entity.ProxyTypeLogger),
				},
			},
			err: entity.ErrEmptyOutputTagKey,
		},
		{
			name: "output io empty name",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("some alias", "in0", "some source", "some key", entity.ProxyTypeLogger),
				},
				iparameter: entity.NewInputParameter("in0", "*some source"),
				output: []*entity.OutputIO{
					entity.NewIOOutputTag("some alias", "", "", "some key", entity.ProxyTypeLogger),
				},
			},
			err: entity.ErrEmptyOutputTagName,
		},
		{
			name: "input io empty source",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("some alias", "in0", "some source", "some key", entity.ProxyTypeLogger),
				},
				iparameter: entity.NewInputParameter("in0", "*some source"),
				output: []*entity.OutputIO{
					entity.NewIOOutputTag("some alias", "some name", "", "some key", entity.ProxyTypeLogger),
				},
			},
			err: entity.ErrEmptyOutputTagSource,
		},
		{
			name: "input io empty parameter",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("some alias", "in0", "some source", "some key", entity.ProxyTypeLogger),
				},
				iparameter: entity.NewInputParameter("in0", "*some source"),
				output: []*entity.OutputIO{
					entity.NewIOOutputTag("some alias", "some name", "some source", "some key", entity.ProxyTypeLogger),
				},
			},
			err: entity.ErrEmptyOutputTagParameter,
		},

		{
			name: "output io success",
			args: args{
				in: entity.Tags{},
				ctx: []*entity.ContextIO{
					entity.NewIOContextTag("some alias", "", "some key", entity.ProxyTypeLogger),
				},
				input: []*entity.InputIO{
					entity.NewIOInputTag("some alias", "in0", "some source", "some key", entity.ProxyTypeLogger),
				},
				iparameter: entity.NewInputParameter("in0", "*some source"),
				output: []*entity.OutputIO{
					entity.NewIOOutputTag("some alias", "out0", "some source", "some key", entity.ProxyTypeLogger),
				},
				oparameter: entity.NewInputParameter("out0", "*some source"),
			},
			err: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.iparameter != nil {
				tt.args.input[0].ApplyParameter(tt.args.iparameter)
			}
			if tt.args.oparameter != nil {
				tt.args.output[0].ApplyParameter(tt.args.oparameter)
			}
			tt.args.in.AddContext(tt.args.ctx...)
			tt.args.in.AddInput(tt.args.input...)
			tt.args.in.AddOutput(tt.args.output...)
			require.ErrorIs(t, validateTags(&tt.args.in), tt.err)
		})
	}
}
