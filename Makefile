PROJECT_NAME=winkoz-plonk
VERSION=$(shell git rev-parse --short HEAD)
TAG=winkoz/plonk:$(VERSION)
DOCKER=docker run -it -v ${PWD}:/go $(TAG)
GO=$(DOCKER) go

.PHONY: clean build test ssh
# -----------------------------------------------
# Top-level targets

clean:
	rm -rf ./bin

build: clean
	$(GO) build -ldflags="-s -w" -o bin/plonk cmd/main.go
	@echo "Applications built successfully!"

test: build
	echo "test"

docker-build:
	docker build -t $(TAG) .
	@echo "Docker built successfully!"

ssh:
	$(DOCKER) bash
