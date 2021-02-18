# Application build settings
BUILD_NAME = xojo-ide-com
GOMAIN ?= main.go
BUILD_DIR ?= bin
BUILD_VERSION ?= 0.0.0
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Go build settings
GOOS ?= $(shell uname -s | tr A-Z a-z)
GOARCH ?= amd64


# Development tasks

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
	@go test $$(go list ./... | grep -v /examples) \
		-v -timeout 30s -race -coverprofile=coverage.txt -covermode=atomic
.PHONY: test

coverage:
	@bash -c "bash <(curl -s https://codecov.io/bash)"
.PHONY: coverage


# Production tasks
ARCHIVE := false

build:
	@echo "Building release binary for $(GOOS)-$(GOARCH)..."
	@go version
	@env \
		GOOS=$(GOOS) \
		GOARCH=$(GOARCH) \
		CGO_ENABLED=0 \
		GO111MODULE=on \
		go build -v \
			-ldflags "\
				-s -w \
				-X 'main.versionNumber=$(BUILD_VERSION)' \
				-X 'main.buildTime=$(BUILD_TIME)'\
			" \
			-a -o $(BUILD_DIR)/$(BUILD_NAME) $(GOMAIN)
	@echo

	@if [[ "$(ARCHIVE)" = "true" ]]; then \
		env \
			BUILD_NAME=$(BUILD_NAME) \
			BUILD_VERSION=$(BUILD_VERSION) \
			BUILD_DIR=$(BUILD_DIR) \
			make archive GOOS=$(GOOS) GOARCH=$(GOARCH); \
	fi
.PHONY: build

archive:
	@if [[ "$(GOOS)" = "darwin" ]] || [[ "$(GOOS)" = "windows" ]]; then \
		OUT_FILE=$(BUILD_NAME)-$(BUILD_VERSION)-$(GOOS)-$(GOARCH); \
		cd $(BUILD_DIR); \
		echo "Archiving and compressing binary using Zip ($${OUT_FILE}.zip)..."; \
		zip -9 -v "$${OUT_FILE}.zip" "$(BUILD_NAME)"; \
		sha256sum "$${OUT_FILE}.zip" >> sha256sum.txt; \
	else \
		OUT_FILE=$(BUILD_NAME)-$(BUILD_VERSION)-$(GOOS)-$(GOARCH); \
		cd $(BUILD_DIR); \
		echo "Archiving and compressing binary using Tar/Gzip ($${OUT_FILE}.tar.gz)..."; \
		tar -vcf - "$(BUILD_NAME)" | gzip -v9 > "$${OUT_FILE}.tar.gz"; \
		sha256sum "$${OUT_FILE}.tar.gz" >> sha256sum.txt; \
	fi
	@echo
.PHONY: archive

release:
	@rm -rf $(BUILD_DIR)
	@make build GOOS=darwin GOARCH=amd64 ARCHIVE=true
	@make build GOOS=darwin GOARCH=arm64 ARCHIVE=true
	@make build GOOS=linux GOARCH=amd64 ARCHIVE=true
	@make build GOOS=linux GOARCH=arm64 ARCHIVE=true
	@make build GOOS=windows GOARCH=amd64 ARCHIVE=true
	@make build GOOS=windows GOARCH=386 ARCHIVE=true
	@echo "Building release binaries done successfully!"
	@ls -ogh $(BUILD_DIR)
	@echo
	@echo "Verifying checksums of archived files..."
	@cd $(BUILD_DIR) && sha256sum -c sha256sum*
	@echo
	@echo "Release files completed and ready to distribute!"
.PHONY: release
