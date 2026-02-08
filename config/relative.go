package config

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"github.com/ninestems/go-proxy-gen/pkg/log"
)

type Relative string

func (r Relative) String() string {
	return string(r)
}

func (r Relative) Module() string {
	elem := strings.Split(string(r), "/")

	if len(elem) > 0 {
		return elem[len(elem)-1]
	}

	return ""
}

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

// getRelative return relative path from root.
func getRelative(filePath string) Relative {
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
		return Relative(moduleName)
	}

	return Relative(moduleName + "/" + filepath.ToSlash(parentPath))
}
