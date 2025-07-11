## Roadmap

### Core Features
- [ ] Read interfaces with tag via `go generate` and logging process
    - [x] Support flags for select input and output folder
    - [x] Support flag for select one and more interfaces for target folder
    - [ ] Support flag for select proxy type layer to generate
- [x] Support generate proxy logging layer via `go generate`
  - [x] Support context tag for logger
  - [x] Support input tags for logger
  - [x] Support output tags for logger
- [ ] Support generation of proxy tracer via `go generate`
  - [ ] Support context tag for tracer
  - [ ] Support input/output tags for tracer
- [ ] Support detecting parameter type while generating
  - [ ] Runtime definer for logger (zap)
  - [ ] Runtime definer for tracer (open telemetry)
  - [ ] Generate definer for logger (zap)
  - [ ] Generate definer for tracer (open telemetry)
  - [ ] ...
- [ ] Support retryable proxy layer
  - [ ] Change domain model for support different type of tags
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
- [x] [Zap logger](https://github.com/uber-go/zap)
- [ ] [OpenTelemetry tracing](https://opentelemetry.io/docs/languages/go/)
- [ ] [Backoff retryer](https://github.com/cenkalti/backoff)
- [ ] ...
