// Package entity describe domain model of tool.
package entity

// Common describe common fields for tags.
type Common struct {
	ttype TagType
	ptype ProxyType
	vtype ValueType
}

// NewCommon builds new instance of Common.
func NewCommon(
	ttype TagType,
	ptype ProxyType,
	vtype ValueType,
) *Common {
	return &Common{
		ttype: ttype,
		ptype: ptype,
		vtype: vtype,
	}
}

// TType return tag type of tag.
func (c *Common) TType() TagType {
	return c.ttype
}

// PType return parameter type of tag.
func (c *Common) PType() ProxyType {
	return c.ptype
}

// VType return type of contains value.
func (c *Common) VType() ValueType {
	return c.vtype
}

// IsSimpleType return flag true tag use common type of golang types.
func (c *Common) IsSimpleType() bool {
	return c.vtype != ValueTypeStruct && c.vtype != ValueTypeUndefined
}

// IsStructType return flag true tag use struct type.
func (c *Common) IsStructType() bool {
	return c.vtype == ValueTypeStruct
}

// IsForLogger returns true if tag for logger setting.
func (c *Common) IsForLogger() bool {
	return c.ptype == ProxyTypeLogger
}

// IsForTracer returns true if tag for tracer setting.
func (c *Common) IsForTracer() bool {
	return c.ptype == ProxyTypeTracer
}

// IsForRetrier returns true if tag for retrier setting.
func (c *Common) IsForRetrier() bool {
	return c.ptype == ProxyTypeRetrier
}
