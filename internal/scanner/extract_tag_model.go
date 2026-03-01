package scanner

import (
	"time"

	"github.com/ninestems/go-proxy-gen/pkg/entity"
)

type tagSpecification struct {
	context []*tagContextSpecification
	input   []*tagInputSpecification
	output  []*tagOutputSpecification
	retry   []*tagRetrySpecification
}

func (s *tagSpecification) Build() *entity.Tags {
	var (
		out entity.Tags
	)

	for _, tag := range s.context {
		out.AddContext(tag.Build())
	}

	for _, tag := range s.input {
		out.AddInput(tag.Build())
	}

	for _, tag := range s.output {
		out.AddOutput(tag.Build())
	}
	for _, tag := range s.retry {
		out.AddRetry(tag.Build())
	}

	return &out
}

type tagContextSpecification struct {
	proxy  string // proxy type
	alias  string // alias for print parameter
	key    string // key to extract data from context
	source string // always sets to context.Context
}

func (s *tagContextSpecification) Build() *entity.ContextIO {
	var ptype entity.ProxyType
	switch s.proxy {
	case "log":
		ptype = entity.ProxyTypeLogger
	case "trace":
		ptype = entity.ProxyTypeTracer
	default:
		ptype = entity.ProxyTypeUndefined
	}

	return entity.NewIOContextTag(s.alias, s.source, s.key, ptype)
}

type tagInputSpecification struct {
	proxy    string // proxy is type layer
	name     string // name of parameter
	source   string // source is import path
	accessor string // accessor is field to call
	alias    string // alias for parameter in select proxy
}

func (s *tagInputSpecification) Build() *entity.InputIO {
	var ptype entity.ProxyType
	switch s.proxy {
	case "log":
		ptype = entity.ProxyTypeLogger
	case "trace":
		ptype = entity.ProxyTypeTracer
	default:
		ptype = entity.ProxyTypeUndefined
	}

	return entity.NewIOInputTag(s.alias, s.name, s.source, s.accessor, ptype)
}

type tagOutputSpecification struct {
	proxy    string // proxy is type layer
	name     string // name of parameter
	source   string // source is import path
	accessor string // accessor is field to call
	alias    string // alias for parameter in select proxy
}

func (s *tagOutputSpecification) Build() *entity.OutputIO {
	var ptype entity.ProxyType
	switch s.proxy {
	case "log":
		ptype = entity.ProxyTypeLogger
	case "trace":
		ptype = entity.ProxyTypeTracer
	default:
		ptype = entity.ProxyTypeUndefined
	}

	return entity.NewIOOutputTag(s.alias, s.name, s.source, s.accessor, ptype)
}

type tagRetrySpecification struct {
	proxy      string        // proxy type.
	start      time.Duration // start interval for retry.
	end        time.Duration // end internal for retry.
	multiplier float32       // multiplier of step.
	attempts   uint64        // attempts count of retry attempts.
}

func (s *tagRetrySpecification) Build() *entity.Retry {
	return entity.NewRetryTag(s.start, s.end, s.multiplier, s.attempts)
}
