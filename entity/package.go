package entity

// Package represents list of interface and imports with extra options.
type Package struct {
	name       string
	relative   string
	logger     ProxyLogger
	tracer     ProxyTracer
	retrier    ProxyRetrier
	imports    []*Import
	interfaces []*Interface
}

// NewPackage builds new Package.
func NewPackage(
	name string,
	relative string,
	logger ProxyLogger,
	tracer ProxyTracer,
	retrier ProxyRetrier,
	imports []*Import,
	interfaces []*Interface,
) *Package {
	return &Package{
		name:       name,
		relative:   relative,
		logger:     logger,
		tracer:     tracer,
		retrier:    retrier,
		imports:    imports,
		interfaces: interfaces,
	}
}

// Name returns the name of the package.
func (p *Package) Name() string {
	return p.name
}

// SetName sets the name of the package.
func (p *Package) SetName(name string) {
	p.name = name
}

// Relative returns the module name.
func (p *Package) Relative() string {
	return p.relative
}

// SetRelative sets the module name.
func (p *Package) SetRelative(rel string) {
	p.relative = rel
}

func (p *Package) Logger() ProxyLogger {
	return p.logger
}

func (p *Package) SetLogger(logger ProxyLogger) {
	p.logger = logger
}

func (p *Package) Tracer() ProxyTracer {
	return p.tracer
}

func (p *Package) SetTracer(tracer ProxyTracer) {
	p.tracer = tracer
}

func (p *Package) Retrier() ProxyRetrier {
	return p.retrier
}

func (p *Package) SetRetrier(retrier ProxyRetrier) {
	p.retrier = retrier
}

// Imports returns the list of imports.
func (p *Package) Imports() []*Import {
	return p.imports
}

// SetImports sets the list of imports.
func (p *Package) SetImports(imports []*Import) {
	p.imports = imports
}

// Interfaces returns the list of interfaces.
func (p *Package) Interfaces() []*Interface {
	return p.interfaces
}

// SetInterfaces sets the list of interfaces.
func (p *Package) SetInterfaces(interfaces []*Interface) {
	p.interfaces = interfaces
}

func (p *Package) Prepare() {
	for _, i := range p.interfaces {
		i.Prepare()
	}
}
