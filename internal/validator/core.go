package validator

import (
	"fmt"

	"github.com/ninestems/go-proxy-gen/entity"
)

// validatePackage validates package info.
func validatePackage(in *entity.Package) error {
	if in.Name() == "" {
		return entity.ErrEmptyPackageName
	}

	return nil
}

// validateInterface validate interface info.
func validateInterface(in *entity.Interface) error {
	if in.Name() == "" {
		return entity.ErrEmptyInterfaceName
	}

	if len(in.Functions()) == 0 {
		return fmt.Errorf("interface %v: %w", in.Name(), entity.ErrEmptyFunctionName)
	}

	return nil
}

// validateFunction validate function info.
func validateFunction(in *entity.Function) error {
	if in.Name() == "" {
		return entity.ErrEmptyFunctionType
	}

	hasContext := false
	for _, parameter := range in.Input() {
		if parameter.Source() == "context.Context" {
			hasContext = true
		}
	}

	if !hasContext {
		return fmt.Errorf("function %v: %w", in.Name(), entity.ErrEmptyContext)
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
			"alias='%v',source='%v':%w",
			in.Alias(),
			in.Source(),
			entity.ErrInvalidContextTagProxyType,
		)
	}

	if in.Alias() == "" {
		return entity.ErrEmptyContextTagAlias
	}

	if in.Key() == "" {
		return entity.ErrEmptyContextTagKey
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
		return fmt.Errorf(
			"alias='%v',source='%v':%w",
			in.Alias(),
			in.Source(),
			entity.ErrInvalidInputTagProxyType,
		)
	}

	if in.Alias() == "" {
		return entity.ErrEmptyInputTagAlias
	}

	if in.Key() == "" {
		return entity.ErrEmptyInputTagKey
	}

	if in.IsEmptyName() {
		return entity.ErrEmptyInputTagName
	}

	if in.Source() == "" {
		return entity.ErrEmptyInputTagSource
	}

	if in.IsEmptyParameter() {
		return fmt.Errorf(
			"alias='%v',source='%v': %w", in.Alias(),
			in.Source(),
			entity.ErrEmptyInputTagParameter,
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
		return fmt.Errorf(
			"alias='%v',source='%v':%w",
			in.Alias(),
			in.Source(),
			entity.ErrInvalidOutputTagProxyType,
		)
	}

	if in.Alias() == "" {
		return entity.ErrEmptyOutputTagAlias
	}

	if in.Key() == "" {
		return entity.ErrEmptyOutputTagKey
	}

	if in.IsEmptyName() {
		return entity.ErrEmptyOutputTagName
	}

	if in.Source() == "" {
		return entity.ErrEmptyOutputTagSource
	}

	if in.IsEmptyParameter() {
		return fmt.Errorf(
			"alias='%v',source='%v': %w", in.Alias(),
			in.Source(),
			entity.ErrEmptyOutputTagParameter,
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
