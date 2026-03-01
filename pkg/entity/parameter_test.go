package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewParameter(t *testing.T) {
	type args struct {
		ptype   ParameterType
		path    string
		names   []string
		pointer bool
	}
	tests := []struct {
		name string
		args args
		want *Parameter
	}{
		{
			name: "simple parameter creation",
			args: args{
				ptype:   ParameterTypeInput,
				path:    "bool",
				names:   []string{},
				pointer: false,
			},
			want: &Parameter{
				CommonParameter: &CommonParameter{
					ptype: ParameterTypeInput,
					vtype: ValueTypeBool,
				},
				names:   []string{},
				path:    "bool",
				pointer: false,
			},
		},
		{
			name: "creating type parameter with relative",
			args: args{
				ptype:   ParameterTypeInput,
				path:    "entity.Struct",
				names:   []string{},
				pointer: false,
			},
			want: &Parameter{
				CommonParameter: &CommonParameter{
					ptype: ParameterTypeInput,
					vtype: ValueTypeStruct,
				},
				names:   []string{},
				path:    "entity.Struct",
				pointer: false,
			},
		},
		{
			name: "creating type parameter with no source relative",
			args: args{
				ptype:   ParameterTypeInput,
				path:    "Struct",
				names:   []string{},
				pointer: false,
			},
			want: &Parameter{
				CommonParameter: &CommonParameter{
					ptype: ParameterTypeInput,
					vtype: ValueTypeStruct,
				},
				names:   []string{},
				path:    "source.Struct",
				pointer: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, NewParameter(tt.args.ptype, tt.args.path, tt.args.names, tt.args.pointer))
		})
	}
}
