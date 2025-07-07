// Package entity contains all models that needed in logic.
package entity

// Package represents list of interface and imports.
type Package struct {
	name       string
	imports    []*Import
	interfaces []*Interface
}

// Name returns the name of the package.
func (p *Package) Name() string {
	return p.name
}

// SetName sets the name of the package.
func (p *Package) SetName(name string) {
	p.name = name
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

// Interface describes read interface.
type Interface struct {
	name      string
	functions []*Function
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

// Import describe all needed imports.
type Import struct {
	alias  string
	source string
}

// Alias returns the alias of the import.
func (im *Import) Alias() string {
	return im.alias
}

// SetAlias sets the alias of the import.
func (im *Import) SetAlias(alias string) {
	im.alias = alias
}

// Source returns the import source path.
func (im *Import) Source() string {
	return im.source
}

// SetSource sets the import source path.
func (im *Import) SetSource(source string) {
	im.source = source
}

// Function describes function of interface.
type Function struct {
	name   string
	input  []*Parameter
	output []*Parameter
	tags   []*Tag
}

// Name returns the function name.
func (f *Function) Name() string {
	return f.name
}

// SetName sets the function name.
func (f *Function) SetName(name string) {
	f.name = name
}

// Input returns the input parameters of the function.
func (f *Function) Input() []*Parameter {
	return f.input
}

// SetInput sets the input parameters of the function.
func (f *Function) SetInput(params []*Parameter) {
	f.input = params
}

// Output returns the output parameters of the function.
func (f *Function) Output() []*Parameter {
	return f.output
}

// SetOutput sets the output parameters of the function.
func (f *Function) SetOutput(params []*Parameter) {
	f.output = params
}

// Tags returns the tags of the function.
func (f *Function) Tags() []*Tag {
	return f.tags
}

// SetTags sets the tags of the function.
func (f *Function) SetTags(tags []*Tag) {
	f.tags = tags
}

// Parameter describe input/output function params.
type Parameter struct {
	name   string
	source string
}

// Name returns the parameter name.
func (p *Parameter) Name() string {
	return p.name
}

// SetName sets the parameter name.
func (p *Parameter) SetName(name string) {
	p.name = name
}

// Source returns the parameter source type.
func (p *Parameter) Source() string {
	return p.source
}

// SetSource sets the parameter source type.
func (p *Parameter) SetSource(source string) {
	p.source = source
}

// Tag describe tags for function.
type Tag struct {
	typ   ProxyType
	name  string
	alias string
	path  string
}

// Type returns the proxy type of the tag.
func (t *Tag) Type() ProxyType {
	return t.typ
}

// SetType sets the proxy type of the tag.
func (t *Tag) SetType(typ ProxyType) {
	t.typ = typ
}

// Name returns the name of the tag.
func (t *Tag) Name() string {
	return t.name
}

// SetName sets the name of the tag.
func (t *Tag) SetName(name string) {
	t.name = name
}

// Alias returns the alias of the tag.
func (t *Tag) Alias() string {
	return t.alias
}

// SetAlias sets the alias of the tag.
func (t *Tag) SetAlias(alias string) {
	t.alias = alias
}

// Path returns the path of the tag.
func (t *Tag) Path() string {
	return t.path
}

// SetPath sets the path of the tag.
func (t *Tag) SetPath(path string) {
	t.path = path
}
