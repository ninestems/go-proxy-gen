### logger example 
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
//  log ctx::traceID::trace_id
//  log input::F:input.F
//  log output::out:out.D
func (a *A) B(ctx context.Context, input *Input) (*Output, error) {
	return nil, nil
}
```