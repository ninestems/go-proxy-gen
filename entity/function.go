package entity

import (
	"strconv"
)

// Function describes function of interface.
type Function struct {
	name   string
	input  []*Parameter
	output []*Parameter
	tags   *Tags
}

// NewFunction builds new Function.
func NewFunction(name string, input []*Parameter, output []*Parameter, tags *Tags) *Function {
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

// Input returns the input parameters of the function.
func (f *Function) Input() []*Parameter {
	return f.input
}

// Output returns the output parameters of the function.
func (f *Function) Output() []*Parameter {
	return f.output
}

// Tags returns main object which contains all tags.
func (f *Function) Tags() *Tags {
	return f.tags
}

// LogContextTags returns context tag for logger.
func (f *Function) LogContextTags() []*ContextIO {
	return f.tags.ContextLogger()
}

// LogInputTags returns input tag for logger.
func (f *Function) LogInputTags() []*InputIO {
	return f.tags.InputLogger()
}

// LogOutputTags returns output tag for logger.
func (f *Function) LogOutputTags() []*OutputIO {
	return f.tags.OutputLogger()
}

// TraceContextTags returns context tag for tracer.
func (f *Function) TraceContextTags() []*ContextIO {
	return f.tags.ContextTracer()
}

// TraceInputTags returns input tag for tracer.
func (f *Function) TraceInputTags() []*InputIO {
	return f.tags.InputTracer()
}

// TraceOutputTags returns output tag for tracer.
func (f *Function) TraceOutputTags() []*OutputIO {
	return f.tags.OutputTracer()
}

// RetryTags returns all retry tags.
func (f *Function) RetryTags() []*Retry {
	return f.tags.Retry()
}

// Prepare generate parameter names and make link between parameters and tags.
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
	for _, p := range f.input {
		for _, tag := range f.Tags().Context() {
			tag.ApplyParameter(p)
		}
		for _, tag := range f.Tags().Input() {
			tag.ApplyParameter(p)
		}
	}
	for _, p := range f.output {
		for _, tag := range f.Tags().Output() {
			tag.ApplyParameter(p)
		}
	}
}
