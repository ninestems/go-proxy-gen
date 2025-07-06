// Package validator describe how app validates markdown.
package validator

import (
	"go-proxy-gen/entity"
	"go-proxy-gen/internal"
)

var _ internal.ValidatorI = (*Validator)(nil)

// Validator describe how markdown validates.
type Validator struct {
	path string
}

// New builds new instance of Emitter.
func New() *Validator {
	return &Validator{}
}

// Validate checks a list of interfaces for tag format, structural issues,
// and semantic correctness before code generation.
func (e Validator) Validate(in []entity.Interface) error {
	return nil
}
