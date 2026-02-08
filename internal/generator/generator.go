// Package generator describe how app works.
package generator

import (
	"github.com/ninestems/go-proxy-gen/internal"
	"github.com/ninestems/go-proxy-gen/pkg/log"
)

type Pair struct {
	parser  internal.ParserI
	definer internal.DefinerI
}

// Generator provides functionality for reading source files and generating proxies for them
type Generator struct {
	pairs   []Pair
	parser  internal.ParserI
	definer internal.DefinerI
}

// New builds new Generator.
func New(opts ...Option) *Generator {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}

	pairs := make([]Pair, 0, len(cfg.pairs))
	for _, p := range cfg.pairs {
		pairs = append(
			pairs,
			Pair{
				parser:  p.parser,
				definer: p.definer,
			},
		)
	}

	log.Debugf("generator initialized")
	return &Generator{
		pairs:   pairs,
		parser:  cfg.parser,
		definer: cfg.definer,
	}
}
