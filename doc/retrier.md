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

type Retrier interface {
	// Example is a some method.
	// 
	// goproxygen: 
	//  retry 1s::60s::1.2::4
	Example(ctx context.Context, input *Input) (*Output, error)
}

```