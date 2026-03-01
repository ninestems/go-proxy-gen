package proxier

import (
	"github.com/ninestems/go-proxy-gen/internal"
)

type options struct {
	// templates list of proxy metainfo for layers.
	templates []internal.TemplaterI
}

// Option describe function for applying config.
type Option func(*options)

// WithTemplater added custom templater.
func WithTemplater(in internal.TemplaterI) Option {
	return func(o *options) {
		o.templates = append(o.templates, in)
	}
}
