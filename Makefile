GOLANGCI_LINT_VERSION := v2.10.1

.PHONY: build codecheck test perft bench race cover clean pgo lint

build: codecheck test
	rm -fr ./vendor
	CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/sachista-chess-perft perft/main.go

codecheck: lint
	go fmt ./...
	go fix ./...
	go vet ./...

perft:
	go run ./perft $(DEPTH) $(if $(FEN),"$(FEN)")

bench:
	go test ./chessboard/... -run=^$$ -bench=. -benchmem -benchtime=3x

race:
	go test -race -skip TestPerfT ./...

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

clean:
	rm -fr bin/ coverage.out coverage.html default.pgo

pgo:
	@echo "Collecting CPU profile via benchmark ..."
	go test -bench=BenchmarkPerfT -benchtime=3x -cpuprofile=default.pgo ./chessboard/...
	@echo "Rebuilding with PGO ..."
	CGO_ENABLED=0 go build -pgo=default.pgo -ldflags="-s -w" -o bin/sachista-chess-perft ./perft
	@echo "PGO binary: bin/sachista-chess-perft"

lint:
	docker run -t --rm -v $(shell pwd):/app:cached \
		-v $(shell go env GOCACHE):/cache/go \
		-v $(shell go env GOPATH)/pkg:/go/pkg \
		-e GOCACHE=/cache/go \
		-e GOLANGCI_LINT_CACHE=/cache/go \
		-w /app golangci/golangci-lint:${GOLANGCI_LINT_VERSION} \
		golangci-lint run --config .golangci.yml
