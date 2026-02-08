package templater

import (
	_ "embed"

	"github.com/ninestems/go-proxy-gen/pkg/log"
)

var (
	//go:embed files/tracer/opentelemetry/opentelemetry.tmpl
	tracerOpenTelemetryTemplate string
)

// Tracer describe ways to get string template of tracer.
type Tracer struct {
	source string
}

// NewTracer builds new instance of Tracer
func NewTracer(source ...string) *Tracer {
	template := baseTracer + tracerTemplate + clearTemplate + tracerOpenTelemetryTemplate
	if len(source) > 0 && len(source[0]) > 0 {
		template = source[0]
	}

	log.Debugf("templater tracer initialized")
	return &Tracer{template}
}

// Template returns template for logger.
func (t *Tracer) Template() string {
	return t.source
}
