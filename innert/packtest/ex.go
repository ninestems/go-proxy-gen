package packtest

import (
	"context"

	"go-proxy-gen/entity"
)

//go:generate go run ../../cmd/generator/generator.go

type ExampleGenerate interface {
	// All example of all types.
	//
	// goproxygen:
	//  log ctx::log_traceID::trace_id
	//  log input::log_some_input::in:entity.IO::Key()
	//  log output::log_some_output::entity.IO::Key()
	//  trace ctx::trace_traceID::trace_id
	//  trace input::trace_some_input::in:entity.IO::Key()
	//  trace output::trace_some_output::entity.IO::Key()
	All(ctx context.Context, in *entity.IO) (*entity.IO, error) // All example of all types.

	// NoNames
	//
	// goproxygen:
	//  log ctx::log_traceID::trace_id
	//  log input::log_some_input::entity.IO::Key()
	//  log output::log_some_output::entity.IO::Key()
	//  trace ctx::trace_traceID::trace_id
	//  trace input::trace_some_input::entity.IO::Key()
	//  trace output::trace_some_output::entity.IO::Key()
	NoNames(context.Context, *entity.IO) (*entity.IO, error)

	// NoOut
	//
	// goproxygen:
	//  log ctx::log_traceID::trace_id
	//  log input::log_some_input::entity.IO::Key()
	//  trace ctx::trace_traceID::trace_id
	//  trace input::trace_some_input::entity.IO::Key()
	NoOut(context.Context, *entity.IO)

	// NoErrorOut
	//
	// goproxygen:
	//  log ctx::log_traceID::trace_id
	//  log input::log_some_input::entity.IO::Key()
	//  log output::log_some_output::entity.IO::Key()
	//  trace ctx::trace_traceID::trace_id
	//  trace input::trace_some_input::entity.IO::Key()
	//  trace output::trace_some_output::entity.IO::Key()
	NoErrorOut(context.Context, *entity.IO) *entity.IO

	// OnlyContext
	//
	// goproxygen:
	//  log ctx::log_traceID::trace_id
	//  trace ctx::trace_traceID::trace_id
	OnlyContext(context.Context)
}
