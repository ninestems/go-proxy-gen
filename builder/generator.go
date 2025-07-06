// Package builder helps to build executable struct.
package builder

import (
	"log"

	"go-proxy-gen/internal/generator"
)

// Build assembles components into an executable case
func Build(in, out string, ifaces, types []string) *generator.Generator {
	log.Printf("initializing tool")
	log.Printf("input path: %v", in)
	log.Printf("output path: %v", out)
	log.Printf("interfaces list: %v", ifaces)
	log.Printf("proxy layers types: %v", types)

	log.Printf("initializing parser")
	// TODO init parser here

	log.Printf("initializing definer")
	// TODO init definer here

	return generator.New(
		generator.WithParser(nil),
		generator.WithDefiner(nil),
	)
}
