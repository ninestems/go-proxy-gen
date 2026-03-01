package emitter

import (
	"os"
	"path/filepath"

	"github.com/ninestems/go-proxy-gen/pkg/log"
)

// Prepare recreate folder to proxy.s
func (e *Emitter) Prepare() error {
	for _, path := range e.paths {
		entries, err := os.ReadDir(path)
		if err != nil && !os.IsNotExist(err) {
			return err
		}

		for _, entry := range entries {
			log.Debug("remove filename:", entry.Name())

			err = os.Remove(filepath.Join(path, entry.Name()))
			if err != nil {
				return err
			}
		}

		log.Debug("remove directory", path)

		err = os.Remove(path)
		if err != nil && !os.IsNotExist(err) {
			return err
		}

		log.Debug("make directory by path", path)

		if err = os.MkdirAll(path, 0755); err != nil {
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
