package entity

// Template describe compile bytes and their name.
type Template struct {
	relative string
	name     string
	data     []byte
}

// NewTemplate create new Template.
func NewTemplate(relative, name string, data []byte) *Template {
	return &Template{
		relative: relative,
		name:     name,
		data:     data,
	}
}

// Path return name of template.
func (t *Template) Path() string {
	return t.relative + "/proxy/" + t.name
}

// Name return name of template.
func (t *Template) Name() string {
	return t.name
}

// Data return compiled bytes.
func (t *Template) Data() []byte {
	return t.data
}
