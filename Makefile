PROJECT_NAME=winkoz-plonk
VERSION=$(shell git rev-parse --short HEAD)
TAG=winkoz/plonk:$(VERSION)
DOCKER=docker run -it -v ${PWD}:/go $(TAG)

.PHONY: clean build test ssh go-build go-test docker-build run
# -----------------------------------------------
# Top-level targets

clean:
	rm -rf ./bin

go-build: clean
	GO111MODULE=on go build -ldflags="-s -w" -o bin/plonk main.go
	@echo "Plonk built successfully!"

go-test:
	GO111MODULE=on go test ./...
	@echo "Plonk finished testing!"

build: clean
	$(DOCKER) make go-build 
	@echo "Applications built successfully!"

test: docker-build
	$(DOCKER) make go-test
	echo "Application tests finished."

docker-build:
	docker build -t $(TAG) .
	@echo "Docker built successfully!"

run: docker-build build
	$(DOCKER) /go/bin/plonk

ssh:
	$(DOCKER) bash
