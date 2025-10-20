package entity

// CommonParameter describe base filed of parameter.
type CommonParameter struct {
	ptype ParameterType // represents type of function parameter - input or output
	vtype ValueType     // represents base type of parameter - e.x. int or struct.
}

// NewCommonParameter builds new instance of CommonParameter.
func NewCommonParameter(ptype ParameterType, vtype ValueType) *CommonParameter {
	return &CommonParameter{
		ptype: ptype,
		vtype: vtype,
	}
}

// Type return type input or output.
func (p *CommonParameter) Type() ParameterType {
	return p.ptype
}

// ValueType return value type from parameter.
func (p *CommonParameter) ValueType() ValueType {
	return p.vtype
}
