PROJECT_NAME=winkoz-plonk
VERSION=$(shell git rev-parse --short HEAD)
TAG=winkoz/plonk:$(VERSION)
DOCKER=docker run -it -v ${PWD}:/go $(TAG)

.PHONY: clean build test ssh go-build docker-build run
# -----------------------------------------------
# Top-level targets

clean:
	rm -rf ./bin

go-build: clean
	GO111MODULE=on go build -ldflags="-s -w" -o bin/plonk main.go
	@echo "Plonk built successfully!"

build: clean
	$(DOCKER) make go-build 
	@echo "Applications built successfully!"

test: build
	echo "test"

docker-build:
	docker build -t $(TAG) .
	@echo "Docker built successfully!"

run:
	$(DOCKER) /go/bin/plonk

ssh:
	$(DOCKER) bash
