package parser

import (
	"go-proxy-gen/entity"
)

// Parse parses the specified path and returns a list of interfaces
// filtered by optional names (if provided).
func (p *Parser) Parse() (*entity.Package, error) {
	pack, err := p.scanner.Scan(p.opt.in, p.opt.ifaces...)
	if err != nil {
		return nil, err
	}

	if err = p.validator.Validate(pack); err != nil {
		return nil, err
	}

	return pack, nil
}
