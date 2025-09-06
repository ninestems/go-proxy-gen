package templater

import (
	_ "embed"
)

var (
	//go:embed files/base/base.tmpl
	baseTemplate string
	//go:embed files/base/clear.tmpl
	clearTemplate string
	//go:embed files/base/logger.tmpl
	loggerBaseTemplate string
	//go:embed files/base/tracer.tmpl
	tracerBaseTemplate string
	//go:embed files/base/retrier.tmpl
	retrierBaseTemplate string
	// commonTemplate contains all base templates,
	commonTemplate = baseTemplate + clearTemplate + loggerBaseTemplate + tracerBaseTemplate + retrierBaseTemplate
)
