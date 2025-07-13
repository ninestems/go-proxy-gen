// Package templater describe templates using for generate layers.
package templater

import (
	_ "embed"

	"github.com/ninestems/go-proxy-gen/pkg/log"
)

//go:embed files/logger/zap.tmpl
var loggerZapTemplate string

// Logger describe ways to get string template of logger.
type Logger struct {
	source string
}

// NewLogger builds new instance of Logger.
func NewLogger(source string) *Logger {
	if source == "" {
		source = loggerZapTemplate
	}

	log.Debugf("templater logger initialized")
	return &Logger{source}
}

// Template returns template for logger.
func (l *Logger) Template() string {
	return l.source
}
