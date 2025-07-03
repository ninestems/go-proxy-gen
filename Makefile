test:
	go test -race ./internal/...

lint:
	golangci-lint run ./... --config .golangci.yml