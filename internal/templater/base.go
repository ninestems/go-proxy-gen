package templater

import (
	_ "embed"
)

var (
	//go:embed files/base/base_logger.tmpl
	baseLogger string
	//go:embed files/base/base_tracer.tmpl
	baseTracer string
	//go:embed files/base/base_retrier.tmpl
	baseRetrier string
	//go:embed files/base/logger.tmpl
	loggerTemplate string
	//go:embed files/base/tracer.tmpl
	tracerTemplate string
	//go:embed files/base/retrier.tmpl
	retrierTemplate string
	//go:embed files/base/clear.tmpl
	clearTemplate string
)
