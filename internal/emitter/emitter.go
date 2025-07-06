// Package emitter describe how app generate saves proxy data to disk.
package emitter

import (
	"go-proxy-gen/internal"
)

var _ internal.EmitterI = (*Emitter)(nil)

// Emitter describes logic saving bytes for file on disk
type Emitter struct {
	path string
}

// New builds new instance of Emitter.
func New() *Emitter {
	return &Emitter{}
}

// Write takes a byte slice and writes it to the specified file path.
func (e Emitter) Write(file []byte) error {
	return nil
}
