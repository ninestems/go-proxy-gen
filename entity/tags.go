package entity

type Tags struct {
	context []*ContextIO
	input   []*InputIO
	output  []*OutputIO
	retry   []*Retry
}

func (t *Tags) AddContext(in ...*ContextIO) {
	t.context = append(t.context, in...)
}

func (t *Tags) AddInput(in ...*InputIO) {
	t.input = append(t.input, in...)
}

func (t *Tags) AddOutput(out ...*OutputIO) {
	t.output = append(t.output, out...)
}

func (t *Tags) AddRetry(in ...*Retry) {
	t.retry = append(t.retry, in...)
}

func (t *Tags) Context() []*ContextIO {
	return t.context
}

func (t *Tags) ContextLogger() []*ContextIO {
	var out []*ContextIO
	for _, tag := range t.context {
		if !tag.IsForLogger() {
			continue
		}
		out = append(out, tag)
	}
	return out
}

func (t *Tags) ContextTracer() []*ContextIO {
	var out []*ContextIO
	for _, tag := range t.context {
		if !tag.IsForTracer() {
			continue
		}
		out = append(out, tag)
	}
	return out
}

func (t *Tags) Input() []*InputIO {
	return t.input
}

func (t *Tags) InputLogger() []*InputIO {
	var out []*InputIO
	for _, tag := range t.input {
		if !tag.IsForLogger() {
			continue
		}
		out = append(out, tag)
	}
	return out
}

func (t *Tags) InputTracer() []*InputIO {
	var out []*InputIO
	for _, tag := range t.input {
		if !tag.IsForTracer() {
			continue
		}
		out = append(out, tag)
	}
	return out
}

func (t *Tags) Output() []*OutputIO {
	return t.output
}

func (t *Tags) OutputLogger() []*OutputIO {
	var out []*OutputIO
	for _, tag := range t.output {
		if !tag.IsForLogger() {
			continue
		}
		out = append(out, tag)
	}
	return out
}

func (t *Tags) OutputTracer() []*OutputIO {
	var out []*OutputIO
	for _, tag := range t.output {
		if !tag.IsForTracer() {
			continue
		}
		out = append(out, tag)
	}
	return out
}

func (t *Tags) Retry() []*Retry {
	return t.retry
}
