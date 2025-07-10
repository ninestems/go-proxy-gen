package parser

import (
	"log"

	"go-proxy-gen/entity"
)

// Parse parses the specified path and returns a list of interfaces
// filtered by optional names (if provided).
func (p *Parser) Parse() (*entity.Package, error) {
	log.Printf("started parse file in %v\n", p.opt.in)

	pack, err := p.scanner.Scan(p.opt.in, p.opt.ifaces...)
	if err != nil {
		return nil, err
	}

	log.Printf("set relative path %v\n", p.opt.relative)

	pack.SetRelative(p.opt.relative)

	log.Printf("preparing package data\n")

	pack.Prepare()

	log.Printf("validate reader package\n")

	if err = p.validator.Validate(pack); err != nil {
		return nil, err
	}

	log.Printf("validate package success\n")

	return pack, nil
}
