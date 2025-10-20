package proxier

import (
	"bytes"
	"text/template"

	"github.com/ninestems/go-proxy-gen/entity"
)

// logger generates Go source code for a proxy wrapper
// for a single interface and returns the code as bytes.
func (p *Proxier) logger(in *entity.Interface) ([]byte, error) {
	funcMap := template.FuncMap{
		"sub":  func(a, b int) int { return a - b },
		"ge":   func(a, b int) bool { return a >= b },
		"dict": dict,
		"list": func(vals ...interface{}) []interface{} { return vals },
	}

	tmpl := template.Must(
		template.New("logger_proxy_" + in.Name()).Funcs(funcMap).Parse(p.opts.lt.Template()),
	)

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, in); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
