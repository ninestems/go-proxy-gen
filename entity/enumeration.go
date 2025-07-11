package entity

type ParameterType uint32

const (
	ParameterTypeUndefined ParameterType = iota
	ParameterTypeInput
	ParameterTypeOutput
)

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

type TagType uint32

const (
	TagTypeUndefined TagType = iota
	TagTypeContext
	TagTypeInput
	TagTypeOutput
	TagTypeRetry
)

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
