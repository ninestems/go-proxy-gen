// Package main describe how app started.
package main

import (
	"flag"
	"log"
	"strings"

	"go-proxy-gen/builder"
)

var (
	// BuildDate need to save date and time of building app.
	BuildDate string
	// BuildVersion need to save tag of building app.
	BuildVersion string
)

func main() {
	log.SetPrefix("[go-proxy-gen] ")
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC | log.Lshortfile)

	var (
		inPath  string
		outPath string
		ifaces  string
		types   string
	)

	flag.StringVar(&inPath, "in", ".", "Path to source package or file")
	flag.StringVar(&outPath, "out", "./proxy", "Output directory for generated files")
	flag.StringVar(&ifaces, "interface", "", "Comma-separated list of interface names")
	flag.StringVar(&types, "types", "log,trace,retry", "Comma-separated list of proxy types (log,trace,retry)")

	flag.Parse()

	log.Printf("build version: %s", BuildVersion)
	log.Printf("build date: %s", BuildDate)

	gen := builder.Build(inPath, outPath, strings.Split(ifaces, ","), strings.Split(types, ","))

	log.Printf("generate is started")
	if err := gen.Generate(); err != nil {
		log.Fatalf("generate ends with error: %v", err)
	}

	log.Printf("generate is done")
}
