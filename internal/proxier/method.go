package proxier

import (
	"bytes"
	"text/template"

	"go-proxy-gen/entity"
)

// DefineLogger generates Go source code for a proxy wrapper
// for a single interface and returns the code as bytes.
func (p *Proxier) DefineLogger(in *entity.Package) ([]byte, error) {
	funcMap := template.FuncMap{
		"sub": func(a, b int) int { return a - b },
		"ge":  func(a, b int) bool { return a >= b },
	}

	tmpl := template.Must(template.New("logger_proxy").Funcs(funcMap).Parse(p.lt.Template()))

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, in); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// DefineTracer generates Go source code for a proxy tracer wrapper
// for a single interface and returns the code as bytes.
func (p *Proxier) DefineTracer(in *entity.Package) ([]byte, error) {
	funcMap := template.FuncMap{
		"sub": func(a, b int) int { return a - b },
		"ge":  func(a, b int) bool { return a >= b },
	}

	tmpl := template.Must(template.New("tracer_proxy").Funcs(funcMap).Parse(p.tt.Template()))

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, in); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
