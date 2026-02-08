package generator

import (
	"github.com/ninestems/go-proxy-gen/internal"
)

type options struct {
	pairs   []pair
	parser  internal.ParserI
	definer internal.DefinerI
}

type pair struct {
	parser  internal.ParserI
	definer internal.DefinerI
}

// Option describe function for applying config.
type Option func(*options)

func WithPair(in internal.ParserI, def internal.DefinerI) Option {
	return func(o *options) {
		o.pairs = append(o.pairs, pair{in, def})
	}
}

// WithParser added parser in config.
func WithParser(in internal.ParserI) Option {
	return func(o *options) {
		o.parser = in
	}
}

// WithDefiner added definer in config.
func WithDefiner(in internal.DefinerI) Option {
	return func(o *options) {
		o.definer = in
	}
}
