package definer

import (
	"github.com/ninestems/go-proxy-gen/internal"
)

type options struct {
	out     string
	proxier internal.ProxierI
	emitter internal.EmitterI
}

// Option describe function for applying config.
type Option func(*options)

// WithOutPath added path where locate files after generate.
func WithOutPath(in string) Option {
	return func(o *options) {
		o.out = in
	}
}

// WithProxier added proxier in options.
func WithProxier(in internal.ProxierI) Option {
	return func(o *options) {
		o.proxier = in
	}
}

// WithEmitter added emitter in options.
func WithEmitter(in internal.EmitterI) Option {
	return func(o *options) {
		o.emitter = in
	}
}
