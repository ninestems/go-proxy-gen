package definer

import (
	"github.com/ninestems/go-proxy-gen/entity"
)

// Define receives a list of interfaces and output path,
// then generates proxy wrappers and writes them to disk.
func (d *Definer) Define(in []*entity.Interface) error {
	var (
		tmplts  []*entity.Template
		ltmplts []*entity.Template
		err     error
	)

	for _, iface := range in {
		ltmplts, err = d.opt.proxier.Build(iface)
		if err != nil {
			return err
		}
		if len(ltmplts) == 0 {
			continue
		}
		tmplts = append(tmplts, ltmplts...)
	}

	if err = d.opt.emitter.Prepare(); err != nil {
		return err
	}

	for _, template := range tmplts {
		if err = d.opt.emitter.Write(template.Name(), template.Data()); err != nil {
			return err
		}
	}

	return nil
}
