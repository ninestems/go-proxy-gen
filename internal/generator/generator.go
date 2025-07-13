// Package generator describe how app works.
package generator

import (
	"github.com/ninestems/go-proxy-gen/internal"
)

// Generator provides functionality for reading source files and generating proxies for them
type Generator struct {
	parser  internal.ParserI
	definer internal.DefinerI
}

// New builds new Generator.
func New(opts ...Option) *Generator {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}

	return &Generator{
		parser:  cfg.parser,
		definer: cfg.definer,
	}
}
