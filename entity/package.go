package entity

// Package represents list of interface and imports with extra options.
type Package struct {
	name       string
	relative   string
	imports    []*Import
	interfaces []*Interface
}

// NewPackage builds new Package.
func NewPackage(
	name string,
	relative string,
	imports []*Import,
	interfaces []*Interface,
) *Package {
	return &Package{
		name:       name,
		relative:   relative,
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

// Prepare make params name and links tags with parameters for each function.
func (p *Package) Prepare() {
	for _, i := range p.interfaces {
		i.Prepare()
	}
}
