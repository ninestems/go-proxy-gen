package parser

type options struct {
	in     string
	ifaces []string
}

// Option describe function for applying config.
type Option func(*options)

// WithInPath added path where locate files to options.
func WithInPath(in string) Option {
	return func(o *options) {
		o.in = in
	}
}

// WithIfaces added list name of interfaces for proxy generate in options.
func WithIfaces(in []string) Option {
	return func(o *options) {
		o.ifaces = in
	}
}
