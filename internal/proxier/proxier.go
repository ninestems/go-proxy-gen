// Package proxier describe how app generate proxy layers by markdown.
package proxier

import (
	"go-proxy-gen/internal"
)

var _ internal.ProxierI = (*Proxier)(nil)

// Proxier describes logic generate of proxy layers
type Proxier struct {
	// lt logger templater.
	lt internal.TemplaterI
	// tt tracer templater.
	tt internal.TemplaterI
	// rt retry templater.
	rt internal.TemplaterI
}

// New builds new instance of Emitter.
func New(opts ...Option) *Proxier {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}

	return &Proxier{
		lt: cfg.lt,
		tt: cfg.tt,
		rt: cfg.rt,
	}
}
