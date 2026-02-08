package emitter

import (
	"os"
	"path/filepath"
)

// Prepare recreate folder to proxy.s
func (e *Emitter) Prepare() error {
	for _, path := range e.paths {
		if err := os.RemoveAll(path); err != nil {
			return err
		}

		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}

	return nil
}

// Write takes a byte slice and writes it to the specified file paths.
func (e *Emitter) Write(name string, file []byte) error {
	for _, path := range e.paths {
		if err := os.WriteFile(filepath.Join(path, name+".go"), file, 0644); err != nil {
			return err
		}
	}

	return nil
}
