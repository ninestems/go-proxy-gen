package definer

import (
	"github.com/ninestems/go-proxy-gen/pkg/log"

	"github.com/ninestems/go-proxy-gen/entity"
)

// Define reads special markdown and makes proxy layers in `out` place.
func (d *Definer) Define(in *entity.Package) error {
	var (
		templates []*entity.Template
		err       error
	)
	{
		log.Info("generate templates: start")

		templates, err = d.proxier.Build(in)
		if err != nil {
			return err
		}

		log.Info("generate templates: success")
	}
	{
		log.Info("prepare folder: start")

		if err = d.emitter.Prepare(); err != nil {
			return err
		}

		log.Info("prepare folder: success")

		func() {
			if err == nil {
				return
			}
			log.Info("remove folder if error found")
			if err = d.emitter.Prepare(); err != nil {
				log.Fatal(err)
			}
		}()
	}
	{
		log.Infof("write files in folder '%v': start", d.opt.out)

		for _, template := range templates {
			log.Infof("write wrapper with name '%v'", template.Name())

			if err = d.emitter.Write(template.Name(), template.Data()); err != nil {
				return err
			}
		}

		log.Info("write files: success")
	}

	return nil
}
