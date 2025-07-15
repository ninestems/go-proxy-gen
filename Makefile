VERSION=0.1.0-alpha
DATE=$(shell date -u +%Y-%m-%dT%H:%M:%SZ)


test:
	go test -race ./internal/...

lint:
	golangci-lint run ./... --config .golangci.yml

build:
	go build -ldflags "-X 'main.Version=$(VERSION)' -X 'main.BuildDate=$(DATE)'" -o go-proxy-gen ./cmd/generator