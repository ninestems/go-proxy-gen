package parser

import (
	"github.com/ninestems/go-proxy-gen/internal"
)

type options struct {
	in        string
	scanner   internal.ScannerI
	validator internal.ValidatorI
}

// Option describe function for applying config.
type Option func(*options)

// WithInPath added path where locate files to options.
func WithInPath(in string) Option {
	return func(o *options) {
		o.in = in
	}
}

// WithScanner added scanner in options.
func WithScanner(scanner internal.ScannerI) Option {
	return func(o *options) {
		o.scanner = scanner
	}
}

// WithValidator added validator in options.
func WithValidator(validator internal.ValidatorI) Option {
	return func(o *options) {
		o.validator = validator
	}
}
