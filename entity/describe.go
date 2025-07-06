// Package entity contains all models that needed in logic.
package entity

// Interface describes read interface.
type Interface struct {
	Imports   []string
	Name      string
	Functions []Function
}

// Function describes function of interface.
type Function struct {
	Name      string
	Signature string
	Fields    []Field
}

// Field describe tags for function.
type Field struct {
	Type ProxyType
	Name string
	Path string
}
