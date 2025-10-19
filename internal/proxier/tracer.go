package proxier

import (
	"bytes"
	"text/template"

	"github.com/ninestems/go-proxy-gen/entity"
)

// tracer generates Go source code for a proxy tracer wrapper
// for a single interface and returns the code as bytes.
func (p *Proxier) tracer(in *entity.Interface) ([]byte, error) {
	funcMap := template.FuncMap{
		"sub":  func(a, b int) int { return a - b },
		"ge":   func(a, b int) bool { return a >= b },
		"dict": dict,
		"list": func(vals ...interface{}) []interface{} { return vals },
	}

	tmpl := template.Must(
		template.New("tracer_proxy_" + in.Name()).Funcs(funcMap).Parse(p.opts.tt.Template()),
	)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, in); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
