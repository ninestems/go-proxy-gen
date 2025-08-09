package entity

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCommon(t *testing.T) {
	type args struct {
		ttype TagType
		ptype ProxyType
		vtype ValueType
	}
	type result struct {
		isLogger  bool
		isTracer  bool
		isRetrier bool
	}
	tests := []struct {
		name   string
		args   args
		result result
	}{
		{
			name: "create tag undefined and proxy undefined",
			args: args{
				ttype: TagTypeUndefined,
				ptype: ProxyTypeUndefined,
				vtype: ValueTypeUndefined,
			},
			result: result{
				isLogger:  false,
				isTracer:  false,
				isRetrier: false,
			},
		},
		{
			name: "create tag context and proxy logger",
			args: args{
				ttype: TagTypeContext,
				ptype: ProxyTypeLogger,
				vtype: ValueTypeUndefined,
			},
			result: result{
				isLogger:  true,
				isTracer:  false,
				isRetrier: false,
			},
		},
		{
			name: "create tag input and proxy tracer",
			args: args{
				ttype: TagTypeInput,
				ptype: ProxyTypeTracer,
				vtype: ValueTypeUndefined,
			},
			result: result{
				isLogger:  false,
				isTracer:  true,
				isRetrier: false,
			},
		},
		{
			name: "create tag output and proxy retrier",
			args: args{
				ttype: TagTypeOutput,
				ptype: ProxyTypeRetrier,
				vtype: ValueTypeUndefined,
			},
			result: result{
				isLogger:  false,
				isTracer:  false,
				isRetrier: true,
			},
		},
		{
			name: "create tag retry and proxy retrier",
			args: args{
				ttype: TagTypeRetry,
				ptype: ProxyTypeRetrier,
				vtype: ValueTypeUndefined,
			},
			result: result{
				isLogger:  false,
				isTracer:  false,
				isRetrier: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			obj := NewCommon(tt.args.ttype, tt.args.ptype, tt.args.vtype)
			require.Equal(t, tt.args.ttype, obj.TType())
			require.Equal(t, tt.args.ptype, obj.PType())
			require.Equal(t, tt.args.vtype, obj.VType())

			require.Equal(t, tt.result.isLogger, obj.IsForLogger())
			require.Equal(t, tt.result.isTracer, obj.IsForTracer())
			require.Equal(t, tt.result.isRetrier, obj.IsForRetrier())
		})
	}
}
