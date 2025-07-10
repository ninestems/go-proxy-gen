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
		inPathFlg  string
		outPathFlg string
		ifacesFlg  string
		typesFlg   string
	)

	// default: use GOFILE if not explicitly set
	defaultIn := os.Getenv("GOFILE")
	if defaultIn == "" {
		defaultIn = "." // fallback
	}

	flag.StringVar(&inPathFlg, "in", defaultIn, "Source to source package or file (default from $GOFILE)")
	flag.StringVar(&outPathFlg, "out", "./proxy", "Output directory for generated files")
	flag.StringVar(&ifacesFlg, "interface", "", "Comma-separated list of interface names")
	flag.StringVar(&typesFlg, "typesFlg", "log,trace,retry", "Comma-separated list of proxy typesFlg (log,trace,retry)")

	flag.Parse()

	log.Printf("build version: %s", BuildVersion)
	log.Printf("build date: %s", BuildDate)

	// make paths absolute
	in, err := filepath.Abs(inPathFlg)
	if err != nil {
		log.Fatalf("invalid input path: %v", err)
	}

	out, err := filepath.Abs(outPathFlg)
	if err != nil {
		log.Fatalf("invalid output path: %v", err)
	}

	ifaces := strings.Split(ifacesFlg, ",")
	if len(ifaces) == 1 && ifaces[0] == "" {
		ifaces = []string{}
	}

	types := strings.Split(typesFlg, ",")
	if len(types) == 1 && types[0] == "" {
		types = []string{}
	}

	gen := builder.Build(
		in,
		out,
		ifaces,
		types,
	)

	log.Printf("generate is started")
	if err = gen.Generate(); err != nil {
		log.Fatalf("generate ends with error: %v", err)
	}

	log.Printf("generate is done")
}
