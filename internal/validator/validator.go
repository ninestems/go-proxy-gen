// Package validator describe how app validates markdown.
package validator

import (
	"errors"

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
func (e *Validator) Validate(in *entity.Package) error {
	if err := validatePackage(in); err != nil {
		return err
	}

	for _, iface := range in.Interfaces() {
		if err := validateInterface(iface); err != nil {
			return err
		}

		for _, fn := range iface.Functions() {
			if err := validateFunction(fn); err != nil {
				return err
			}

			for _, tag := range fn.Tags() {
				if err := validateTag(tag); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// validatePackage validates package info.
func validatePackage(in *entity.Package) error {
	if in.Name() == "" {
		return errors.New("empty package name")
	}

	if len(in.Interfaces()) == 0 {
		return errors.New("empty package interfaces")
	}

	return nil
}

// validateInterface validate interface info.
func validateInterface(in *entity.Interface) error {
	if in.Name() == "" {
		return errors.New("empty interface name")
	}

	if len(in.Functions()) == 0 {
		return errors.New("empty interface functions")
	}

	return nil
}

// validateFunction validate function info.
func validateFunction(in *entity.Function) error {
	if in.Name() == "" {
		return errors.New("empty function name")
	}

	return nil
}

// validateTag validate tag info.
func validateTag(in *entity.Tag) error {
	switch in.TagType() {
	case entity.TagTypeContext:
		return validateCtxTag(in)
	case entity.TagTypeInput, entity.TagTypeOutput:
		return validateIOTag(in)
	case entity.TagTypeUndefined:
		return errors.New("invalid tag type")
	}

	return nil
}

func validateIOTag(in *entity.Tag) error {
	if in.ProxyType() == entity.ProxyTypeUndefined {
		return errors.New("invalid proxy type")
	}

	if in.Alias() == "" {
		return errors.New("empty tag alias")
	}

	if in.Key() == "" {
		return errors.New("empty tag key parameter")
	}

	if in.Path().Name() == "" {
		return errors.New("empty tag name parameter")
	}

	if in.Path().Source() == "" {
		return errors.New("empty tag path")
	}

	if in.Parameter() == nil {
		return errors.New("tag have no linked parameter")
	}

	return nil
}

func validateCtxTag(in *entity.Tag) error {
	if in.ProxyType() == entity.ProxyTypeUndefined {
		return errors.New("invalid proxy type")
	}

	if in.Alias() == "" {
		return errors.New("empty tag alias")
	}

	if in.Key() == "" {
		return errors.New("empty tag key parameter")
	}

	return nil
}
