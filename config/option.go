// Package config describe how app must be configured.
package config

// Template describes a configuration template for logger, tracer or retrier.
type Template struct {
	Custom string // path to custom template directory (optional)
	Value  string // template value
}

// Path describes source and destination folders.
type Path struct {
	From string // path to the source Go code
	To   string // path to write generated files
}

// Config holds all settings for generation.
type Config struct {
	Debug   bool
	Logger  Template
	Tracer  Template
	Retrier Template
	Path    Path
	Ifaces  []string
}

// Option is a function that modifies the Config.
type Option func(*Config)

// WithDebug enables or disables debug mode.
//
// Now its unused parameter, added for future compatibility.
func WithDebug(debug bool) Option {
	return func(c *Config) {
		c.Debug = debug
	}
}

// WithLogger sets the logger template.
func WithLogger(t Template) Option {
	return func(c *Config) {
		c.Logger = t
	}
}

// WithTracer sets the tracer template.
func WithTracer(t Template) Option {
	return func(c *Config) {
		c.Tracer = t
	}
}

// WithRetrier sets the retrier template.
func WithRetrier(t Template) Option {
	return func(c *Config) {
		c.Retrier = t
	}
}

// WithPath sets the input/output paths.
func WithPath(from, to string) Option {
	return func(c *Config) {
		c.Path.From = from
		c.Path.To = to
	}
}

// WithInterfaces sets the target interface names.
func WithInterfaces(ifaces []string) Option {
	return func(c *Config) {
		c.Ifaces = ifaces
	}
}
