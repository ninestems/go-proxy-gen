// Package log represents simple logger wrapper.
package log

import (
	"log"
	"os"
	"strings"
)

var (
	level = "INFO"

	flags = log.Ldate | log.Ltime | log.LUTC | log.Lshortfile

	info  = log.New(os.Stdout, "[GO-PROXY-GEN] [INFO] ", flags)
	debug = log.New(os.Stdout, "[GO-PROXY-GEN] [DEBUG] ", flags)
	err   = log.New(os.Stderr, "[GO-PROXY-GEN] [ERROR] ", flags)
)

// SetLevel sets logger level.
func SetLevel(in string) {
	level = strings.ToUpper(in)
}

// Info prints info log.
func Info(v ...any) {
	if level == "INFO" || level == "DEBUG" {
		info.Println(v...)
	}
}

// Infof prints info log with format string.
func Infof(format string, v ...any) {
	if level == "INFO" || level == "DEBUG" {
		info.Printf(format, v...)
	}
}

// Debug prints debuf log.
func Debug(v ...any) {
	if level == "DEBUG" {
		debug.Println(v...)
	}
}

// Debugf prints debug log with format string.
func Debugf(format string, v ...any) {
	if level == "DEBUG" {
		debug.Printf(format, v...)
	}
}

// Error prints error log.
func Error(v ...any) {
	err.Println(v...)
}

// Errorf prints error log with format string.
func Errorf(format string, v ...any) {
	err.Printf(format, v...)
}

// Fatal prints error and stop app.
func Fatal(v ...any) {
	err.Fatalln(v...)
}

// Fatalf prints error with format string. and stop app.
func Fatalf(format string, v ...any) {
	err.Fatalf(format, v...)
}
