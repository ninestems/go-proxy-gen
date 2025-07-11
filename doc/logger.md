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
	//  log ctx::traceID::trace_id
	//  log input::some_input::input:Input::F
	//  log output::some_output::Output::D
	//  log output::error::Error()
	Example(ctx context.Context, input *Input) (*Output, error)
}

```