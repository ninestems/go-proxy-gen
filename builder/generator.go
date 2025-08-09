// Package builder helps to build executable struct.
package builder

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

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

// getModuleName читает имя модуля из файла go.mod по указанному пути.
func getModuleName(goModPath string) string {
	f, err := os.Open(goModPath)
	if err != nil {
		log.Fatalf("failed to open go.mod: %v", err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatalf("failed to close go.mod file: %v", err)
		}
	}()

	scnnr := bufio.NewScanner(f)
	for scnnr.Scan() {
		line := strings.TrimSpace(scnnr.Text())
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module "))
		}
	}

	return ""
}

// findGoModRoot ищет директорию с go.mod, начиная с файла и двигаясь вверх.
func findGoModRoot(startPath string) string {
	dir := filepath.Dir(startPath)
	for {
		modPath := filepath.Join(dir, "go.mod")
		if _, err := os.Stat(modPath); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}
	log.Fatal("go.mod not found")
	return ""
}

// GetImportPathWithoutPackage возвращает модульный путь без имени пакета,
// то есть до родительской директории относительно пакета.
func getRelative(filePath string) string {
	moduleRoot := findGoModRoot(filePath)
	moduleName := getModuleName(filepath.Join(moduleRoot, "go.mod"))

	if moduleName == "" {
		log.Fatalf("failed to find go.mod for module root: %s", moduleRoot)
	}

	fileDir := filepath.Dir(filePath)

	relPath, err := filepath.Rel(moduleRoot, fileDir)
	if err != nil {
		log.Fatalf("failed to get relative path: %v", err)
	}

	parentPath := filepath.Dir(relPath)

	if parentPath == "." {
		return moduleName
	}
	return moduleName + "/" + filepath.ToSlash(parentPath)
}

// Build assembles components into an executable case
func Build(
	in, out string,
	ifaces, types []string,
) *generator.Generator {
	log.Info("initializing tool: start")

	log.Debugf("input path: %v", in)
	log.Debugf("output path: %v", out)
	log.Debugf("interfaces list: %v", ifaces)
	log.Debugf("proxy layers types: %v", types)

	pars := parser.New(
		parser.WithInPath(in),
		parser.WithRelativePath(getRelative(in)),
		parser.WithIfaces(ifaces),
		parser.WithScanner(scanner.New()),
		parser.WithValidator(validator.New()),
	)

	prxr := proxier.New(
		proxier.WithLoggerTemplater(templater.NewLogger("")),
		proxier.WithTracerTemplater(templater.NewTracer("")),
	)

	emtr := emitter.New(
		emitter.WithPath(out),
	)

	def := definer.New(
		definer.WithOutPath(out),
		definer.WithProxier(prxr),
		definer.WithEmitter(emtr),
	)

	log.Info("initializing tool: success")
	return generator.New(
		generator.WithParser(pars),
		generator.WithDefiner(def),
	)
}
