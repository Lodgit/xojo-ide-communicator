BUILD_TIME ?= $(shell date -u '+%Y-%m-%dT%H:%m:%S')

install:
	@go version
	@go get -v golang.org/x/lint/golint
.PHONY: install

test:
	@go version
	@golint -set_exit_status ./...
	@go vet ./...
	@go test -v -timeout 30s -race -coverprofile=coverage.txt -covermode=atomic ./...
.PHONY: test

build:
	@go version
	@go build -v \
		-ldflags "-s -w -X 'main.version=0.0.0' -X 'main.buildTime=$(BUILD_TIME)'" \
		-a -o bin/goipcc main.go
.PHONY: build
