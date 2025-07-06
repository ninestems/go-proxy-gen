// Package definer describe how app generates proxy layers,
package definer

import (
	"go-proxy-gen/entity"
	"go-proxy-gen/internal"
)

var _ internal.DefinerI = (*Definer)(nil)

// Definer generates proxy layer files based on special markdown annotations.
type Definer struct {
}

// New builds new instance of Definer.
func New() *Definer {
	return &Definer{}
}

// Define reads special markdown and makes proxy layers in `out` place.
func (d *Definer) Define(out string, in []entity.Interface) error {
	return nil
}
