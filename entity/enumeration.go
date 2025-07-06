package entity

// ProxyType describe end forms for proxy layers.
type ProxyType string

const (
	// ProxyTypeLogger for logger.
	ProxyTypeLogger ProxyType = "logger"
	// ProxyTypeTracer for tracer.
	ProxyTypeTracer ProxyType = "tracer"
	// ProxyTypeRetrier for retrier.
	ProxyTypeRetrier ProxyType = "retrier"
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
