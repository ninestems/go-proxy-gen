package generator

import (
	"github.com/ninestems/go-proxy-gen/pkg/log"
)

// Generate read source and build proxy layers.
func (g *Generator) Generate() error {
	log.Infof("generate: start")
	ifas, err := g.parser.Parse()
	if err != nil {
		return err
	}

	if err = g.definer.Define(ifas); err != nil {
		return err
	}

	log.Infof("generate: success")
	return nil
}
