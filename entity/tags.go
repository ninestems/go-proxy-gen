package entity

// OwnershipInput описывает входной объект
type ownershipper interface {
	IsForLogger() bool
	IsForTracer() bool
	IsForRetrier() bool
}

// ownership collect mark of having types of tag.
type ownership struct {
	logger  bool
	tracer  bool
	retrier bool
}

// Tags describe all list of tags.
type Tags struct {
	ownership
	context []*ContextIO
	input   []*InputIO
	output  []*OutputIO
	retry   []*Retry
}

// setOwnership set mark to true, if tag is correct type.
func (t *Tags) setOwnership(in ownershipper) {
	switch {
	case in.IsForLogger():
		t.logger = true
	case in.IsForTracer():
		t.tracer = true
	case in.IsForRetrier():
		t.retrier = true
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
func (t *Tags) AddRetry(in ...*Retry) {
	for _, ow := range in {
		t.setOwnership(ow)
	}
	t.retry = append(t.retry, in...)
}

// Context returns list of context tag/
func (t *Tags) Context() []*ContextIO {
	return t.context
}

// ContextLogger returns list of tag for context logger.
func (t *Tags) ContextLogger() []*ContextIO {
	out := make([]*ContextIO, 0, len(t.context))
	for _, tag := range t.context {
		if !tag.IsForLogger() {
			continue
		}
		out = append(out, tag)
	}
	return out
}

// ContextTracer returns list of tag for context tracer.
func (t *Tags) ContextTracer() []*ContextIO {
	out := make([]*ContextIO, 0, len(t.context))
	for _, tag := range t.context {
		if !tag.IsForTracer() {
			continue
		}
		out = append(out, tag)
	}
	return out
}

// Input returns full list of input tags.
func (t *Tags) Input() []*InputIO {
	return t.input
}

// InputLogger returns list of input tags for logger.
func (t *Tags) InputLogger() []*InputIO {
	out := make([]*InputIO, 0, len(t.input))
	for _, tag := range t.input {
		if !tag.IsForLogger() {
			continue
		}
		out = append(out, tag)
	}
	return out
}

// InputTracer returns list of input tags for tracer.
func (t *Tags) InputTracer() []*InputIO {
	out := make([]*InputIO, 0, len(t.input))
	for _, tag := range t.input {
		if !tag.IsForTracer() {
			continue
		}
		out = append(out, tag)
	}
	return out
}

// Output returns full list of output tags.
func (t *Tags) Output() []*OutputIO {
	return t.output
}

// OutputLogger returns list of output tags for logger.
func (t *Tags) OutputLogger() []*OutputIO {
	out := make([]*OutputIO, 0, len(t.output))
	for _, tag := range t.output {
		if !tag.IsForLogger() {
			continue
		}
		out = append(out, tag)
	}
	return out
}

// OutputTracer returns list of output tags for tracer.
func (t *Tags) OutputTracer() []*OutputIO {
	out := make([]*OutputIO, 0, len(t.output))
	for _, tag := range t.output {
		if !tag.IsForTracer() {
			continue
		}
		out = append(out, tag)
	}
	return out
}

// Retry returns list of retry tags.
func (t *Tags) Retry() []*Retry {
	return t.retry
}

// IsHaveTags return true if have one or more tag of input type.
func (t *Tags) IsHaveTags(in ProxyType) bool {
	switch in {
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
