// Package main describe how app started.
package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
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

	// default: use GOFILE if not explicitly set
	defaultIn := os.Getenv("GOFILE")
	if defaultIn == "" {
		defaultIn = "." // fallback
	}

	flag.StringVar(&inPath, "in", defaultIn, "Path to source package or file (default from $GOFILE)")
	flag.StringVar(&outPath, "out", "./proxy", "Output directory for generated files")
	flag.StringVar(&ifaces, "interface", "", "Comma-separated list of interface names")
	flag.StringVar(&types, "types", "log,trace,retry", "Comma-separated list of proxy types (log,trace,retry)")

	flag.Parse()

	log.Printf("build version: %s", BuildVersion)
	log.Printf("build date: %s", BuildDate)

	// make paths absolute
	inPath, err := filepath.Abs(inPath)
	if err != nil {
		log.Fatalf("invalid input path: %v", err)
	}
	outPath, err = filepath.Abs(outPath)
	if err != nil {
		log.Fatalf("invalid output path: %v", err)
	}

	gen := builder.Build(inPath, outPath, strings.Split(ifaces, ","), strings.Split(types, ","))

	log.Printf("generate is started")
	if err = gen.Generate(); err != nil {
		log.Fatalf("generate ends with error: %v", err)
	}

	log.Printf("generate is done")
}

