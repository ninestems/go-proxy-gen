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
