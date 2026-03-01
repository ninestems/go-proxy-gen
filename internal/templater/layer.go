package templater

import (
	"fmt"

	"github.com/ninestems/go-proxy-gen/pkg/entity"
)

type Layer struct {
	name           string
	proxy          entity.ProxyType
	implementation entity.ImplementationType
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
		proxy:          entity.NewProxyType(proxy),
		implementation: entity.NewImplementationType(implementation),
		path:           path,
	}
}

func (l *Layer) Name() string {
	return l.name
}

func (l *Layer) Proxy() entity.ProxyType {
	return l.proxy
}

func (l *Layer) Implementation() entity.ImplementationType {
	return l.implementation
}

func (l *Layer) Path() string {
	return l.path
}

// Template returns source for logger.
func (l *Layer) Template() string {
	return l.source
}

// Init define template value of read it from disk.
func (l *Layer) Init() error {
	var err error
	l.source, err = define(l)
	if err != nil {
		return fmt.Errorf("read: %w", err)
	}

	return nil
}
