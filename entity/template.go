package entity

// Template describe compile bytes and their name.
type Template struct {
	name string
	data []byte
}

// NewTemplate create new Template.
func NewTemplate(name string, data []byte) *Template {
	return &Template{
		name: name,
		data: data,
	}
}

// Name return name of template.
func (t *Template) Name() string {
	return t.name
}

// Data return compiled bytes.
func (t *Template) Data() []byte {
	return t.data
}
