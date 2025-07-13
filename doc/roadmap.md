# Roadmap

This roadmap reflects the core evolution plan of `go-proxy-gen`. Milestones align with versioning stages.

---

## Core Features

- [x] Parse interface declarations via `go generate`
- [x] Accept CLI flags for:
  - Input/output paths
  - Target interfaces
  - Proxy layer types (log/trace)
- [x] Generate logging proxy layer
- [x] Generate tracing proxy layer
- [ ] Generate retry proxy layer

---

## Configuration and Templates

- [x] Code-based configuration via inline `goproxygen:` comments
- [ ] External file-based config support
- [ ] Template injection per implementation (e.g., zap, opentelemetry)

---

## Dependency Support

- [x] Zap logger
- [x] OpenTelemetry tracer
- [ ] Backoff retryer

---

## Planned

- [ ] Template override support
- [ ] Interface parsing for unnamed parameters
- [ ] Plugin-style design for extensibility
