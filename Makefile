include mk/docs.mk
include mk/go.mk

PROJECT_NAME=winkoz-plonk
VERSION=$(shell git rev-parse --short HEAD)
TAG=winkoz/plonk:$(VERSION)
INTERACTIVE?=-it
GO-BIN-FOLDER?=
DOCKER=docker run $(INTERACTIVE) -v $(shell pwd):/go $(TAG)

.PHONY: clean build test ssh docker-build run
# -----------------------------------------------
# Top-level targets

clean:
	rm -rf ./bin

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
