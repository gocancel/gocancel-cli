test:
	go test -race -v ./...

cover:
	go test -cover -coverprofile .cover.out .
	go tool cover -html=.cover.out -o coverage.html
	open coverage.html

build:
	go generate ./...
	go build -o gocancel cmd/gocancel/main.go
.PHONY: build

lint:
	golangci-lint run

mocks:
	@scripts/generate-mocks.sh
