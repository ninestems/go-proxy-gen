// Package builder helps to build executable struct.
package builder

import (
	"log"

	"go-proxy-gen/internal/definer"
	"go-proxy-gen/internal/emitter"
	"go-proxy-gen/internal/generator"
	"go-proxy-gen/internal/parser"
	"go-proxy-gen/internal/proxier"
	"go-proxy-gen/internal/scanner"
	"go-proxy-gen/internal/validator"
)

// Build assembles components into an executable case
func Build(in, out string, ifaces, types []string) *generator.Generator {
	log.Printf("initializing tool")
	log.Printf("input path: %v", in)
	log.Printf("output path: %v", out)
	log.Printf("interfaces list: %v", ifaces)
	log.Printf("proxy layers types: %v", types)

	log.Printf("initializing scanner")
	scnner := scanner.New()

	log.Printf("initializing validator")
	vldtr := validator.New()

	log.Printf("initializing parser")
	pars := parser.New(
		parser.WithInPath(in),
		parser.WithIfaces(ifaces),
		parser.WithScanner(scnner),
		parser.WithValidator(vldtr),
	)

	log.Printf("initializing proxier")
	prxr := proxier.New()

	log.Printf("initializing emitter")
	emtr := emitter.New()

	log.Printf("initializing definer")
	def := definer.New(
		definer.WithProxier(prxr),
		definer.WithEmitter(emtr),
	)

	return generator.New(
		generator.WithParser(pars),
		generator.WithDefiner(def),
	)
}
