// Package templater describe templates using for generate layers.
package templater

import (
	_ "embed"

	"github.com/ninestems/go-proxy-gen/pkg/log"
)

//go:embed files/retrier/backoff/backoff.tmpl
var retrierBackoffTemplate string

// Retrier describe ways to get string template of retrier.
type Retrier struct {
	source string
}

// NewRetrier builds new instance of Retrier.
func NewRetrier(source ...string) *Retrier {
	template := baseRetrier + retrierTemplate + clearTemplate + retrierBackoffTemplate
	if len(source) > 0 && len(source[0]) > 0 {
		template = source[0]
	}

	log.Debugf("templater retrier initialized")
	return &Retrier{template}
}

// Template returns template for retrier.
func (l *Retrier) Template() string {
	return l.source
}
