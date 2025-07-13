// Package validator describe how app validates markdown.
package validator

import (
	"github.com/ninestems/go-proxy-gen/internal"
	"github.com/ninestems/go-proxy-gen/pkg/log"
)

var _ internal.ValidatorI = (*Validator)(nil)

// Validator describe how markdown validates.
type Validator struct{}

// New builds new instance of Emitter.
func New() *Validator {
	log.Debugf("validator initialized")
	return &Validator{}
}
