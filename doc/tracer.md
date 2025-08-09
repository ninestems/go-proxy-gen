### tracer example
```golang
package main

import (
	"context"
)

type Input struct {
	F int
}
type Output struct {
	D int
}

type Tracer interface {
	// Example is a some method.
	// 
	// goproxygen:
	//  trace ctx::trace_id::traceID
	//  trace input::input:Input::F::some_input
	//  trace output::Output::D::some_output
	Example(ctx context.Context, input *Input) (*Output, error)
}

```