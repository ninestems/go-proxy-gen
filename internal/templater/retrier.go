// Package templater describe templates using for generate layers.
package templater

import (
	_ "embed"

	"github.com/ninestems/go-proxy-gen/pkg/log"
)

var (
	//go:embed files/retrier/backoff/backoff.tmpl
	retrierBackoffTemplate string
)

// Retrier describe ways to get string template of retrier.
type Retrier struct {
	source string
}

// NewRetrier builds new instance of Retrier.
func NewRetrier(source string) *Retrier {
	if source == "" {
		source = commonTemplate + retrierBackoffTemplate
	}

	log.Debugf("templater retrier initialized")
	return &Retrier{source}
}

// Template returns template for retrier.
func (l *Retrier) Template() string {
	return l.source
}
