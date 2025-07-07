package generator

import (
	"log"
)

// Generate read source and build proxy layers.
func (g *Generator) Generate() error {
	log.Printf("read file, parse tags and validate\n")
	pack, err := g.parser.Parse()
	if err != nil {
		return err
	}
	log.Printf("parsed tags and validate %v\t", pack)

	return nil
}
