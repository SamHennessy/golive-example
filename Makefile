# Development helpers

# Allow alternative test runner e.g.: `export GOTEST=gotest`
GOTEST ?= go test
# Allow alternative formatter e.g.: `export GOFMT=gofumports`
GOFMT ?= gofmt

.PHONY: build
build:
	go build -o ./bin/web ./cmd/web

run:
	./bin/web

# Run static code analysis
.PHONY: lint
lint:
	golangci-lint run

# format code in a way the linter will be happy
.PHONY: format
format:
	go mod tidy
	$(GOFMT) -l -w ./

# Update local modules
.PHONY: mod-update
mod-update:
	go get -u ./...
