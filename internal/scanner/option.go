package scanner

type options struct {
	relative string
	ifaces   []string
}

// Option describe function for applying config.
type Option func(*options)

// WithRelativePath added relative path where with go mod name..
func WithRelativePath(in string) Option {
	return func(o *options) {
		o.relative = in
	}
}

// WithIfaces added list name of interfaces for proxy generate in options.
func WithIfaces(in []string) Option {
	return func(o *options) {
		o.ifaces = in
	}
}
