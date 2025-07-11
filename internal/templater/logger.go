// Package templater describe templates using for generate layers.
package templater

import _ "embed"

//go:embed files/logger/zap.tmpl
var loggerZapTemplate string

type Logger struct {
	source string
}

func NewLogger(source string) *Logger {
	if source == "" {
		source = loggerZapTemplate
	}
	return &Logger{source}
}

// Template returns template for logger.
func (l *Logger) Template() string {
	return l.source
}
