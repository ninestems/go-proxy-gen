package templater

import _ "embed"

//go:embed files/tracer/ot.tmpl
var tracerOpenTelemetryTemplate string

// Tracer describe ways to get string template of tracer.
type Tracer struct {
	source string
}

// NewTracer builds new instance of Tracer
func NewTracer(source string) *Tracer {
	if source == "" {
		source = tracerOpenTelemetryTemplate
	}
	return &Tracer{source}
}

// Template returns template for logger.
func (t *Tracer) Template() string {
	return t.source
}
