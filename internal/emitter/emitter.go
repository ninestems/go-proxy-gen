// Package emitter describe how app generate saves proxy data to disk.
package emitter

import (
	"go-proxy-gen/internal"
)

var _ internal.EmitterI = (*Emitter)(nil)

// Emitter describes logic saving bytes for file on disk
type Emitter struct {
	path string
}

// New builds new instance of Emitter.
func New(opts ...Option) *Emitter {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}

	return &Emitter{
		path: cfg.path,
	}
}
