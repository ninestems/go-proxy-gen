// Package definer describe how app generates proxy layers,
package definer

import (
	"log"

	"go-proxy-gen/entity"
	"go-proxy-gen/internal"
)

var _ internal.DefinerI = (*Definer)(nil)

// Definer generates proxy layer files based on special markdown annotations.
type Definer struct {
	opt     *options
	proxier internal.ProxierI
	emitter internal.EmitterI
}

// New builds new instance of Definer.
func New(opts ...Option) *Definer {
	var cfg options
	for _, opt := range opts {
		opt(&cfg)
	}
	return &Definer{
		opt:     &cfg,
		proxier: cfg.proxier,
		emitter: cfg.emitter,
	}
}

// Define reads special markdown and makes proxy layers in `out` place.
func (d *Definer) Define(in *entity.Package) error {
	log.Printf("generate logger bytes start %s\n", in.Name())
	lbytes, err := d.proxier.DefineLogger(in)
	if err != nil {
		return err
	}
	log.Printf("generate logger bytes success\n")

	log.Printf("creating logger proxy file %s\n", in.Name())
	if err = d.emitter.Write("logger", lbytes); err != nil {
		return err
	}

	log.Printf("creating logger proxy file success %s\n", in.Name())

	return nil
}
