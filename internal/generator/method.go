package generator

import (
	"log"
)

// Generate read source and build proxy layers.
func (g *Generator) Generate() error {
	log.Printf("starting parsing file and creating markdown\n")
	pack, err := g.parser.Parse()
	if err != nil {
		return err
	}
	log.Printf("parsing file and creating markdown success\n")

	log.Printf("starting defining proxy layer by markdown and template\n")
	err = g.definer.Define(pack)
	if err != nil {
		return err
	}
	log.Printf("defining proxy layer by markdown and template success\n")

	log.Printf("parsing file and creating markdown success\n")

	return nil
}
