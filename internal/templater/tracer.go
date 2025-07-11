package templater

import _ "embed"

//go:embed files/tracer/ot.tmpl
var tracerOpenTelemetryTemplate string

type Tracer struct {
	source string
}

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
