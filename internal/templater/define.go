package templater

import (
	"fmt"
	"os"

	"github.com/ninestems/go-proxy-gen/pkg/entity"
)

func define(l *Layer) (string, error) {
	switch l.Proxy() {
	case entity.ProxyTypeLogger:
		return logger(l)
	case entity.ProxyTypeTracer:
		return tracer(l)
	case entity.ProxyTypeRetrier:
		return retrier(l)
	case entity.ProxyTypeCustom:
		panic("custom implementation not yet supported")
	default:
		return "", nil
	}
}

func logger(l *Layer) (string, error) {
	var template = baseLogger + loggerTemplate + clearTemplate

	switch l.Implementation() {
	case entity.ImplementationTypeZap:
		template += loggerZapTemplate
	case entity.ImplementationTypeCustom:
		readed, err := read(l.Path())
		if err != nil {
			return "", fmt.Errorf("read: %w", err)
		}

		template += readed
	}

	return template, nil
}

func tracer(l *Layer) (string, error) {
	var template = baseTracer + tracerTemplate + clearTemplate

	switch l.Implementation() {
	case entity.ImplementationTypeOpenTelemetry:
		template += tracerOpenTelemetryTemplate
	case entity.ImplementationTypeCustom:
		readed, err := read(l.Path())
		if err != nil {
			return "", fmt.Errorf("read: %w", err)
		}

		template += readed
	}

	return template, nil
}

func retrier(l *Layer) (string, error) {
	var template = baseRetrier + retrierTemplate + clearTemplate

	switch l.Implementation() {
	case entity.ImplementationTypeBackoff:
		template += retrierBackoffTemplate
	case entity.ImplementationTypeCustom:
		readed, err := read(l.Path())
		if err != nil {
			return "", fmt.Errorf("read: %w", err)
		}

		template += readed
	}

	return template, nil
}

func read(path string) (string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read template %s: %w", path, err)
	}

	return string(f), err
}
