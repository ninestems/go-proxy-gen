package validator

import (
	"fmt"

	"github.com/ninestems/go-proxy-gen/entity"
)

// validatePackage validates package info.
func validatePackage(in *entity.Package) error {
	if in.Name() == "" {
		return fmt.Errorf("empty package name")
	}

	return nil
}

// validateInterface validate interface info.
func validateInterface(in *entity.Interface) error {
	if in.Name() == "" {
		return fmt.Errorf("empty interface name")
	}

	if len(in.Functions()) == 0 {
		return fmt.Errorf("empty functions of interface '%v'", in.Name())
	}

	return nil
}

// validateFunction validate function info.
func validateFunction(in *entity.Function) error {
	if in.Name() == "" {
		return fmt.Errorf("empty function name")
	}

	hasContext := false
	for _, parameter := range in.Input() {
		if parameter.Source() == "context.Context" {
			hasContext = true
		}
	}

	if !hasContext {
		return fmt.Errorf("missing context input parameter with name '%v'", in.Name())
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
		return fmt.Errorf("empty context tag alias")
	}

	if in.Key() == "" {
		return fmt.Errorf("empty context tag key parameter")
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
