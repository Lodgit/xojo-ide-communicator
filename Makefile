APP_VERSION ?= 0.0.0
APP_BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
GOOS ?= darwin

install:
	@go version
	@go get -v golang.org/x/lint/golint
.PHONY: install

run:
	@go version
	@go run main.go
.PHONY: run

test:
	@go version
	@golint -set_exit_status ./...
	@go vet ./...
	@go test -v -timeout 30s -race -coverprofile=coverage.txt -covermode=atomic ./...
.PHONY: test

coverage:
	@bash -c "bash <(curl -s https://codecov.io/bash)"
.PHONY: coverage

build:
	@go version
	@env \
		GOOS=$(GOOS) \
		CGO_ENABLED=0 \
		GO111MODULE=on \
			go build -v \
				-ldflags "\
					-s -w \
					-X 'main.versionNumber=$(APP_VERSION)' \
					-X 'main.buildTime=$(APP_BUILD_TIME)'\
				" \
				-a -o bin/xojo-ide-com main.go
	@du -sh bin/xojo-ide-com
.PHONY: build
