// Package proxier describe how app generate proxy layers by markdown.
package proxier

import (
	"github.com/ninestems/go-proxy-gen/internal"
	"github.com/ninestems/go-proxy-gen/pkg/log"
)

var _ internal.ProxierI = (*Proxier)(nil)

// Proxier describes logic generate of proxy layers
type Proxier struct {
	opts options
}

// New builds new instance of Emitter.
func New(opts ...Option) *Proxier {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}

	log.Debugf("proxier initialized")
	return &Proxier{
		opts: cfg,
	}
}
