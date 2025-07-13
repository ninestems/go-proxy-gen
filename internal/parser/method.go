package parser

import (
	"github.com/ninestems/go-proxy-gen/pkg/log"

	"github.com/ninestems/go-proxy-gen/entity"
)

// Parse parses the specified path and returns a list of interfaces
// filtered by optional names (if provided).
func (p *Parser) Parse() (*entity.Package, error) {
	log.Infof("scan file in path '%v': start", p.opt.in)

	pack, err := p.scanner.Scan(p.opt.in, p.opt.ifaces...)
	if err != nil {
		return nil, err
	}

	log.Info("scan file: success")

	log.Debugf("interfaces: %v for package '%v' parsed", pack.Interfaces(), pack.Name())
	log.Debugf("set relative path='%v' for package '%v' parsed", p.opt.relative, pack.Name())

	pack.SetRelative(p.opt.relative)

	log.Debug("prepare markdown data")

	pack.Prepare()

	log.Info("validate markdown: start")

	if err = p.validator.Validate(pack); err != nil {
		return nil, err
	}

	log.Info("validate markdown: success")

	return pack, nil
}
