package entity

import (
	"strings"
)


// Tag describe tags for function.
type Tag struct {
	proxyType ProxyType
	tagType   TagType
	alias     string
	path      *Path
	key       string
	parameter *Parameter
}

// NewTag build new Tag.
func NewTag(
	proxyType ProxyType,
	tagType TagType,
	alias string,
	name string,
	path string,
	key string,
) *Tag {
	return &Tag{
		proxyType: proxyType,
		tagType:   tagType,
		alias:     alias,
		path:      NewPath(name, path),
		key:       key,
	}
}

// ProxyType returns the proxy type of the tag.
func (t *Tag) ProxyType() ProxyType {
	return t.proxyType
}

// SetProxyType sets the proxy type of the tag.
func (t *Tag) SetProxyType(typ ProxyType) {
	t.proxyType = typ
}

// TagType returns the type of tag parameter.
func (t *Tag) TagType() TagType {
	return t.tagType
}

// SetTagType sets the type of tag parameter.
func (t *Tag) SetTagType(typ TagType) {
	t.tagType = typ
}

// Alias returns the alias of the tag.
func (t *Tag) Alias() string {
	return t.alias
}

// SetAlias sets the alias of the tag.
func (t *Tag) SetAlias(alias string) {
	t.alias = alias
}

// Path returns the source of the tag.
func (t *Tag) Path() *Path {
	return t.path
}

// SetPath sets the source of the tag.
func (t *Tag) SetPath(path *Path) {
	t.path = path
}

// Key returns the key of the tag.
func (t *Tag) Key() string {
	return t.key
}

// SetKey sets the key of the tag.
func (t *Tag) SetKey(key string) {
	t.key = key
}

// Parameter returns linked Parameter.
func (t *Tag) Parameter() *Parameter {
	return t.parameter
}

// SetParameter sets Parameter.
func (t *Tag) SetParameter(parameter *Parameter) {
	t.parameter = parameter
}

func (t *Tag) Name() string {
	return t.path.name
}

func (t *Tag) Source() string {
	return t.path.source
}

// ApplyParameter checks input Parameter and apply if conditions are true.
func (t *Tag) ApplyParameter(param *Parameter) {
	if !(t.Source() == strings.TrimLeft(param.Source(), "*")) {
		return
	}

	t.SetParameter(param)

	if t.path.Name() == "" {
		t.path.SetName(param.Name())
	}
}
