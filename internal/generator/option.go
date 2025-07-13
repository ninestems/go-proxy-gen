package generator

import (
	"github.com/ninestems/go-proxy-gen/internal"
)

type options struct {
	parser  internal.ParserI
	definer internal.DefinerI
}

// Option describe function for applying config.
type Option func(*options)

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
