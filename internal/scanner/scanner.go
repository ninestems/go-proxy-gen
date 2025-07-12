// Package scanner describe how app read files and build it to markdown.
package scanner

// Scanner contains logic how scan source and format in to markdown.
type Scanner struct{}

// New build new instance of Scanner.
func New() *Scanner {
	return &Scanner{}
}
