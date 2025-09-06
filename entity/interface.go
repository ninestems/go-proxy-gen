package entity

// Layer describe selected layer to generate proxy.
type Layer struct {
	logger  bool
	tracer  bool
	retrier bool
}

// Interface describes read interface.
type Interface struct {
	name      string
	layer     Layer
	functions []*Function
}

// NewInterface builds new Interface.
func NewInterface(name string, fns []*Function) *Interface {
	return &Interface{
		name:      name,
		functions: fns,
	}
}

// Name returns the interface name.
func (i *Interface) Name() string {
	return i.name
}

// SetName sets the interface name.
func (i *Interface) SetName(name string) {
	i.name = name
}

// Functions returns the list of functions in the interface.
func (i *Interface) Functions() []*Function {
	return i.functions
}

// SetFunctions sets the list of functions in the interface.
func (i *Interface) SetFunctions(funcs []*Function) {
	i.functions = funcs
}

// Prepare generate parameter names and make link between parameters and tags.
func (i *Interface) Prepare() {
	for _, fn := range i.functions {
		fn.Prepare()
		fn.LinkParameters()
	}
}

// SetLayer select active layer to generate code.
func (i *Interface) SetLayer(in ProxyType) {
	switch in {
	case ProxyTypeLogger:
		i.layer.logger = true
		i.layer.tracer = false
		i.layer.retrier = false
	case ProxyTypeTracer:
		i.layer.logger = false
		i.layer.tracer = true
		i.layer.retrier = false
	case ProxyTypeRetrier:
		i.layer.logger = false
		i.layer.tracer = false
		i.layer.retrier = true
	default:
	}
}

// IsLogger return true if logger layer was sets.
func (i *Interface) IsLogger() bool {
	return i.layer.logger
}

// IsTracer return true if tracer layer was sets.
func (i *Interface) IsTracer() bool {
	return i.layer.tracer
}

// IsRetrier return true if retrier layer was sets.
func (i *Interface) IsRetrier() bool {
	return i.layer.retrier
}
