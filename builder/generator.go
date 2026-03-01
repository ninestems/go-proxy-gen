// Package builder helps to build executable struct.
package builder

import (
	"fmt"

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
func Build(cfg *config.Config) (*generator.Generator, error) {
	log.Info("initializing tool: start")

	for _, path := range cfg.Paths {
		log.Debugf("input path: %v", path.In)
		log.Debugf("output paths list: %v", path.Outwards)
		log.Debugf("relative path: %v", path.Relative)
		log.Debugf("interfaces names list: %v", path.Names)
	}

	var opts []proxier.Option
	for _, l := range cfg.Layers {
		templateLayer := templater.NewLayer(l.Name, l.Proxy, l.Implementation, l.Path)

		err := templateLayer.Init()
		if err != nil {
			return nil, fmt.Errorf("failed to initialize layer %q: %w", l.Name, err)
		}

		opts = append(opts, proxier.WithTemplater(templateLayer))
	}

	prxr := proxier.New(
		opts...,
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
	), nil
}
