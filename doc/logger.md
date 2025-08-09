### logger example 
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

type Logger interface {
	// Example is a some method.
	// 
	// goproxygen: 
	//  log ctx::trace_id::traceID
	//  log input::input:Input::F::some_input
	//  log output::Output::D::some_output
	//  log output::error::Error()
	Example(ctx context.Context, input *Input) (*Output, error)
}

```