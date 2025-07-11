package emitter

import (
	"os"
	"path/filepath"
)

// Prepare recreate folder to proxy.s
func (e *Emitter) Prepare() error {
	if err := os.RemoveAll(e.path); err != nil {
		return err
	}

	if err := os.MkdirAll(e.path, 0755); err != nil {
		return err
	}

	return nil
}


// Write takes a byte slice and writes it to the specified file path.
func (e *Emitter) Write(name string, file []byte) error {
	if err := os.WriteFile(filepath.Join(e.path, name+".go"), file, 0644); err != nil {
		return err
	}

	return nil
}
