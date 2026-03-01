package proxier

import (
	"github.com/ninestems/go-proxy-gen/internal"
)

type enables struct {
	logger  bool // enable logger generate if interface has tags.
	tracer  bool // enable tracer generate if interface has tags.
	retrier bool // enable retrier generate if interface has tags.
}

type options struct {
	// enables sets enable for generation layers
	enables enables

	templates []internal.TemplaterI
	// lt logger templater.
	lt internal.TemplaterI
	// tt tracer templater.
	tt internal.TemplaterI
	// rt retry templater.
	rt internal.TemplaterI
	// ct custom templater
	ct internal.TemplaterI
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

// WithTemplater added custom templater.
func WithTemplater(in internal.TemplaterI) Option {
	return func(o *options) {
		o.templates = append(o.templates, in)
	}
}

// WithEnableLoggerTemplater enables/disables generation proxy logger decorator even if empty tags.
func WithEnableLoggerTemplater(in bool) Option {
	return func(o *options) {
		o.enables.logger = in
	}
}

// WithEnableTracerTemplater enables/disables generation proxy tracer decorator even if empty tags.
func WithEnableTracerTemplater(in bool) Option {
	return func(o *options) {
		o.enables.tracer = in
	}
}

// WithEnableRetrierTemplater enables/disables generation proxy retrier decorator even if empty tags.
func WithEnableRetrierTemplater(in bool) Option {
	return func(o *options) {
		o.enables.retrier = in
	}
}
