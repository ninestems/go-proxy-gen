package generator

import (
	"github.com/ninestems/go-proxy-gen/internal"
)

type options struct {
	pairs []pair
}

type pair struct {
	parser  internal.ParserI
	definer internal.DefinerI
}

// Option describe function for applying config.
type Option func(*options)

// WithPair saves parser and definer implementations for pair struct.
func WithPair(in internal.ParserI, def internal.DefinerI) Option {
	return func(o *options) {
		o.pairs = append(o.pairs, pair{in, def})
	}
}
