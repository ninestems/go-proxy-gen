// Package internal contains all internal app logic.
package internal

import (
	"github.com/ninestems/go-proxy-gen/entity"
)

// ParserI defines interface for parsing Go packages
// and extracting interface declarations from source code.
type ParserI interface {
	// Parse parses the specified path and returns a list of interfaces
	// filtered by optional names (if provided).
	Parse() (*entity.Package, error)
}

// ScannerI defines interface for scanning Go interfaces
// and their methods for embedded proxy tags.
type ScannerI interface {
	// Scan scans the provided path and returns interface descriptions
	// with parsed metadata such as proxy tags.
	Scan(path string, names ...string) (*entity.Package, error)
}

// ValidatorI defines interface for validating correctness of
// parsed interface metadata and tag structure.
type ValidatorI interface {
	// Validate checks a list of interfaces for tag format, structural issues,
	// and semantic correctness before code generation.
	Validate(in *entity.Package) error
}

// DefinerI defines interface for generating proxy implementations
// (e.g., logging, tracing) based on parsed metadata.
type DefinerI interface {
	// Define receives a list of interfaces and output path,
	// then generates proxy wrappers and writes them to disk.
	Define(in *entity.Package) error
}

// ProxierI defines interface for building in-memory proxy code
// for a single interface.
type ProxierI interface {
	// DefineLogger generates Go source code for a proxy logger wrapper
	// for a single interface and returns the code as bytes.
	DefineLogger(in *entity.Package) ([]byte, error)
	// DefineTracer generates Go source code for a proxy tracer wrapper
	// for a single interface and returns the code as bytes.
	DefineTracer(in *entity.Package) ([]byte, error)
}

// EmitterI defines interface for persisting generated code
// to disk or another target.
type EmitterI interface {
	// Prepare recreate folder to proxy.
	Prepare() error
	// Write takes a byte slice and writes it to the specified file path.
	Write(name string, file []byte) error
}

// TemplaterI defines how logic extract template for next generation.
type TemplaterI interface {
	// Template returns template for logger.
	Template() string
}
