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
//  log input::F::input.F
//  log output::out::out.D
//  trace input::F::input.F
//  trace output::F::input.D
func Name(ctx context.Context, input *Input) (*Output, error)
```
1. Text between `goproxygen:` and `func ...` line describes options for the proxy generator.  
2. Always use `//` comments — do not use `/* ... */` blocks.  
3. `goproxygen:` must have one space after `//`.  
4. Each option line must have two spaces after `//`.  

### Option specification
- `type ctx::[name::]key` - extract a value from context.Context for logging/tracing.  
- `type input::[name::]path.field` - extract a value from input parameters.
- `type output::[out::]path.field` - extract a value from output values.

`type` can be log/trace/retry.   
`ctx` means the value is taken from context.Context.     
`input` means the value is taken from function input parameters.
`output` means the value is taken from function return values.  
`name` is optional – sets the key name in logs/traces.  
`key` is the context key.  
`path.field` is the full path to the field inside a struct.  

### Examples
- [logger example](doc/logger.md)
- [tracer example](doc/tracer.md)
