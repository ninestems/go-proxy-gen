package proxier

import (
	"github.com/ninestems/go-proxy-gen/entity"
)

// Build generates Go source code for all proxy layer.
func (p *Proxier) Build(in *entity.Package) ([]*entity.Template, error) {
	var (
		out  []*entity.Template
		tmpl []byte
		err  error
	)

	tmpl, err = p.logger(in)
	if err != nil {
		return nil, err
	}

	out = append(out, entity.NewTemplate("logger", tmpl))

	tmpl, err = p.tracer(in)
	if err != nil {
		return nil, err
	}

	out = append(out, entity.NewTemplate("tracer", tmpl))

	tmpl, err = p.retrier(in)
	if err != nil {
		return nil, err
	}

	out = append(out, entity.NewTemplate("retrier", tmpl))

	return out, nil
}
