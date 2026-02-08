package parser

import (
	"github.com/ninestems/go-proxy-gen/pkg/log"

	"github.com/ninestems/go-proxy-gen/entity"
)

// Parse parses the specified path and returns a list of interfaces
// filtered by optional names (if provided).
func (p *Parser) Parse() ([]*entity.Interface, error) {
	log.Infof("scaning files: start")

	ifas, err := p.scanner.Scan(p.opt.in)
	if err != nil {
		return nil, err
	}

	log.Info("scaning files: success")

	log.Info("validate markdown: start")

	if err = p.validator.Validate(ifas); err != nil {
		return nil, err
	}

	log.Info("validate markdown: success")

	return ifas, nil
}
