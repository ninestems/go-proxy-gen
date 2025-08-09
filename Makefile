VERSION=v0.1.0-alpha
DATE=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)


test:
	go test -race -coverprofile=coverage.out ./internal/... ./entity/... ./pkg/...

coverage:
	go tool cover -func=coverage.out | grep total

lint:
	golangci-lint run ./... --config .golangci.yml

build:
	go build -ldflags "-X 'main.BuildVersion=$(VERSION)' -X 'main.BuildDate=$(DATE)'" -o go-proxy-gen ./cmd/generator