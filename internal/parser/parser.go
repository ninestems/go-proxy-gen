// Package parser describe how app read interfaces for proxy.
package parser

import (
	"go-proxy-gen/entity"
	"go-proxy-gen/internal"
)

var _ internal.ParserI = (*Parser)(nil)

// Parser helps parse source files and compacts them into a specialized markdown format.
type Parser struct {
	opt       options
	scanner   any
	validator any
}

// New builds new instance of Parser.
func New(opts ...Option) *Parser {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}
	return &Parser{
		opt:       cfg,
		scanner:   nil,
		validator: nil,
	}
}

// Parse parses the specified path and returns a list of interfaces
// filtered by optional names (if provided).
func (p *Parser) Parse() ([]entity.Interface, error) {
	return nil, nil
}
