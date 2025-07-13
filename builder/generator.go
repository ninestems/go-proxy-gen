// Package builder helps to build executable struct.
package builder

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

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
	f.Close()
	log.Fatal("module directive not found in go.mod")
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
		parser.WithRelativePath(getRelative(in)),
		parser.WithIfaces(ifaces),
		parser.WithScanner(scnner),
		parser.WithValidator(vldtr),
	)

	log.Printf("initializing proxier")
	prxr := proxier.New(
		proxier.WithLoggerTemplater(templater.NewLogger("")),
		proxier.WithTracerTemplater(templater.NewTracer("")),
		proxier.WithRetrierTemplater(nil),
	)

	log.Printf("initializing emitter")
	emtr := emitter.New(
		emitter.WithPath(out),
	)

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
