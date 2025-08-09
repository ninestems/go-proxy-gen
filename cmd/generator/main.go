// Package main describe how app started.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/ninestems/go-proxy-gen/builder"
	"github.com/ninestems/go-proxy-gen/pkg/log"
)

var (
	// BuildDate need to save date and time of building app.
	BuildDate = time.Now().UTC().Format(time.RFC3339)
	// BuildVersion need to save tag of building app.
	BuildVersion = "UNKNOW"
	// BuildGoVersion need to show version used to build executable file.
	BuildGoVersion = runtime.Version()
)

var (
	// inPathFlag represents path to input files, which need to read.
	inPathFlg string
	// outPathFlg represents path to output directory, where need placed generated proxy.
	outPathFlg string
	// ifacesFlg list of interfaces which need to get a proxy layers.
	ifacesFlg string
	// typesFlg represents list of types proxy, which need to generate.
	typesFlg string
	// logLevel represents log level.
	logLevel string
	// showVersion represents standard way to show build version and
	showVersion bool
)

func prepareFlags() {
	// default: use GOFILE if not explicitly set
	defaultIn := os.Getenv("GOFILE")
	if defaultIn == "" {
		defaultIn = "." // fallback
	}

	flag.StringVar(&inPathFlg, "in", defaultIn, "Source to source package or file (default from $GOFILE)")
	flag.StringVar(&outPathFlg, "out", "./proxy", "Output directory for generated files")
	flag.StringVar(&ifacesFlg, "interfaces", "", "Comma-separated list of interface names")
	flag.StringVar(&typesFlg, "types", "log,trace,retry", "Comma-separated list of proxy typesFlg (log,trace,retry)")
	flag.StringVar(&logLevel, "log-level", "info", "Set level log to debug, default value info")
	flag.BoolVar(&showVersion, "version", false, "Print version and exit")

	flag.Parse()
}

// show displays build flags and exit.
func show() bool {
	if showVersion {
		fmt.Printf("version: %s\n", BuildVersion)
		fmt.Printf("build date utc: %s\n", BuildDate)
		fmt.Printf("go version: %s\n", BuildGoVersion)
	}
	return showVersion
}

// flagsToParameters prepare parameters.
func flagsToParameters() (string, string, []string, []string, error) {
	in, err := filepath.Abs(inPathFlg)
	if err != nil {
		return "", "", nil, nil, fmt.Errorf("failed to resolve absolute path for -in flag: %w", err)
	}

	out, err := filepath.Abs(outPathFlg)
	if err != nil {
		return "", "", nil, nil, fmt.Errorf("failed to resolve absolute path for -out flag: %w", err)
	}

	ifaces := strings.Split(ifacesFlg, ",")
	if len(ifaces) == 1 && ifaces[0] == "" {
		ifaces = []string{}
	}

	types := strings.Split(typesFlg, ",")
	if len(types) == 1 && types[0] == "" {
		types = []string{}
	}

	return in, out, ifaces, types, nil
}

func main() {
	prepareFlags()

	log.SetLevel(logLevel)

	if show() {
		// only for version print
		return
	}

	in, out, ifaces, types, err := flagsToParameters()
	if err != nil {
		log.Fatal(err)
	}

	gen := builder.Build(in, out, ifaces, types)
	if err = gen.Generate(); err != nil {
		log.Fatalf("generate proxy ends with error: %v", err)
	}
}
