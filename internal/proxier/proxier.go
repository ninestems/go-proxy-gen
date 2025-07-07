// Package proxier describe how app generate proxy layers by markdown.
package proxier

import (
	"go-proxy-gen/entity"
	"go-proxy-gen/internal"
)

var _ internal.ProxierI = (*Proxier)(nil)

// Proxier describes logic generate of proxy layers
type Proxier struct {
	path string
}

// New builds new instance of Emitter.
func New() *Proxier {
	return &Proxier{}
}

// Define generates Go source code for a proxy wrapper
// for a single interface and returns the code as bytes.
func (p Proxier) Define(in *entity.Interface) ([]byte, error) {
	return nil, nil
}
