// Package builder helps to build executable struct.
package builder

import (
	"github.com/ninestems/go-proxy-gen/config"
	"github.com/ninestems/go-proxy-gen/pkg/log"

	"github.com/ninestems/go-proxy-gen/internal/definer"
	"github.com/ninestems/go-proxy-gen/internal/emitter"
	"github.com/ninestems/go-proxy-gen/internal/generator"
	"github.com/ninestems/go-proxy-gen/internal/parser"
	"github.com/ninestems/go-proxy-gen/internal/proxier"
	"github.com/ninestems/go-proxy-gen/internal/scanner"
	"github.com/ninestems/go-proxy-gen/internal/templater"
	"github.com/ninestems/go-proxy-gen/internal/validator"
)

// Build assembles components into an executable case
func Build(
	opts ...config.Option,
) *generator.Generator {
	cfg := config.DefaultConfig()
	for _, opt := range opts {
		opt(&cfg)
	}

	log.Info("initializing tool: start")

	for _, path := range cfg.Paths {
		log.Debugf("input path: %v", path.In)
		log.Debugf("output paths list: %v", path.Outwards)
		log.Debugf("relative path: %v", path.Relative)
		log.Debugf("interfaces names list: %v", path.Names)
	}

	prxr := proxier.New(
		proxier.WithLoggerTemplater(templater.NewLogger()),   // TODO 18
		proxier.WithTracerTemplater(templater.NewTracer()),   // TODO 18
		proxier.WithRetrierTemplater(templater.NewRetrier()), // TODO 18
		proxier.WithEnableLoggerTemplater(cfg.Proxy.Logger),
		proxier.WithEnableTracerTemplater(cfg.Proxy.Tracer),
		proxier.WithEnableRetrierTemplater(cfg.Proxy.Retrier),
	)

	var gopts = make([]generator.Option, 0, len(cfg.Paths))
	for _, path := range cfg.Paths {
		pars := parser.New(
			parser.WithInPath(path.In),
			parser.WithScanner(
				scanner.New(
					scanner.WithRelativePath(path.Relative),
					scanner.WithIfaces(path.Names),
				)),
			parser.WithValidator(validator.New()),
		)

		def := definer.New(
			definer.WithProxier(prxr),
			definer.WithEmitter(
				emitter.New(
					emitter.WithPath(path.Outwards...),
				),
			),
		)

		gopts = append(gopts, generator.WithPair(pars, def))
	}

	log.Info("initializing tool: success")

	return generator.New(
		gopts...,
	)
}
