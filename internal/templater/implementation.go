package templater

import (
	_ "embed"
)

var (
	//go:embed files/logger/zap/zap.tmpl
	loggerZapTemplate string
	//go:embed files/retrier/backoff/backoff.tmpl
	retrierBackoffTemplate string
	//go:embed files/tracer/opentelemetry/opentelemetry.tmpl
	tracerOpenTelemetryTemplate string
)
