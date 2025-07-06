// Package scanner describe how app read files and build it to markdown.
package scanner

import (
	"go-proxy-gen/entity"
	"go-proxy-gen/internal/parser"
)

// Scanner contains logic how scan source and format in to markdown.
type Scanner struct{}

// New build new instance of Scanner.
func New(options ...parser.Option) *Scanner {
	return &Scanner{}
}

// Scan parses the specified path and returns a list of interfaces
// filtered by optional names (if provided).
func (s *Scanner) Scan(in string, names ...string) ([]entity.Interface, error) {
	return nil, nil
}
