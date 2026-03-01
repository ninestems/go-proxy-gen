package proxier

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/ninestems/go-proxy-gen/pkg/entity"
)

// Build generates Go source code for all proxy layer.
func (p *Proxier) Build(in *entity.Interface) ([]*entity.Template, error) {
	var (
		out  []*entity.Template
		err  error
		name string
	)

	for _, tmpl := range p.opts.templates {
		name = fmt.Sprintf(
			"%s_%s_%s_%s",
			tmpl.Proxy(),
			tmpl.Name(),
			tmpl.Implementation(),
			strings.ToLower(in.Name()),
		)

		result := template.Must(template.New(name).Funcs(funcMap).Parse(tmpl.Template()))

		var buf bytes.Buffer
		if err = result.Execute(&buf, in); err != nil {
			return nil, fmt.Errorf("Execute: %w", err)
		}

		out = append(out, entity.NewTemplate(in.Path(), name, buf.Bytes()))
	}

	return out, nil
}
