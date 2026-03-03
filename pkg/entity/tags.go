package entity

// OwnershipInput описывает входной объект
type ownershipper interface {
	IsForProxy(ProxyType) bool
}

// ownership collect mark of having types of tag.
type ownership struct {
	logger  bool
	tracer  bool
	retrier bool
	custom  bool
}

// Tags describe all list of tags.
type Tags struct {
	ownership
	context []*ContextIO
	input   []*InputIO
	output  []*OutputIO
	retry   *Retry
	custom  []any
}

// setOwnership set mark to true, if tag is correct type.
func (t *Tags) setOwnership(in ownershipper) {
	switch {
	case in.IsForProxy(ProxyTypeLogger):
		t.logger = true
	case in.IsForProxy(ProxyTypeTracer):
		t.tracer = true
	case in.IsForProxy(ProxyTypeRetrier):
		t.retrier = true
	case in.IsForProxy(ProxyTypeCustom):
		panic("custom proxy type is not supported")
	}
}

// AddContext added context tag to inner list.
func (t *Tags) AddContext(in ...*ContextIO) {
	for _, ow := range in {
		t.setOwnership(ow)
	}
	t.context = append(t.context, in...)
}

// AddInput added input tag to inner list.
func (t *Tags) AddInput(in ...*InputIO) {
	for _, ow := range in {
		t.setOwnership(ow)
	}
	t.input = append(t.input, in...)
}

// AddOutput added output tag to inner list.
func (t *Tags) AddOutput(out ...*OutputIO) {
	for _, ow := range out {
		t.setOwnership(ow)
	}
	t.output = append(t.output, out...)
}

// AddRetry added retry tag to inner list.
func (t *Tags) AddRetry(in *Retry) {
	t.setOwnership(in)
	t.retry = in
}

// Context returns list of context tag/
func (t *Tags) Context(ins ...ProxyType) []*ContextIO {
	if len(ins) == 0 {
		return t.context
	}

	typ := ins[0]

	out := make([]*ContextIO, 0, len(t.context))

	for _, tag := range t.context {
		if tag.IsForProxy(typ) {
			out = append(out, tag)
		}
	}

	return out
}

// Input returns full list of input tags.
func (t *Tags) Input(ins ...ProxyType) []*InputIO {
	if len(ins) == 0 {
		return t.input
	}

	typ := ins[0]

	out := make([]*InputIO, 0, len(t.input))

	for _, tag := range t.input {
		if tag.IsForProxy(typ) {
			out = append(out, tag)
		}
	}

	return out
}

// Output returns full list of output tags.
func (t *Tags) Output(ins ...ProxyType) []*OutputIO {
	if len(ins) == 0 {
		return t.output
	}

	typ := ins[0]

	out := make([]*OutputIO, 0, len(t.output))

	for _, tag := range t.output {
		if tag.IsForProxy(typ) {
			out = append(out, tag)
		}
	}

	return out
}

// Retry returns retry tag.
func (t *Tags) Retry() *Retry {
	return t.retry
}

// IsHaveTags return true if have one or more tag of input type.
func (t *Tags) IsHaveTags(in ProxyType) bool {
	switch in {
	case ProxyTypeCustom:
		panic("custom proxy type not supported")
	case ProxyTypeLogger:
		return t.logger
	case ProxyTypeTracer:
		return t.tracer
	case ProxyTypeRetrier:
		return t.retrier
	default:
		return false
	}
}
