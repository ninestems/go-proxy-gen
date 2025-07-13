package entity

import (
	"strings"
)

// IO represents tag for input/output action, including context.
type IO struct {
	*Common
	alias     string
	name      string
	source    string
	key       string
	parameter *Parameter
}

// NewTagIO builds new instance of IO.
func NewTagIO(
	alias string,
	name string,
	source string,
	key string,
	ttype TagType,
	ptype ProxyType,
) *IO {
	return &IO{
		alias:  alias,
		name:   name,
		source: source,
		key:    key,
		Common: NewCommon(ttype, ptype),
	}
}

// Alias returns alias.
func (t *IO) Alias() string {
	return t.alias
}

// Name returns name.
func (t *IO) Name() string {
	return t.name
}

// Source returns source.
func (t *IO) Source() string {
	return t.source
}

// Key returns key.
func (t *IO) Key() string {
	return t.key
}

// IsEmptyParameter check linked parameter.
func (t *IO) IsEmptyParameter() bool {
	return t.parameter == nil
}

// IsEmptyName check name.
func (t *IO) IsEmptyName() bool {
	return t.name == ""
}

// IsName compare name with in.
func (t *IO) IsName(in string) bool {
	return t.name == in
}

// IsSource compare source with in.
func (t *IO) IsSource(in string) bool {
	return t.source == strings.TrimLeft(in, "*")
}

// IsParentParameter check that parameter is parent for tag.
func (t *IO) IsParentParameter(p *Parameter) bool {
	isSource := t.IsSource(p.Source())               // source from parameter equals source from tag.
	isNameEmpty := t.IsEmptyName()                   // name from tag is empty.
	isName := !t.IsEmptyName() && t.IsName(p.Name()) // name is not empty and equals name from tag.
	return isSource && (isNameEmpty || isName)
}

// ApplyParameter checks input Parameter and apply if conditions are true.
func (t *IO) ApplyParameter(p *Parameter) {
	if !(t.IsParentParameter(p)) {
		return
	}

	t.parameter = p   // sets parent parameter
	if t.name == "" { // if name of param empty, means user does not know param alias, set generated name.
		t.name = p.Name()
	}
}

// ContextIO describe tag for context.
type ContextIO struct {
	*IO
}

// NewIOContextTag builds new instance of ContextIO.
func NewIOContextTag(
	alias string,
	name string,
	source string,
	key string,
	ptype ProxyType,
) *ContextIO {
	return &ContextIO{
		IO: NewTagIO(
			alias,
			name,
			source,
			key,
			TagTypeContext,
			ptype,
		),
	}
}

// InputIO describe input tags.
type InputIO struct {
	*IO
}

// NewIOInputTag builds new instance of InputIO.
func NewIOInputTag(
	alias string,
	name string,
	source string,
	key string,
	ptype ProxyType,
) *InputIO {
	return &InputIO{
		IO: NewTagIO(
			alias,
			name,
			source,
			key,
			TagTypeInput,
			ptype,
		),
	}
}

// OutputIO describe output tags.
type OutputIO struct {
	*IO
}

// NewIOOutputTag builds new instance of OutputIO.
func NewIOOutputTag(
	alias string,
	name string,
	source string,
	key string,
	ptype ProxyType,
) *OutputIO {
	return &OutputIO{
		IO: NewTagIO(
			alias,
			name,
			source,
			key,
			TagTypeContext,
			ptype,
		),
	}
}
