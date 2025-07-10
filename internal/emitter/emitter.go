// Package emitter describe how app generate saves proxy data to disk.
package emitter

import (
	"os"
	"path/filepath"

	"go-proxy-gen/internal"
)

var _ internal.EmitterI = (*Emitter)(nil)

// Emitter describes logic saving bytes for file on disk
type Emitter struct {
	path string
}

// New builds new instance of Emitter.
func New(opts ...Option) *Emitter {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}

	return &Emitter{
		path: cfg.path,
	}
}

// Write takes a byte slice and writes it to the specified file path.
func (e Emitter) Write(name string, file []byte) error {
	if err := os.RemoveAll(e.path); err != nil {
		return err
	}

	if err := os.MkdirAll(e.path, 0755); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(e.path, name+".go"), file, 0644); err != nil {
		return err
	}

	return nil
}
