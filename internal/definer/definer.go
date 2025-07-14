// Package definer describe how app generates proxy layers,
package definer

import (
	"github.com/ninestems/go-proxy-gen/internal"
	"github.com/ninestems/go-proxy-gen/pkg/log"
)

var _ internal.DefinerI = (*Definer)(nil)

// Definer generates proxy layer files based on special markdown annotations.
type Definer struct {
	opt     *options
	proxier internal.ProxierI
	emitter internal.EmitterI
}

// New builds new instance of Definer.
func New(opts ...Option) *Definer {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}

	log.Debugf("definer initialized")
	return &Definer{
		opt:     &cfg,
		proxier: cfg.proxier,
		emitter: cfg.emitter,
	}
}
