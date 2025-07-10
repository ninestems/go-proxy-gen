package emitter

type options struct {
	path string
}

// Option describe function for applying config.
type Option func(*options)

// WithPath added path where locate files after generate.
func WithPath(in string) Option {
	return func(o *options) {
		o.path = in
	}
}
