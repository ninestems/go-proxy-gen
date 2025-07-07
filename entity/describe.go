// Package entity contains all models that needed in logic.
package entity

// Package represents list of interface and imports.
type Package struct {
	Name       string
	Imports    []*Import
	Interfaces []*Interface
}

// Interface describes read interface.
type Interface struct {
	Name      string
	Functions []*Function
}

// Import describe all needed imports.
type Import struct {
	Alias  string
	Source string
}

// Function describes function of interface.
type Function struct {
	Name   string
	Input  []*Parameter
	Output []*Parameter
	Tags   []*Tag
}

// Parameter describe input/output function params.
type Parameter struct {
	Name   string
	Source string
}

// Tag describe tags for function.
type Tag struct {
	Type  ProxyType
	Name  string
	Alias string
	Path  string
}
