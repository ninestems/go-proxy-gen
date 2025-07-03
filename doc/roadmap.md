## Roadmap

### Core Features
- [ ] Read interfaces with tag via `go generate` and logging process
    - Support flags for select input and output folder
    - Support flag for select one and more interfaces for target folder
    - Support flag for select proxy type layer to generate
- [ ] Support generate proxy logging layer via `go generate`
  - Support context tag for logger
  - Support input/output tags for logger
- [ ] Support generation of proxy tracer via `go generate`
  - Support context tag for tracer
  - Support input/output tags for tracer
- [ ] Support retryable proxy layer
- [ ] Allow selection of logger implementation (e.g., Zap)
- [ ] ...

### Code-based Configuration Features
- [ ] `log` tag – for printing values to logs
- [ ] `ctx` tag – for extracting values from `context.Context`
- [ ] `input` param – for referencing input arguments
- [ ] `output` param – for referencing return values
- [ ] Optional `name` parameter – used as log field name
- [ ] `trace` tag – for recording data into tracing spans
- [ ] Optional `name` parameter – used as trace field name
- [ ] `retry` tag – for configuring retries for functions
- [ ] ...


### Supported Dependencies
- [ ] [Zap logger](https://github.com/uber-go/zap)
- [ ] [OpenTelemetry tracing](https://opentelemetry.io/docs/languages/go/)
- [ ] [Backoff retryer](https://github.com/cenkalti/backoff)
- [ ] ...
