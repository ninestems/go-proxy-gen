package entity

import (
	"strconv"
)

// Function describes function of interface.
type Function struct {
	name   string
	input  []*Parameter
	output []*Parameter
	tags   []*Tag
}

// NewFunction builds new Function.
func NewFunction(name string, input []*Parameter, output []*Parameter, tags []*Tag) *Function {
	return &Function{
		name:   name,
		input:  input,
		output: output,
		tags:   tags,
	}
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

func (f *Function) Prepare() {
	for idx := range f.input {
		f.input[idx].Prepare(strconv.Itoa(idx))
	}

	for idx := range f.output {
		f.output[idx].Prepare(strconv.Itoa(idx))
	}
}

// LinkParameters links input/output parameters with tags.
func (f *Function) LinkParameters() {
	for _, tag := range f.tags {
		switch tag.TagType() {
		case TagTypeInput:
			for _, param := range f.input {
				tag.ApplyParameter(param)
			}
		case TagTypeOutput:
			for _, param := range f.output {
				tag.ApplyParameter(param)
			}
		default:
		}
	}
}
