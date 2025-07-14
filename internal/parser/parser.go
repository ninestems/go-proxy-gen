// Package parser describe how app read interfaces for proxy.
package parser

import (
	"github.com/ninestems/go-proxy-gen/internal"
	"github.com/ninestems/go-proxy-gen/pkg/log"
)

var _ internal.ParserI = (*Parser)(nil)

// Parser helps parse source files and compacts them into a specialized markdown format.
type Parser struct {
	opt       options
	scanner   internal.ScannerI
	validator internal.ValidatorI
}

// New builds new instance of Parser.
func New(opts ...Option) *Parser {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}

	log.Debugf("parser initialized")
	return &Parser{
		opt:       cfg,
		scanner:   cfg.scanner,
		validator: cfg.validator,
	}
}
