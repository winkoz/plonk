PROJECT_NAME=winkoz-plonk
VERSION=$(shell git rev-parse --short HEAD)
TAG=winkoz/plonk:$(VERSION)
DOCKER=docker run -it -v ${PWD}:/go $(TAG)
ADR_DOCKER=docker run -it -v ${PWD}:/plonk -w /plonk brianskarda/adr-tools-docker:latest

.PHONY: clean build test ssh go-build go-test docker-build run adr_init adr super_adr
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
	@echo "Application tests finished."

docker-build:
	docker build -t $(TAG) .
	@echo "Docker built successfully!"

run: docker-build build
	$(DOCKER) /go/bin/plonk

ssh:
	$(DOCKER) bash

adr_init:
	$(ADR_DOCKER) adr init ./adrs/

adr:
ifdef ADR_TITLE
	$(ADR_DOCKER) adr new $(ADR_TITLE)
else
	@echo "ADR Title missing. Try like this: 'make adr ADR_TITLE=\"ADR Title here\"'"
	@exit
endif

super_adr:
ifdef ADR_TITLE OLD_ADR
	$(ADR_DOCKER) adr new -s $(OLD_ADR) $(ADR_TITLE)
else
	@echo "ADR Title missing as well as superseeded ADR number. Try the command like this: 'make super_adr ADR_TITLE'\"New ADR title\" OLD_ADR=<Number of the ADR to be superseded>"
	@exit
endif

