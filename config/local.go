package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// LocalConfig represents the root YAML configuration structure
type LocalConfig struct {
	Version string             `yaml:"version"`
	Debug   bool               `yaml:"debug"`
	Proxy   []LayerDescription `yaml:"proxy"`
}

func (cfg *LocalConfig) Options() []Option {
	opts := make([]Option, 0, len(cfg.Proxy)+3)

	opts = append(opts, WithConfigVersion(cfg.Version))
	opts = append(opts, WithAppBuildDebug(cfg.Debug))

	if len(cfg.Proxy) > 0 {
		opts = append(opts, WithLayerReset())
	}

	for _, layer := range cfg.Proxy {
		opts = append(opts, layer.Option())
	}
	return opts
}

// LayerDescription represents a custom proxy configuration
type LayerDescription struct {
	Name  string `yaml:"name"`
	PType string `yaml:"ptype"`
	IType string `yaml:"itype"`
	Path  string `yaml:"path"`
}

// Option builds extra option for layer description.
func (l *LayerDescription) Option() Option {
	args := make([]string, 0, 4)
	args = append(args, l.Name)
	args = append(args, l.PType)
	args = append(args, l.IType)

	if len(l.Path) > 0 {
		args = append(args, l.Path)
	}

	return WithLayer(args...)
}

// readLocal read and return yaml config if found.
func readLocal(path string) (*LocalConfig, error) {
	if path == "" {
		return nil, nil
	}

	bytes, err := os.ReadFile(path)
	switch {
	case os.IsNotExist(err):
		return nil, nil
	case err != nil:
		return nil, err
	}

	var cfg LocalConfig
	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
