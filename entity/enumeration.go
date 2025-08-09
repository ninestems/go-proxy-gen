package entity

// ParameterType describe types for parameters of function.
type ParameterType uint32

const (
	// ParameterTypeUndefined undefined.
	ParameterTypeUndefined ParameterType = iota
	// ParameterTypeInput for input parameters.
	ParameterTypeInput
	// ParameterTypeOutput for output parameters.
	ParameterTypeOutput
)

// String implements stringer for ParameterType.
func (p ParameterType) String() string {
	switch p {
	case ParameterTypeInput:
		return "in"
	case ParameterTypeOutput:
		return "out"
	default:
		return "undefined"
	}
}

// TagType describe tag type.
type TagType uint32

const (
	// TagTypeUndefined undefined.
	TagTypeUndefined TagType = iota
	// TagTypeContext context type tag.
	TagTypeContext
	// TagTypeInput tag for input setting.
	TagTypeInput
	// TagTypeOutput tag for output setting.
	TagTypeOutput
	// TagTypeRetry tag for retry setting.
	TagTypeRetry
)

// String implements stringer for TagType.
func (p TagType) String() string {
	switch p {
	case TagTypeContext:
		return "ctx"
	case TagTypeInput:
		return "in"
	case TagTypeOutput:
		return "out"
	case TagTypeRetry:
		return "retry"
	default:
		return "undefined"
	}
}

// ProxyType describe end forms for proxy layers.
type ProxyType uint32

const (
	// ProxyTypeUndefined for undefined type.
	ProxyTypeUndefined ProxyType = iota
	// ProxyTypeLogger for logger.
	ProxyTypeLogger
	// ProxyTypeTracer for tracer.
	ProxyTypeTracer
	// ProxyTypeRetrier for retrier.
	ProxyTypeRetrier
)

// String implements stringer for ProxyType.
func (p ProxyType) String() string {
	switch p {
	case ProxyTypeUndefined:
		return "undefined"
	case ProxyTypeLogger:
		return "logger"
	case ProxyTypeTracer:
		return "tracer"
	case ProxyTypeRetrier:
		return "retrier"
	default:
		return "undefined"
	}
}

// ValueType describe which type of source in tag.
type ValueType uint32

const (
	// ValueTypeUndefined for undefined or unknown value type.
	ValueTypeUndefined ValueType = iota
	// ValueTypeBool represents a boolean value (true or false).
	ValueTypeBool
	// ValueTypeInt represents a signed integer of platform-dependent size.
	ValueTypeInt
	// ValueTypeInt8 represents a signed 8-bit integer.
	ValueTypeInt8
	// ValueTypeInt16 represents a signed 16-bit integer.
	ValueTypeInt16
	// ValueTypeInt32 represents a signed 32-bit integer.
	ValueTypeInt32
	// ValueTypeInt64 represents a signed 64-bit integer.
	ValueTypeInt64
	// ValueTypeUint represents an unsigned integer of platform-dependent size.
	ValueTypeUint
	// ValueTypeUint8 represents an unsigned 8-bit integer.
	ValueTypeUint8
	// ValueTypeUint16 represents an unsigned 16-bit integer.
	ValueTypeUint16
	// ValueTypeUint32 represents an unsigned 32-bit integer.
	ValueTypeUint32
	// ValueTypeUint64 represents an unsigned 64-bit integer.
	ValueTypeUint64
	// ValueTypeUintptr represents an unsigned integer large enough to store a pointer address.
	ValueTypeUintptr
	// ValueTypeFloat32 represents a 32-bit floating point number.
	ValueTypeFloat32
	// ValueTypeFloat64 represents a 64-bit floating point number.
	ValueTypeFloat64
	// ValueTypeComplex64 represents a complex number with float32 real and imaginary parts.
	ValueTypeComplex64
	// ValueTypeComplex128 represents a complex number with float64 real and imaginary parts.
	ValueTypeComplex128
	// ValueTypeString represents a string value.
	ValueTypeString
	// ValueTypeByte represents a byte value (alias for uint8).
	ValueTypeByte
	// ValueTypeRune represents a rune value (alias for int32, Unicode code point).
	ValueTypeRune
	// ValueTypeStruct represents a struct type.
	ValueTypeStruct
)

// ProxyLogger describe various logger types for proxy logger.
type ProxyLogger uint32

const (
	// ProxyLoggerUndefined for undefined type.
	ProxyLoggerUndefined ProxyLogger = iota
	// ProxyLoggerZap flag for zap implementation.
	ProxyLoggerZap
)

// ProxyTracer describe various tracer types for proxy tracer.
type ProxyTracer uint32

const (
	// ProxyTracerUndefined for undefined type.
	ProxyTracerUndefined ProxyTracer = iota
	// ProxyTracerOpenTelemetry flag for open telemetry implementation.
	ProxyTracerOpenTelemetry
)

// ProxyRetrier describe various tracer types for proxy retrier.
type ProxyRetrier uint32

const (
	// ProxyRetrierUndefined for undefined type.
	ProxyRetrierUndefined ProxyRetrier = iota
	// ProxyRetrierBackoff flag for backoff implementation.
	ProxyRetrierBackoff
)
