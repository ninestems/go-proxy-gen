// Package entity describe domain model of tool.
package entity

// CommonTag describe common fields for tags.
type CommonTag struct {
	ttype TagType
	ptype ProxyType
	vtype ValueType
}

// NewCommonTag builds new instance of CommonTag.
func NewCommonTag(
	ttype TagType,
	ptype ProxyType,
	vtype ValueType,
) *CommonTag {
	return &CommonTag{
		ttype: ttype,
		ptype: ptype,
		vtype: vtype,
	}
}

// TType return tag type of tag.
func (c *CommonTag) TType() TagType {
	return c.ttype
}

// PType return parameter type of tag.
func (c *CommonTag) PType() ProxyType {
	return c.ptype
}

// VType return type of contains value.
func (c *CommonTag) VType() ValueType {
	return c.vtype
}

// IsSimpleType return flag true tag use common type of golang types.
func (c *CommonTag) IsSimpleType() bool {
	return c.vtype != ValueTypeStruct && c.vtype != ValueTypeUndefined
}

// IsStructType return flag true tag use struct type.
func (c *CommonTag) IsStructType() bool {
	return c.vtype == ValueTypeStruct
}

// IsForLogger returns true if tag for logger setting.
func (c *CommonTag) IsForLogger() bool {
	return c.ptype == ProxyTypeLogger
}

// IsForTracer returns true if tag for tracer setting.
func (c *CommonTag) IsForTracer() bool {
	return c.ptype == ProxyTypeTracer
}

// IsForRetrier returns true if tag for retrier setting.
func (c *CommonTag) IsForRetrier() bool {
	return c.ptype == ProxyTypeRetrier
}

func (c *CommonTag) IsForCustom() bool {
	return c.ptype == ProxyTypeCustom
}
