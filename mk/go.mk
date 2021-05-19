.PHONY: go-build go-test go-build-assets
GOOS?=darwin
GOARCH?=amd64
BINARY_NAME?=plonk

go-build-assets:
	GO111MODULE=on $(GO-BIN-FOLDER)go-bindata -prefix "data/" -pkg data -o data/data.go data/...

go-build: go-build-assets
	GO111MODULE=on GOOS=$(GOOS) GOARCH=$(GOARCH) go build -mod=mod -ldflags="-s -w" -o bin/$(BINARY_NAME) main.go
	@echo "Plonk built successfully!"

go-test: go-build-assets
	GO111MODULE=on $(GO-BIN-FOLDER)gotestsum --junitfile unit-tests.xml --format pkgname-and-test-fails -- -mod=mod -cover ./...
	@echo "Plonk finished testing!"
