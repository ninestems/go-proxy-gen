package definer

import (
	"github.com/ninestems/go-proxy-gen/pkg/log"

	"github.com/ninestems/go-proxy-gen/entity"
)

// Define reads special markdown and makes proxy layers in `out` place.
func (d *Definer) Define(in *entity.Package) error {
	var (
		lbytes, tbytes []byte
		err            error
	)
	{
		log.Info("generate bytes: start")

		lbytes, err = d.proxier.DefineLogger(in)
		if err != nil {
			return err
		}

		tbytes, err = d.proxier.DefineTracer(in)
		if err != nil {
			return err
		}

		log.Info("generate bytes: success")
	}

	log.Info("prepare folder: start")

	if err = d.emitter.Prepare(); err != nil {
		return err
	}

	func() {
		if err == nil {
			return
		}
		log.Info("remove folder if error found")
		if err = d.emitter.Prepare(); err != nil {
			log.Fatal(err)
		}
	}()

	log.Info("prepare folder: success")

	{
		log.Infof("write files in folder '%v': start", d.opt.out)

		if err = d.emitter.Write("logger", lbytes); err != nil {
			return err
		}

		if err = d.emitter.Write("tracer", tbytes); err != nil {
			return err
		}

		log.Info("write file: success")
	}

	return nil
}
