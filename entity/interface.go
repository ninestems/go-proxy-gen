package entity

// Interface describes read interface.
type Interface struct {
	name      string
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
