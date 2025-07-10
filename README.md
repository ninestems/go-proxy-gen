# go-proxy-gen

## aim
As a web service developer, I often need to inject observability (logging, tracing) into my business logic without modifying it directly.  
The best way to solve this is using the proxy pattern — but handling many cases or writing implementations manually wastes a ton of time.

This package was created with two goals in mind:
1. Embrace Go's `go generate` approach
2. Provide a clean, declarative configuration format

You just describe what you want — and the tool takes care of the rest.

## roadmap

[tap here](doc/roadmap.md)

## Сonfigure tool as executable

### Flags
- `-in` — path to the file or directory with interfaces to scan
- `-out` — path to the folder where generated proxies will be placed
    - Default: `./proxy`
- `-interface` — comma-separated list of interface names to process
- `-types` — list of proxy types to generate: `log`, `trace`, `retry`
    - Default: `log,trace`

## Сonfigure tool as a code

example configuration for function from interface

```
// Name do something.
// 
// goproxygen: 
//  log ctx::traceID::trace_id
//  log input::NAME::input entity.Interface::Name()
//  log output::FUNCS::input entity.Interface::Functions()
//  log output::error::error.Error()
//  trace input::NAME::input entity.Interface::Name()
//  trace output::FUNCS::input entity.Interface::Functions()
//  trace output::error::error.Error()
func Name(ctx context.Context, input *Input) (*Output, error)
```
1. Text between `goproxygen:` and `func ...` line describes options for the proxy generator.  
2. Always use `//` comments — do not use `/* ... */` blocks.  
3. `goproxygen:` must have one space after `//`.  
4. Each option line must have two spaces after `//`.  

### Option Specification

Each directive describes a logging, tracing, or retry instruction based on function context, inputs, or outputs.

**Format**:  
`type section::label::path::accessor`

#### Fields

| Field      | Description                                                                                                                |
|------------|----------------------------------------------------------------------------------------------------------------------------|
| `type`     | One of: `log`, `trace`, `retry`                                                                                            |
| `section`  | Source of the data: `ctx`, `input`, or `output`                                                                            |
| `label`    | Optional name used in logs/traces (e.g., `traceID`, `error`)                                                               |
| `path`     | Full type path, or a combination with name of the parameter(e.g. `entity.Interface` or `input entity.Interface`) |
| `accessor` | Field or method to extract value (e.g., `Name()`, `Error()`, `Meta.ID`)                                                    |

Notes:  
- `ctx` uses only `label::key` (no path or accessor)
- `input` and `output` support access to nested fields or methods
-  One directive per line

#### Parsing Notes
- each directive defines a single log/trace operation.
- label is used as a key for structured output (e.g., logs, trace spans).
- path supports fallback matching even when parameters are unnamed.
- accessor may be a field (.Field) or method (.Method()).
- designed to be readable, explicit, and compatible with common Go interfaces.

### Examples
- [logger example](doc/logger.md)
- [tracer example](doc/tracer.md)
