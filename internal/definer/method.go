package definer

import (
	"log"

	"github.com/ninestems/go-proxy-gen/entity"
)

// Define reads special markdown and makes proxy layers in `out` place.
func (d *Definer) Define(in *entity.Package) error {
	log.Printf("generate logger bytes start %s\n", in.Name())
	var (
		lbytes, tbytes []byte
		err            error
	)
	{
		lbytes, err = d.proxier.DefineLogger(in)
		if err != nil {
			return err
		}
		log.Printf("generate logger bytes success\n")

		log.Printf("generate tracer bytes start %s\n", in.Name())
		tbytes, err = d.proxier.DefineTracer(in)
		if err != nil {
			return err
		}

		log.Printf("generate tracer bytes success\n")
	}

	log.Printf("start preparing folder to files\n")
	if err = d.emitter.Prepare(); err != nil {
		return err
	}

	log.Printf("preparing folder to files success\n")

	log.Printf("saving proxy files start\n")
	{
		log.Printf("creating logger proxy file %s\n", in.Name())
		if err = d.emitter.Write("logger", lbytes); err != nil {
			return err
		}
		log.Printf("creating logger proxy file success %s\n", in.Name())

		log.Printf("creating tracer proxy file %s\n", in.Name())
		if err = d.emitter.Write("tracer", tbytes); err != nil {
			return err
		}
		log.Printf("creating tracer proxy file success %s\n", in.Name())
	}
	log.Printf("saving proxy files success\n")

	return nil
}
