package scanner

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ninestems/go-proxy-gen/entity"
)

func Test_parameterSpecification_BuildInput(t *testing.T) {
	type fields struct {
		names   []string
		source  string
		pointer bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *entity.Parameter
	}{
		{
			name: "not pointer build",
			fields: fields{
				names:   []string{"name"},
				source:  "int",
				pointer: false,
			},
			want: entity.NewInputParameter([]string{"name"}, "int", false),
		},
		{
			name: "pointer build",
			fields: fields{
				names:   []string{"name"},
				source:  "int",
				pointer: true,
			},
			want: entity.NewInputParameter([]string{"name"}, "int", true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &parameterSpecification{
				names:   tt.fields.names,
				source:  tt.fields.source,
				pointer: tt.fields.pointer,
			}
			require.Equal(t, tt.want, s.BuildInput())
		})
	}
}

func Test_parameterSpecification_BuildOutput(t *testing.T) {
	type fields struct {
		names   []string
		source  string
		pointer bool
	}
	tests := []struct {
		name   string
		fields fields
		want   *entity.Parameter
	}{
		{
			name: "not pointer build",
			fields: fields{
				names:   []string{"name"},
				source:  "int",
				pointer: false,
			},
			want: entity.NewOutputParameter([]string{"name"}, "int", false),
		},
		{
			name: "pointer build",
			fields: fields{
				names:   []string{"name"},
				source:  "int",
				pointer: true,
			},
			want: entity.NewOutputParameter([]string{"name"}, "int", true),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &parameterSpecification{
				names:   tt.fields.names,
				source:  tt.fields.source,
				pointer: tt.fields.pointer,
			}
			require.Equal(t, tt.want, s.BuildOutput())
		})
	}
}
