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
	"github.com/ninestems/go-proxy-gen/config"
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
	// outPathFlg represents list of paths to output directory, where need placed generated proxy.
	outPathFlg string
	// ifacesFlg names of interface from file for generating.
	ifacesFlg string
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

	flag.StringVar(&inPathFlg, "in", defaultIn, "Source to source package or file (default from $GOFILE).")
	flag.StringVar(&outPathFlg, "out", "", "Comma-separated list path to destination package for generated files, required relative path from root of project.") //nolint:lll
	flag.StringVar(&ifacesFlg, "interfaces", "", "Comma-separated list of interface names.")
	flag.StringVar(&logLevel, "log-level", "info", "Set level log to debug (default value info).")
	flag.BoolVar(&showVersion, "version", false, "Print version and exit.")

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

func prepareInputPath() (string, error) {
	in, err := filepath.Abs(inPathFlg)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path for -in flag: %w", err)
	}

	return in, nil
}

func prepareOutputPath() []string {
	outs := strings.Split(outPathFlg, ",")
	if len(outs) == 1 && outs[0] == "" {
		outs = []string{}
	}

	return outs
}

func prepareInterfaceNames() []string {
	ifaces := strings.Split(ifacesFlg, ",")
	if len(ifaces) == 1 && ifaces[0] == "" {
		ifaces = []string{}
	}

	return ifaces
}

func main() {
	prepareFlags()

	log.SetLevel(logLevel)

	if show() {
		return // only for version print
	}

	log.Debug("FOR TESTING PURPOSE in path", inPathFlg)
	log.Debug("FOR TESTING PURPOSE out path", outPathFlg)

	inputPath, err := prepareInputPath()
	if err != nil {
		log.Fatal(err)
	}

	outputPaths := prepareOutputPath()

	names := prepareInterfaceNames()

	log.Debug("FOR TESTING PURPOSE input path", inputPath)
	log.Debug("FOR TESTING PURPOSE out paths", outputPaths)
	log.Debug("FOR TESTING PURPOSE names", names)

	gen := builder.Build(
		config.WithAppBuildDate(BuildDate),
		config.WithAppBuildVersion(BuildVersion),
		config.WithAppBuildGoVersion(BuildGoVersion),
		config.WithProxyLoggerEnable(true),
		config.WithProxyTracerEnable(true),
		config.WithProxyRetrierEnable(true),
		config.WithPath(inputPath, names, outputPaths),
	)

	if err = gen.GenerateV2(); err != nil {
		log.Fatalf("generate proxy ends with error: %v", err)
	}
}
