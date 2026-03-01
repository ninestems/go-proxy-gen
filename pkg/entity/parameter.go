package entity

import (
	"strings"
)

// Parameter describe input/output function params.
type Parameter struct {
	*CommonParameter          // embedded base fields.
	names            []string // list of parameter names.
	path             string   // path is source of parameter, sets `source.{value}` if struct and no point in path.
	pointer          bool     // mark set true if parameter is pointer.
}

// NewParameter builds new Parameter instance.
func NewParameter(
	ptype ParameterType,
	path string,
	names []string,
	pointer bool,
) *Parameter {
	out := Parameter{
		CommonParameter: NewCommonParameter(ptype, define(path)),
		names:           names,
		path:            path,
		pointer:         pointer,
	}

	if out.ValueType() == ValueTypeStruct && !strings.Contains(path, ".") {
		if strings.Contains(path, "*") {
			out.path = "*source." + strings.TrimLeft(path, "*")
		} else {
			out.path = "source." + path
		}
	}

	return &out
}

// Path return relative of Parameter.
func (p *Parameter) Path() string {
	return p.path
}

// Name return name of Parameter.
func (p *Parameter) Name() string {
	if len(p.names) > 0 {
		return p.names[0]
	}
	return ""
}

// Call prints parameter in template.
func (p *Parameter) Call() string {
	return p.names[0] + " " + p.path
}

// Prepare prepares parameter for latest generate.
func (p *Parameter) Prepare(idx string) {
	switch len(p.names) {
	case 0:
		p.names = append(p.names, p.Type().String()+idx)
	case 1:
		if p.names[0] == "" {
			p.names[0] = p.Type().String() + idx
		}
	}
}

// NewInputParameter build input type parameter.
func NewInputParameter(names []string, path string, pointer bool) *Parameter {
	return NewParameter(ParameterTypeInput, path, names, pointer)
}

// NewOutputParameter build output type parameter.
func NewOutputParameter(names []string, path string, pointer bool) *Parameter {
	return NewParameter(ParameterTypeOutput, path, names, pointer)
}
