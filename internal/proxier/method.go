package proxier

import (
	"strings"

	"github.com/ninestems/go-proxy-gen/entity"
)

// Build generates Go source code for all proxy layer.
func (p *Proxier) Build(in *entity.Interface) ([]*entity.Template, error) {
	var (
		out []*entity.Template
		err error
	)

	if p.opts.enables.logger || in.IsLogger() {
		var tmpl []byte
		tmpl, err = p.logger(in)
		if err != nil {
			return nil, err
		}
		out = append(out, entity.NewTemplate(in.Path(), "logger_"+strings.ToLower(in.Name()), tmpl))
	}

	if p.opts.enables.tracer || in.IsTracer() {
		var tmpl []byte
		tmpl, err = p.tracer(in)
		if err != nil {
			return nil, err
		}
		out = append(out, entity.NewTemplate(in.Path(), "tracer_"+strings.ToLower(in.Name()), tmpl))
	}

	if p.opts.enables.retrier || in.IsRetrier() {
		var tmpl []byte
		tmpl, err = p.retrier(in)
		if err != nil {
			return nil, err
		}
		out = append(out, entity.NewTemplate(in.Path(), "retrier_"+strings.ToLower(in.Name()), tmpl))
	}

	return out, nil
}
