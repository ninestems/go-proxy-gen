### tracer example
```golang
package main

import (
	"context"
)

type A struct {}

type Input struct {
	F int
}
type Output struct {
	D int
}

// B is a some method.
// 
// goproxygen: 
//  trace ctx::traceID::trace_id
//  trace input::F:input.F
//  trace output::out:out.D
func (a *A) B(ctx context.Context, input *Input) (*Output, error) {
	return nil, nil
}
```