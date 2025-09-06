package proxier

import (
	"bytes"
	"text/template"

	"github.com/ninestems/go-proxy-gen/entity"
)

// tracer generates Go source code for a proxy tracer wrapper
// for a single interface and returns the code as bytes.
func (p *Proxier) tracer(in *entity.Package) ([]byte, error) {
	funcMap := template.FuncMap{
		"sub":  func(a, b int) int { return a - b },
		"ge":   func(a, b int) bool { return a >= b },
		"dict": dict,
		"list": func(vals ...interface{}) []interface{} { return vals },
	}

	in.SetLayer(entity.ProxyTypeTracer)

	tmpl := template.Must(template.New("tracer_proxy").Funcs(funcMap).Parse(p.tt.Template()))

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, in); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
