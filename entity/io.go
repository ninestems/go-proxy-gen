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

func (t *IO) Alias() string {
	return t.alias
}

func (t *IO) Name() string {
	return t.name
}

func (t *IO) Source() string {
	return t.source
}

func (t *IO) Key() string {
	return t.key
}

func (t *IO) IsEmptyParameter() bool {
	return t.parameter == nil
}

func (t *IO) IsEmptyName() bool {
	return t.name == ""
}

func (t *IO) IsName(in string) bool {
	return t.name == in
}

func (t *IO) IsSource(in string) bool {
	return t.source == strings.TrimLeft(in, "*")
}

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

type ContextIO struct {
	*IO
}

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

type InputIO struct {
	*IO
}

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

type OutputIO struct {
	*IO
}

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
