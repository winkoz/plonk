PROJECT_NAME=winkoz-plonk
VERSION=$(shell git rev-parse --short HEAD)
TAG=winkoz/plonk:$(VERSION)
INTERACTIVE?=-it
GO-BIN-FOLDER?=
DOCKER=docker run $(INTERACTIVE) -v $(shell pwd):/go $(TAG)

.PHONY: clean build test ssh go-build go-test docker-build run go-build-assets
# -----------------------------------------------
# Top-level targets

clean:
	rm -rf ./bin

go-build-assets:
	GO111MODULE=on $(GO-BIN-FOLDER)go-bindata -prefix "data/" -pkg data -o data/data.go data/...

go-build: clean go-build-assets
	GO111MODULE=on GOOS=darwin GOARCH=amd64 go build -mod=mod -ldflags="-s -w" -o bin/plonk main.go
	@echo "Plonk built successfully!"

go-test: go-build-assets
	GO111MODULE=on $(GO-BIN-FOLDER)gotestsum --junitfile unit-tests.xml --format pkgname-and-test-fails -- -mod=mod -cover ./...
	@echo "Plonk finished testing!"

build: clean docker-build
	$(DOCKER) make go-build 
	@echo "Applications built successfully!"

test: docker-build
	$(DOCKER) make go-test
	@echo "Application tests finished."

docker-build:
	docker build -t $(TAG) .
	@echo "Docker built successfully!"

run: build
	$(DOCKER) /go/bin/plonk

ssh: build
	$(DOCKER) bash
