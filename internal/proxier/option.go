package proxier

import (
	"go-proxy-gen/internal"
)

type options struct {
	// lt logger templater.
	lt internal.TemplaterI
	// tt tracer templater.
	tt internal.TemplaterI
	// rt retry templater.
	rt internal.TemplaterI
}

// Option describe function for applying config.
type Option func(*options)

// WithLoggerTemplater added logger templater.
func WithLoggerTemplater(in internal.TemplaterI) Option {
	return func(o *options) {
		o.lt = in
	}
}

// WithTracerTemplater added tracer templater.
func WithTracerTemplater(in internal.TemplaterI) Option {
	return func(o *options) {
		o.tt = in
	}
}

// WithRetrierTemplater added logger templater.
func WithRetrierTemplater(in internal.TemplaterI) Option {
	return func(o *options) {
		o.rt = in
	}
}
