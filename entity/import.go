package entity

// Import describe all needed imports.
type Import struct {
	alias  string
	source string
}

// NewImport builds new Import.
func NewImport(alias string, source string) *Import {
	return &Import{
		alias:  alias,
		source: source,
	}
}

// Alias returns the alias of the import.
func (im *Import) Alias() string {
	return im.alias
}

// SetAlias sets the alias of the import.
func (im *Import) SetAlias(alias string) {
	im.alias = alias
}

// Source returns the import source.
func (im *Import) Source() string {
	return im.source
}

// SetSource sets the import source.
func (im *Import) SetSource(source string) {
	im.source = source
}
