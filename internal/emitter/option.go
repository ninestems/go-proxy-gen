package emitter

type options struct {
	paths []string // out path for saving files.
}

// Option describe function for applying config.
type Option func(*options)

// WithPath added paths where locate files after generate.
func WithPath(in ...string) Option {
	return func(o *options) {
		o.paths = append(o.paths, in...)
	}
}
