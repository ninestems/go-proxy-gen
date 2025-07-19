package entity

import (
	"errors"
)

var (
	// ErrEmptyPackageName is returned when the package name is not set.
	ErrEmptyPackageName = errors.New("empty package name")
)
var (
	// ErrEmptyInterfaceName is returned when the interface name is not specified.
	ErrEmptyInterfaceName = errors.New("empty interface name")

	// ErrEmptyFunctionName is returned when the function name is not specified.
	ErrEmptyFunctionName = errors.New("empty function name")
)

var (
	// ErrEmptyFunctionType is returned when the function type is missing or undefined.
	ErrEmptyFunctionType = errors.New("empty function name")

	// ErrEmptyContext is returned when the context parameter is missing.
	ErrEmptyContext = errors.New("empty context parameter")
)

var (
	// ErrInvalidContextTagProxyType is returned when the context tag contains an unknown proxy type.
	ErrInvalidContextTagProxyType = errors.New("invalid context tag proxy type")

	// ErrEmptyContextTagAlias is returned when the context tag alias is not provided.
	ErrEmptyContextTagAlias = errors.New("empty context tag alias")

	// ErrEmptyContextTagKey is returned when the context tag key is not specified.
	ErrEmptyContextTagKey = errors.New("empty context tag key")
)

var (
	// ErrInvalidInputTagProxyType is returned when the input tag contains an unknown proxy type.
	ErrInvalidInputTagProxyType = errors.New("invalid input tag proxy type")

	// ErrEmptyInputTagAlias is returned when the input tag alias is not provided.
	ErrEmptyInputTagAlias = errors.New("empty input tag alias")

	// ErrEmptyInputTagKey is returned when the input tag key is not specified.
	ErrEmptyInputTagKey = errors.New("empty input tag key")

	// ErrEmptyInputTagName is returned when the input tag name is not provided.
	ErrEmptyInputTagName = errors.New("empty input tag name")

	// ErrEmptyInputTagSource is returned when the input tag source is not specified.
	ErrEmptyInputTagSource = errors.New("empty input tag source")

	// ErrEmptyInputTagParameter is returned when the input tag parameter is not set.
	ErrEmptyInputTagParameter = errors.New("empty input tag parameter")
)

var (
	// ErrInvalidOutputTagProxyType is returned when the output tag contains an unknown proxy type.
	ErrInvalidOutputTagProxyType = errors.New("invalid output tag proxy type")

	// ErrEmptyOutputTagAlias is returned when the output tag alias is not provided.
	ErrEmptyOutputTagAlias = errors.New("empty output tag alias")

	// ErrEmptyOutputTagKey is returned when the output tag key is not specified.
	ErrEmptyOutputTagKey = errors.New("empty output tag key")

	// ErrEmptyOutputTagName is returned when the output tag name is not provided.
	ErrEmptyOutputTagName = errors.New("empty output tag name")

	// ErrEmptyOutputTagSource is returned when the output tag source is not specified.
	ErrEmptyOutputTagSource = errors.New("empty output tag source")

	// ErrEmptyOutputTagParameter is returned when the output tag parameter is not set.
	ErrEmptyOutputTagParameter = errors.New("empty output tag parameter")
)
