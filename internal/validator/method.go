package validator

import (
	"fmt"

	"github.com/ninestems/go-proxy-gen/entity"
)

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
				return fmt.Errorf("interface %s has invalid function: %w", iface.Name(), err)
			}

			if err := validateTags(fn.Tags()); err != nil {
				return fmt.Errorf("function %s has invalid tags: %w", fn.Name(), err)
			}
		}
	}

	return nil
}
