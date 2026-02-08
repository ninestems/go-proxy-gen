package config

import (
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type App struct {
	Date      string
	Version   string
	GoVersion string
	Debug     bool
}

// Template describes a configuration template for custom tmpl file implementations.
type Template struct {
	Path string // path to custom template
}

type Templates struct {
	Logger  Template
	Tracer  Template
	Retrier Template
}

type Proxy struct {
	Logger  bool // Logger enables generate for proxy logger.
	Tracer  bool // Tracer enables generate for proxy tracer.
	Retrier bool // Retrier enables generate for proxy retrier.
}

// Path describes source and destination folders.
type Path struct {
	In       string   // path to the source Go code
	Relative string   // relative path
	Outwards []string // paths to location for store generated files
	Names    []string // list names of interfaces to include generation
}

type Config struct {
	App       App
	Templates Templates
	Proxy     Proxy
	Paths     []Path
}

func DefaultConfig() Config {
	return Config{
		App: App{
			Date:      time.Now().UTC().Format(time.RFC3339),
			Version:   "UNKNOW",
			GoVersion: runtime.Version(),
			Debug:     false,
		},
		Templates: Templates{
			Logger:  Template{},
			Tracer:  Template{},
			Retrier: Template{},
		},
		Proxy: Proxy{
			Logger:  true,
			Tracer:  true,
			Retrier: true,
		},
		Paths: nil,
	}
}

type Option func(v2 *Config)

func WithAppBuildDate(date string) Option {
	return func(cfg *Config) {
		cfg.App.Date = date
	}
}

func WithAppBuildVersion(v string) Option {
	return func(cfg *Config) {
		cfg.App.Version = v
	}
}

func WithAppBuildGoVersion(v string) Option {
	return func(cfg *Config) {
		cfg.App.GoVersion = v
	}
}

// WithAppBuildDebug TODO maybe useless.
func WithAppBuildDebug(v bool) Option {
	return func(cfg *Config) {
		cfg.App.Debug = v
	}
}

func WithTemplateLogger(in string) Option {
	return func(cfg *Config) {
		cfg.Templates.Logger = Template{Path: in}
	}
}

func WithTemplateTracer(in string) Option {
	return func(cfg *Config) {
		cfg.Templates.Tracer = Template{Path: in}
	}
}

func WithTemplateRetrier(in string) Option {
	return func(cfg *Config) {
		cfg.Templates.Retrier = Template{Path: in}
	}
}

func WithProxyLoggerEnable(in bool) Option {
	return func(cfg *Config) {
		cfg.Proxy.Logger = in
	}
}

func WithProxyTracerEnable(in bool) Option {
	return func(cfg *Config) {
		cfg.Proxy.Tracer = in
	}
}

func WithProxyRetrierEnable(in bool) Option {
	return func(cfg *Config) {
		cfg.Proxy.Retrier = in
	}
}

func WithPath(in string, names []string, outs []string) Option {
	var (
		idx      int
		outPaths []string
		inPaths  = strings.Split(in, "/")
		relative = getRelative(in)
	)

	for idx = range inPaths {
		if inPaths[idx] == relative.Module() {
			break
		}
	}

	switch {
	case len(outs) > 0:
		rootPath := "/" + filepath.Join(inPaths[:idx+1]...)

		for _, out := range outs {
			outPaths = append(outPaths, filepath.Join(rootPath, out))
		}
	default:
		outPaths = append(outPaths, strings.Replace(in, inPaths[len(inPaths)-1], "proxy", 1))
	}

	return func(cfg *Config) {
		cfg.Paths = append(cfg.Paths, Path{
			In:       in,
			Relative: relative.String(),
			Outwards: outPaths,
			Names:    names,
		})
	}
}
