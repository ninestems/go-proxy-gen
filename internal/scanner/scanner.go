// Package scanner describe how app read files and build it to markdown.
package scanner

import (
	"github.com/ninestems/go-proxy-gen/pkg/log"
)

// Scanner contains logic how scan source and format in to markdown.
type Scanner struct {
	names    []string
	relative string
}

// New build new instance of Scanner.
func New(opts ...Option) *Scanner {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}

	log.Debugf("scanner initialized")
	return &Scanner{
		names:    cfg.ifaces,
		relative: cfg.relative,
	}
}
