package entity

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
type ProxyLogger string

const (
	// ProxyLoggerZap flag for zap implementation.
	ProxyLoggerZap ProxyLogger = "zap"
)

// ProxyTracer describe various tracer types for proxy tracer.
type ProxyTracer string

const (
	// ProxyTracerOpenTelemetry flag for open telemetry implementation.
	ProxyTracerOpenTelemetry ProxyTracer = "openTelemetry"
)
