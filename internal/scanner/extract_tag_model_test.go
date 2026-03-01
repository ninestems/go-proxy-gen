package scanner

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/ninestems/go-proxy-gen/pkg/entity"
)

func Test_tagContextSpecification_Build(t *testing.T) {
	type fields struct {
		proxy  string
		alias  string
		key    string
		source string
	}
	tests := []struct {
		name   string
		fields fields
		want   *entity.ContextIO
	}{
		{
			name: "build proxy logger context tag",
			fields: fields{
				proxy:  "log",
				alias:  "custom key",
				key:    "context_value",
				source: "context.Context",
			},
			want: entity.NewIOContextTag(
				"custom key",
				"context.Context",
				"context_value",
				entity.ProxyTypeLogger,
			),
		},
		{
			name: "build proxy tracer context tag",
			fields: fields{
				proxy:  "trace",
				alias:  "custom key",
				key:    "context_value",
				source: "context.Context",
			},
			want: entity.NewIOContextTag(
				"custom key",
				"context.Context",
				"context_value",
				entity.ProxyTypeTracer,
			),
		},
		{
			name: "build proxy tracer context tag",
			fields: fields{
				proxy:  "amy",
				alias:  "custom key",
				key:    "context_value",
				source: "context.Context",
			},
			want: entity.NewIOContextTag(
				"custom key",
				"context.Context",
				"context_value",
				entity.ProxyTypeUndefined,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &tagContextSpecification{
				proxy:  tt.fields.proxy,
				alias:  tt.fields.alias,
				key:    tt.fields.key,
				source: tt.fields.source,
			}
			require.Equal(t, tt.want, s.Build())
		})
	}
}

func Test_tagInputSpecification_Build(t *testing.T) {
	type fields struct {
		proxy    string
		name     string
		source   string
		accessor string
		alias    string
	}
	tests := []struct {
		name   string
		fields fields
		want   *entity.InputIO
	}{
		{
			name: "build proxy logger input tag",
			fields: fields{
				proxy:    "log",
				name:     "name",
				source:   "int",
				accessor: "name",
				alias:    "custom_name",
			},
			want: entity.NewIOInputTag(
				"custom_name",
				"name",
				"int",
				"name",
				entity.ProxyTypeLogger,
			),
		},
		{
			name: "build proxy tracer input tag",
			fields: fields{
				proxy:    "trace",
				name:     "name",
				source:   "int",
				accessor: "name",
				alias:    "custom_name",
			},
			want: entity.NewIOInputTag(
				"custom_name",
				"name",
				"int",
				"name",
				entity.ProxyTypeTracer,
			),
		},
		{
			name: "build proxy tracer input tag",
			fields: fields{
				proxy:    "any",
				name:     "name",
				source:   "int",
				accessor: "name",
				alias:    "custom_name",
			},
			want: entity.NewIOInputTag(
				"custom_name",
				"name",
				"int",
				"name",
				entity.ProxyTypeUndefined,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &tagInputSpecification{
				proxy:    tt.fields.proxy,
				name:     tt.fields.name,
				source:   tt.fields.source,
				accessor: tt.fields.accessor,
				alias:    tt.fields.alias,
			}
			require.Equal(t, tt.want, s.Build())
		})
	}
}

func Test_tagOutputSpecification_Build(t *testing.T) {
	type fields struct {
		proxy    string
		name     string
		source   string
		accessor string
		alias    string
	}
	tests := []struct {
		name   string
		fields fields
		want   *entity.OutputIO
	}{
		{
			name: "build proxy logger input tag",
			fields: fields{
				proxy:    "log",
				name:     "name",
				source:   "int",
				accessor: "name",
				alias:    "custom_name",
			},
			want: entity.NewIOOutputTag(
				"custom_name",
				"name",
				"int",
				"name",
				entity.ProxyTypeLogger,
			),
		},
		{
			name: "build proxy tracer input tag",
			fields: fields{
				proxy:    "trace",
				name:     "name",
				source:   "int",
				accessor: "name",
				alias:    "custom_name",
			},
			want: entity.NewIOOutputTag(
				"custom_name",
				"name",
				"int",
				"name",
				entity.ProxyTypeTracer,
			),
		},
		{
			name: "build proxy tracer input tag",
			fields: fields{
				proxy:    "any",
				name:     "name",
				source:   "int",
				accessor: "name",
				alias:    "custom_name",
			},
			want: entity.NewIOOutputTag(
				"custom_name",
				"name",
				"int",
				"name",
				entity.ProxyTypeUndefined,
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &tagOutputSpecification{
				proxy:    tt.fields.proxy,
				name:     tt.fields.name,
				source:   tt.fields.source,
				accessor: tt.fields.accessor,
				alias:    tt.fields.alias,
			}
			require.Equal(t, tt.want, s.Build())
		})
	}
}

func Test_tagRetrySpecification_Build(t *testing.T) {
	type fields struct {
		proxy      string
		start      time.Duration
		end        time.Duration
		multiplier float32
		attempts   uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   *entity.Retry
	}{
		{
			name: "build retry tag",
			fields: fields{
				proxy:      "retry",
				start:      time.Second,
				end:        time.Minute,
				multiplier: 1.2,
				attempts:   4,
			},
			want: entity.NewRetryTag(time.Second, time.Minute, 1.2, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &tagRetrySpecification{
				proxy:      tt.fields.proxy,
				start:      tt.fields.start,
				end:        tt.fields.end,
				multiplier: tt.fields.multiplier,
				attempts:   tt.fields.attempts,
			}
			require.Equal(t, tt.want, s.Build())
		})
	}
}
