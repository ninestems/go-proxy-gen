## Roadmap

### Core Features
- [x] Read interfaces with tag via `go generate` and logging process
    - [x] Support flags for select input and output folder
    - [x] Support flag for select one and more interfaces for target folder
    - [x] Support flag for select proxy type layer to generate
- [x] Support generate proxy logging layer via `go generate`
  - [x] Support context tag for logger
  - [x] Support input tags for logger
  - [x] Support output tags for logger
- [x] Support generation of proxy tracer via `go generate`
  - [x] Support context tag for tracer
  - [x] Support input/output tags for tracer
- [ ] Support detecting parameter type 
  - [x] Runtime definer for logger (zap)
  - [x] Runtime definer for tracer (open telemetry)
  - [ ] Generate definer for logger (zap)
  - [ ] Generate definer for tracer (open telemetry)
  - [ ] ...
- [ ] Support retryable proxy layer
  - [ ] Teach to read tags
  - [ ] Create template
  - [ ] Modify Definer
- [ ] Allow selection of logger implementation (e.g., Zap)
- [ ] ...

### Code-based Configuration Features
- [x] `log` tag – for printing values to logs
- [x] `ctx` tag – for extracting values from `context.Context`
- [x] `input` param – for referencing input arguments
- [x] `output` param – for referencing return values
- [x] Optional `name` parameter – used as log field name
- [x] `trace` tag – for recording data into tracing spans
- [x] Optional `name` parameter – used as trace field name
- [ ] `retry` tag – for configuring retries for functions
- [ ] ...


### Way versions by functionality
- alfa
  - started by `go generate`
  - read flags when running
  - correct read tags from interface for logger and tracer
  - generate log/proxy layers for one implementation
  - have enough deep in documentation
- beta
  - tool have test coverage ~80% for core logic
  - correct read tags from interface for retrier
  - generate retry proxy layer
  - read local config to change deep settings of tool
- release candidate 
  - ...
- release 
  - ...

### Supported Dependencies
- [x] [Zap logger](https://github.com/uber-go/zap)
- [x] [OpenTelemetry tracing](https://opentelemetry.io/docs/languages/go/)
- [ ] [Backoff retryer](https://github.com/cenkalti/backoff)
- [ ] ...
