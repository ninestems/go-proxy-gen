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
	//  trace ctx::traceID::trace_id
	//  trace input::some_input::input:Input::F
	//  trace output::some_output::Output::D
	Example(ctx context.Context, input *Input) (*Output, error)
}

```