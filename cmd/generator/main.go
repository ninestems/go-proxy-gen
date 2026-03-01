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
	// configPath represents path to local yaml file.
	configPath string
	// showVersion represents standard way to show build version and
	showVersion bool
)

func prepareFlags() {
	// default: use GOFILE if not explicitly set
	defaultIn := os.Getenv("GOFILE")
	if defaultIn == "" {
		defaultIn = "." // fallback
	}

	flag.StringVar(&inPathFlg, "in", defaultIn, "Path to source package or file (default from $GOFILE).")
	flag.StringVar(&outPathFlg, "out", "", "Comma-separated list path to destination package for generated files, required relative path from root of project.") //nolint:lll
	flag.StringVar(&ifacesFlg, "interfaces", "", "Comma-separated list of interface names.")
	flag.StringVar(&logLevel, "log-level", "info", "Set level log to debug (default value info).")
	flag.StringVar(&configPath, "config-file", ".go-proxy-gen.yaml", "Using local config (if exists) for extra control of cli functionality.") //nolint:lll
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

	inputPath, err := prepareInputPath()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := config.Init(
		config.WithLocalConfig(configPath),
		config.WithAppBuildDate(BuildDate),
		config.WithAppBuildVersion(BuildVersion),
		config.WithAppBuildGoVersion(BuildGoVersion),
		config.WithPath(inputPath, prepareInterfaceNames(), prepareOutputPath()),
	)
	if err != nil {
		log.Fatal(err)
	}

	gen, err := builder.Build(cfg)
	if err != nil {
		log.Fatalf("builds generator ends with fail: %v", err)
	}

	if err = gen.Generate(); err != nil {
		log.Fatalf("generate proxy ends with error: %v", err)
	}
}
