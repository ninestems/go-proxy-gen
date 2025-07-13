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

func SetLevel(in string) {
	level = strings.ToUpper(in)
}

func Info(v ...any) {
	if level == "INFO" || level == "DEBUG" {
		info.Println(v...)
	}
}
func Infof(format string, v ...any) {
	if level == "INFO" || level == "DEBUG" {
		info.Printf(format, v...)
	}
}

func Debug(v ...any) {
	if level == "DEBUG" {
		debug.Println(v...)
	}
}

func Debugf(format string, v ...any) {
	if level == "DEBUG" {
		debug.Printf(format, v...)
	}
}

func Error(v ...any) {
	err.Println(v...)
}

func Errorf(format string, v ...any) {
	err.Printf(format, v...)
}

func Fatal(v ...any) {
	err.Fatalln(v...)
}

func Fatalf(format string, v ...any) {
	err.Fatalf(format, v...)
}
