package entity

// Path describe source variable name and type.
type Path struct {
	name   string
	source string
}

// NewPath builds instance of Path.
func NewPath(name string, path string) *Path {
	return &Path{
		name:   name,
		source: path,
	}
}

func (p *Path) Name() string {
	return p.name
}

func (p *Path) SetName(name string) {
	p.name = name
}

func (p *Path) Source() string {
	return p.source
}

func (p *Path) SetPath(path string) {
	p.source = path
}

