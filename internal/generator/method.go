package generator

import (
	"fmt"

	"github.com/ninestems/go-proxy-gen/pkg/log"
)

// Generate read source and build proxy layers.
func (g *Generator) Generate() error {
	log.Infof("generate: start")

	for _, p := range g.pairs {
		ifas, err := p.parser.Parse()
		if err != nil {
			return fmt.Errorf("parser error: %w", err)
		}

		if err = p.definer.Define(ifas); err != nil {
			return fmt.Errorf("definer error: %w", err)
		}
	}

	log.Infof("generate: success")

	return nil

}
