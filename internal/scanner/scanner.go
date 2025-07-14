// Package scanner describe how app read files and build it to markdown.
package scanner

import (
	"github.com/ninestems/go-proxy-gen/pkg/log"
)

// Scanner contains logic how scan source and format in to markdown.
type Scanner struct{}

// New build new instance of Scanner.
func New() *Scanner {
	log.Debugf("scanner initialized")
	return &Scanner{}
}
