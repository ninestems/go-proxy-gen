package entity

type Common struct {
	ttype TagType
	ptype ProxyType
}

func NewCommon(ttype TagType, ptype ProxyType) *Common {
	return &Common{
		ttype: ttype,
		ptype: ptype,
	}
}

func (c *Common) TType() TagType {
	return c.ttype
}

func (c *Common) PType() ProxyType {
	return c.ptype
}

func (c *Common) IsForLogger() bool {
	return c.ptype == ProxyTypeLogger
}

func (c *Common) IsForTracer() bool {
	return c.ptype == ProxyTypeTracer
}

func (c *Common) IsForRetrier() bool {
	return c.ptype == ProxyTypeRetrier
}
