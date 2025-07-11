package entity

type Tags []*Tag

func (t Tags) Input() []InputTag {
	var out []InputTag
	for _, tag := range t {
		if tag.TagType() == TagTypeContext || tag.TagType() == TagTypeInput { // todo ref as method
			out = append(out, tag)
		}
	}
	return out
}

func (t Tags) Output() []OutputTag {
	var out []OutputTag
	for _, tag := range t {
		if tag.TagType() == TagTypeOutput { // todo ref as method
			out = append(out, tag)
		}
	}
	return out
}

func (t Tags) LogContext() []LogContextTag {
	var out []LogContextTag
	for _, tag := range t {
		if tag.TagType() == TagTypeContext && // todo ref as method
			tag.ProxyType() == ProxyTypeLogger { // todo ref as method
			out = append(out, tag)
		}
	}
	return out
}

func (t Tags) LogInput() []LogInputTag {
	var out []LogInputTag
	for _, tag := range t {
		if tag.TagType() == TagTypeInput && // todo ref as method
			tag.ProxyType() == ProxyTypeLogger { // todo ref as method
			out = append(out, tag)
		}
	}
	return out
}

func (t Tags) LogOutput() []LogOutputTag {
	var out []LogOutputTag
	for _, tag := range t {
		if tag.TagType() == TagTypeOutput && // todo ref as method
			tag.ProxyType() == ProxyTypeLogger { // todo ref as method
			out = append(out, tag)
		}
	}
	return out
}

func (t Tags) TraceContext() []TraceContextTag {
	var out []TraceContextTag
	for _, tag := range t {
		if tag.TagType() == TagTypeContext && // todo ref as method
			tag.ProxyType() == ProxyTypeTracer { // todo ref as method
			out = append(out, tag)
		}
	}
	return out
}

func (t Tags) TraceInput() []TraceInputTag {
	var out []TraceInputTag
	for _, tag := range t {
		if tag.TagType() == TagTypeInput && // todo ref as method
			tag.ProxyType() == ProxyTypeTracer { // todo ref as method
			out = append(out, tag)
		}
	}
	return out
}

func (t Tags) TraceOutput() []TraceOutputTag {
	var out []TraceOutputTag
	for _, tag := range t {
		if tag.TagType() == TagTypeOutput && // todo ref as method
			tag.ProxyType() == ProxyTypeTracer { // todo ref as method
			out = append(out, tag)
		}
	}
	return out
}

func (t Tags) Retry() []RetryTag {
	var out []RetryTag
	for _, tag := range t {
		if tag.ProxyType() == ProxyTypeRetrier { // todo ref as method
			out = append(out, tag)
		}
	}
	return out
}
