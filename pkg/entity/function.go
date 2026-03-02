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

// IsHaveTags checks function for having proxy tag.
func (f *Function) IsHaveTags(in ProxyType) bool {
	return f.tags.IsHaveTags(in)
}

// Layer function need for template compile.
//
// Calls for select way to generate clear template for proxy layer.
func (f *Function) Layer(in string) bool {
	return f.tags.IsHaveTags(NewProxyType(in))
}

// IsHaveLoggerTag returns true if function contains logger tag.
func (f *Function) IsHaveLoggerTag() bool {
	return f.tags.IsHaveTags(ProxyTypeLogger)
}

// IsHaveTracerTag returns true if function contains tracer tag.
func (f *Function) IsHaveTracerTag() bool {
	return f.tags.IsHaveTags(ProxyTypeTracer)
}

// IsHaveRetrierTag returns true if function contains retrier tag.
func (f *Function) IsHaveRetrierTag() bool {
	return f.tags.IsHaveTags(ProxyTypeRetrier)
}

// ContextTags function need for template compile.
//
// Return context tags for select proxy type.
func (f *Function) ContextTags(in string) []*ContextIO {
	switch NewProxyType(in) {
	case ProxyTypeLogger:
		return f.tags.ContextLogger()
	case ProxyTypeTracer:
		return f.tags.ContextTracer()
	case ProxyTypeCustom:
		panic("custom proxy type not supported")
	default:
		return nil
	}
}

// LogContextTags returns context tag for logger.
//func (f *Function) LogContextTags() []*ContextIO {
//	return f.tags.ContextLogger()
//}

// InputTags function need for template compile.
//
// Return input tags for select proxy type.
func (f *Function) InputTags(in string) []*InputIO {
	switch NewProxyType(in) {
	case ProxyTypeLogger:
		return f.tags.InputLogger()
	case ProxyTypeTracer:
		return f.tags.InputTracer()
	case ProxyTypeCustom:
		panic("custom proxy type not supported")
	default:
		return nil
	}
}

// OutputTags function need for template compile.
//
// Return output tags for select proxy type.
func (f *Function) OutputTags(in string) []*OutputIO {
	switch NewProxyType(in) {
	case ProxyTypeLogger:
		return f.tags.OutputLogger()
	case ProxyTypeTracer:
		return f.tags.OutputTracer()
	case ProxyTypeCustom:
		panic("custom proxy type not supported")
	default:
		return nil
	}
}

// RetryTag returns all retry tags.
func (f *Function) RetryTag() *Retry {
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

	f.LinkParameters()
}

// LinkParameters links input/output parameters with tags.
func (f *Function) LinkParameters() {
	if f.Tags() == nil {
		return // think about panic or error here
	}
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
