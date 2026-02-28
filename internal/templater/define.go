package templater

import (
	"fmt"
	"os"
)

func define(proxy, implementation string) string {
	switch proxy {
	case "logger":
		return logger(implementation)
	case "tracer":
		return tracer(implementation)
	case "retrier":
		return retrier(implementation)
	}
	return ""
}

func logger(implementation string) string {
	var template = baseLogger + loggerTemplate + clearTemplate

	switch implementation {
	case "zap":
		template += loggerZapTemplate
	}

	return template
}

func tracer(implementation string) string {
	switch implementation {
	case "opentelemetry":
		return baseTracer + tracerTemplate + clearTemplate + tracerOpenTelemetryTemplate
	default:
		return ""
	}
}

func retrier(implementation string) string {
	switch implementation {
	case "backoff":
		return baseRetrier + retrierTemplate + clearTemplate + retrierBackoffTemplate
	default:
		return ""
	}
}

func read(path string) (string, error) {
	f, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("failed to read template %s: %w", path, err)
	}

	return string(f), err
}
