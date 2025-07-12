// Package validator describe how app validates markdown.
package validator

import (
	"errors"
	"fmt"

	"go-proxy-gen/entity"
	"go-proxy-gen/internal"
)

var _ internal.ValidatorI = (*Validator)(nil)

// Validator describe how markdown validates.
type Validator struct{}

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

			if err := validateTags(fn.Tags()); err != nil {
				return err
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

	hasContext := false
	for _, parameter := range in.Input() {
		if parameter.Source() == "context.Context" {
			hasContext = true
		}
	}

	if !hasContext {
		return errors.New("missing context input parameter")
	}

	return nil
}

// validateTag validate tag info.
func validateTags(in *entity.Tags) error {
	if err := validateContextIOTags(in.Context()...); err != nil {
		return err
	}
	if err := validateInputIOTags(in.Input()...); err != nil {
		return err
	}
	if err := validateOutputIOTags(in.Output()...); err != nil {
		return err
	}
	if err := validateRetryTags(in.Retry()...); err != nil {
		return err
	}

	return nil
}

func validateContextIOTags(in ...*entity.ContextIO) error {
	for _, tag := range in {
		if err := validateContextIOTag(tag); err != nil {
			return err
		}
	}

	return nil
}

func validateContextIOTag(in *entity.ContextIO) error {
	if in.PType() == entity.ProxyTypeUndefined {
		return fmt.Errorf(
			"invalid proxy type for context tag: alias='%v',source='%v'",
			in.Alias(),
			in.Source(),
		)
	}

	if in.Alias() == "" {
		return fmt.Errorf("empty tag alias")
	}

	if in.Key() == "" {
		return fmt.Errorf("empty tag key parameter")
	}

	return nil
}

func validateInputIOTags(in ...*entity.InputIO) error {
	for _, tag := range in {
		if err := validateInputIOTag(tag); err != nil {
			return err
		}
	}

	return nil
}

func validateInputIOTag(in *entity.InputIO) error {
	if in.PType() == entity.ProxyTypeUndefined {
		return fmt.Errorf("invalid proxy type for input tag alias='%v',source='%v'",
			in.Alias(),
			in.Source(),
		)
	}

	if in.Alias() == "" {
		return fmt.Errorf("empty input tag alias")
	}

	if in.Key() == "" {
		return fmt.Errorf("empty input tag key parameter")
	}

	if in.IsEmptyName() {
		return fmt.Errorf("empty input tag name parameter")
	}

	if in.Source() == "" {
		return fmt.Errorf("empty input tag path")
	}

	if in.IsEmptyParameter() {
		return fmt.Errorf("input tag have no linked parameter alias='%v',source='%v'",
			in.Alias(),
			in.Source(),
		)
	}

	return nil
}

func validateOutputIOTags(in ...*entity.OutputIO) error {
	for _, tag := range in {
		if err := validateOutputIOTag(tag); err != nil {
			return err
		}
	}

	return nil
}

func validateOutputIOTag(in *entity.OutputIO) error {
	if in.PType() == entity.ProxyTypeUndefined {
		return fmt.Errorf("invalid proxy type output tag")
	}

	if in.Alias() == "" {
		return fmt.Errorf("empty output tag alias")
	}

	if in.Key() == "" {
		return fmt.Errorf("empty output tag key parameter")
	}

	if in.IsEmptyName() {
		return fmt.Errorf("empty output tag name parameter")
	}

	if in.Source() == "" {
		return fmt.Errorf("empty output tag path")
	}

	if in.IsEmptyParameter() {
		return fmt.Errorf("output tag have no linked parameter alias='%v',source='%v'",
			in.Alias(),
			in.Source(),
		)
	}

	return nil
}

func validateRetryTags(in ...*entity.Retry) error {
	for _, tag := range in {
		if err := validateRetryTag(tag); err != nil {
			return err
		}
	}

	return nil
}

func validateRetryTag(_ *entity.Retry) error {
	panic("implement me")
}
