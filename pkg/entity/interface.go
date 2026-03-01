package entity

// Layer describe selected layer to generate proxy.
type Layer struct {
	logger  bool
	tracer  bool
	retrier bool
}

// Interface describes read interface.
type Interface struct {
	name      string      // name of read interface
	packname  string      // packname is name of read package
	relative  string      // relative relative to original interface
	imports   []*Import   // list of imports
	layer     Layer       // marks for layer - will clear later
	functions []*Function // list of functions of readed interfcase.
}

// NewInterface builds new Interface.
//
// Deprecated
func NewInterface(name string, fns []*Function) *Interface {
	return &Interface{
		name:      name,
		functions: fns,
	}
}

// NewInterfaceV2 builds new Interface.
func NewInterfaceV2(
	name string,
	packname string,
	relative string,
	imprts []*Import,
	fns []*Function,
) *Interface {
	var (
		logger  = false
		tracer  = false
		retrier = false
	)

	for _, fn := range fns {
		logger = logger || fn.IsHaveTags(ProxyTypeLogger)
		tracer = tracer || fn.IsHaveTags(ProxyTypeTracer)
		retrier = retrier || fn.IsHaveTags(ProxyTypeRetrier)
	}

	out := Interface{
		name:     name,
		packname: packname,
		relative: relative,
		imports:  imprts,
		layer: Layer{
			logger:  logger,
			tracer:  tracer,
			retrier: retrier,
		},
		functions: fns,
	}

	out.Prepare()

	return &out
}

// Name returns the interface name.
func (i *Interface) Name() string {
	return i.name
}

// Path return relative path to save interface.
func (i *Interface) Path() string {
	return i.relative + "/" + i.packname
}

// Functions returns the list of functions in the interface.
func (i *Interface) Functions() []*Function {
	return i.functions
}

// Imports return all imports need for interface.
func (i *Interface) Imports() []*Import {
	return i.imports
}

// Prepare generate parameter names and make link between parameters and tags.
func (i *Interface) Prepare() {
	for _, fn := range i.functions {
		fn.Prepare()
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
