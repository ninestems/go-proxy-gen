package proxier

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/ninestems/go-proxy-gen/entity"
)

var funcMap = template.FuncMap{
	"sub":  func(a, b int) int { return a - b },
	"ge":   func(a, b int) bool { return a >= b },
	"not":  func(b bool) bool { return !b },
	"dict": dict,
	"list": func(vals ...interface{}) []interface{} { return vals },
}

// Build generates Go source code for all proxy layer.
func (p *Proxier) Build(in *entity.Interface) ([]*entity.Template, error) {
	var (
		out []*entity.Template
		err error
	)

	for _, tmpl := range p.opts.templates {
		result := template.Must(
			template.New(fmt.Sprintf(
				"%s_%s_%s_%s",
				tmpl.Proxy(),
				tmpl.Name(),
				tmpl.Implementation(),
				in.Name(),
			)).Funcs(funcMap).Parse(tmpl.Template()),
		)

		var buf bytes.Buffer
		if err = result.Execute(&buf, in); err != nil {
			return nil, err
		}

		out = append(out, entity.NewTemplate(
			in.Path(),
			fmt.Sprintf(
				"%s_%s_%s_%s",
				tmpl.Proxy(),
				tmpl.Name(),
				tmpl.Implementation(),
				strings.ToLower(in.Name()),
			),
			buf.Bytes(),
		))
	}

	return out, nil
}
