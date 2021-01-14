# Development helpers

# Allow alternative formatter e.g.: `export GOFMT=gofumports`
GOFMT ?= gofmt
PORT ?= 3000

.PHONY: build
build:
	go build -o ./bin/web ./cmd/web

run:
	PORT=$(PORT) ./bin/web

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

# Vendor the dependencies, useful for testing local version of a package (e.g. GoLive) in Docker
.PHONY: vendor
vendor:
	go mod vendor

# Package the app as a docker container
.PHONY: pkg
pkg:
	@docker build --tag golive-example:local .

# Run the packaged dockerized app
.PHONY: pkgrun
pkgrun:
	docker run --rm -p $(PORT):$(PORT) --env PORT=$(PORT) golive-example:local

.PHONY: push
push:
	go mod vendor
	heroku container:push web -a golive-example
	rm -fr vendor

.PHONY: deploy
deploy:
	heroku container:release web -a golive-example

.PHONY: hlog
hlog:
	heroku logs --tail -a golive-example
