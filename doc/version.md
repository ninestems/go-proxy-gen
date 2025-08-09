# Versioning Criteria

This document outlines the requirements for promoting `go-proxy-gen` across release stages.  
The process follows a progressive, quality-focused approach.

---

## Alpha

> First functional and testable version of the tool.

- [x] 40%+ test coverage for core logic (`parser`, `definer`, `generator`)
- [x] CLI flag parsing (`--in`, `--out`, `--interface`, `--types`)
- [x] Proxy generation for logging layer
- [x] Proxy generation for tracing layer
- [x] Clean package structure (internal/cmd/entity/...)
- [x] Documentation includes:
    - Architecture and soft design
    - Roadmap
    - Examples for logger/tracer tags
    - Versioning overview

---

## Beta

> Feature-complete tool with extended functionality and improved coverage.

- [ ] 80%+ test coverage for core logic
- [ ] Local configuration support (file-based settings)
    - For switching logger/tracer implementations
- [ ] Retry proxy layer generation
- [ ] Documentation updated with:
    - Retry configuration examples
    - Template customization explanation

---

## Release Candidate

> Ready for stable release, feature-locked.

- [ ] 100% test coverage for all core logic
- [ ] Templates respond to selected implementation (e.g., Zap/OpenTelemetry)
- [ ] CI/CD support:
    - Linting
    - Unit and integration tests
- [ ] No open critical bugs from previous stages
- [ ] Documentation covers all intended use cases

---

## Release

> Stable, production-ready version.

- [ ] All planned features are implemented and verified
- [ ] All features are documented with examples
- [ ] Tagged as `v1.0.0` using [semantic versioning](https://semver.org)
