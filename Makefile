GOLANGCI_LINT_VERSION := v2.10.1

build: codecheck test
	rm -fr ./vendor
	CGO_ENABLED=0  go build -ldflags="-s -w" -o bin/sachista-chess-perft perft/main.go

codecheck: lint
	go fmt ./...
	go fix ./...
	go vet ./...

test:
	go test ./...

lint:
	docker run -t --rm -v $(shell pwd):/app:cached \
		-v $(shell go env GOCACHE):/cache/go \
		-v $(shell go env GOPATH)/pkg:/go/pkg \
		-e GOCACHE=/cache/go \
		-e GOLANGCI_LINT_CACHE=/cache/go \
		-e GOPRIVATE=oddin.gg,github.com/oddin-gg \
		-w /app golangci/golangci-lint:${GOLANGCI_LINT_VERSION} \
		golangci-lint run --config .golangci.yml
