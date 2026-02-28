package templater

import (
	"fmt"
)

type Layer struct {
	name           string
	proxy          string
	implementation string
	path           string
	source         string
}

func NewLayer(
	name string,
	proxy string,
	implementation string,
	path string,
) *Layer {
	return &Layer{
		name:           name,
		proxy:          proxy,
		implementation: implementation,
		path:           path,
		//source:         define(proxy, implementation),
	}
}

func (l *Layer) Name() string {
	return l.name
}

func (l *Layer) Proxy() string {
	return l.proxy
}

func (l *Layer) Implementation() string {
	return l.implementation
}

// Template returns source for logger.
func (l *Layer) Template() string {
	return l.source
}

// Init define template value of read it from disk.
func (l *Layer) Init() error {
	if len(l.path) == 0 {
		l.source = define(l.proxy, l.implementation)
		return nil
	}

	var err error

	l.source, err = read(l.path)

	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	return nil
}
