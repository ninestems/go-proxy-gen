// Package config contains structs and way to configuration of cli.
package config

import (
	"errors"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

// App describe base info about cli.
type App struct {
	Debug         bool
	Date          string
	Version       string
	GoVersion     string
	ConfigVersion string
	LocalConfig   string
}

// Layer describes a configuration template for custom tmpl file implementations.
type Layer struct {
	Name           string
	Proxy          string
	Implementation string
	Path           string
}

// Path describes source and destination folders.
type Path struct {
	In       string   // path to the source Go code
	Relative string   // relative path
	Outwards []string // paths to location for store generated files
	Names    []string // list names of interfaces to include generation
}

type Config struct {
	IsCorrupted bool
	App         App
	Layers      []Layer
	Paths       []Path
}

func Init(opts ...Option) (*Config, error) {
	cfg := DefaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	if len(cfg.App.LocalConfig) > 0 {
		local, err := readLocal(cfg.App.LocalConfig)
		if err != nil {
			return nil, err
		}

		if local != nil {
			for _, opt := range local.Options() {
				opt(&cfg)
			}
		}
	}

	if cfg.IsCorrupted {
		return nil, errors.New("corrupted config")
	}

	return &cfg, nil
}

func (cfg *Config) ApplyLocal(in *LocalConfig) {
	cfg.App.Debug = in.Debug
}

func DefaultConfig() Config {
	return Config{
		App: App{
			Debug:         false,
			Date:          time.Now().UTC().Format(time.RFC3339),
			Version:       "UNKNOW",
			GoVersion:     runtime.Version(),
			ConfigVersion: "UNKNOW",
		},
		Layers: []Layer{
			{
				Name:           "logger",
				Proxy:          "logger",
				Implementation: "zap",
			},
			{
				Name:           "tracer",
				Proxy:          "tracer",
				Implementation: "opentelemetry",
			},
			{
				Name:           "retrier",
				Proxy:          "retrier",
				Implementation: "backoff",
			},
		},
		Paths: nil,
	}
}

type Option func(cfg *Config)

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

func WithLocalConfig(in string) Option {
	return func(cfg *Config) {
		if len(in) > 0 {
			return
		}
		cfg.App.LocalConfig = in
	}
}

// WithAppBuildDebug
func WithAppBuildDebug(v bool) Option {
	return func(cfg *Config) {
		cfg.App.Debug = v
	}
}

func WithConfigVersion(v string) Option {
	return func(cfg *Config) {
		cfg.App.ConfigVersion = v
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

// WithLayer set layer setting.
//
// Wait for list of arguments `layer name`, `proxy type`, `implementation type` and optional `path` to template.
//
// If list of parameters wrong - corrupting config flag.
func WithLayer(in ...string) Option {
	return func(cfg *Config) {
		switch len(in) {
		case 3:
			cfg.Layers = append(cfg.Layers, Layer{
				Name:           in[0],
				Proxy:          in[1],
				Implementation: in[2],
			})
		case 4:
			cfg.Layers = append(cfg.Layers, Layer{
				Name:           in[0],
				Proxy:          in[1],
				Implementation: in[2],
				Path:           in[3],
			})
		default:
			cfg.IsCorrupted = true
		}
	}
}
