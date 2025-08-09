package entity

// Parameter describe input/output function params.
type Parameter struct {
	typ    ParameterType
	name   string
	source string
}

// NewInputParameter build input type parameter.
func NewInputParameter(name, source string) *Parameter {
	return &Parameter{
		typ:    ParameterTypeInput,
		name:   name,
		source: source,
	}
}

// NewOutputParameter build input type parameter.
func NewOutputParameter(name, source string) *Parameter {
	return &Parameter{
		typ:    ParameterTypeOutput,
		name:   name,
		source: source,
	}
}

// Type returns parameter type.
func (p *Parameter) Type() ParameterType {
	return p.typ
}

// Name returns the parameter name.
func (p *Parameter) Name() string {
	return p.name
}

// Source returns the parameter source type.
func (p *Parameter) Source() string {
	return p.source
}

// Prepare prepares parameter for latest generate.
func (p *Parameter) Prepare(idx string) {
	if p.name == "" {
		p.name = p.typ.String() + idx
	}
}
