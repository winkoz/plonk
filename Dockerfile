FROM golang:1.19.3

ENV PATH="${PATH}:${GOPATH}/bin"

WORKDIR /tmp/go

# Copy go mod and sum files
COPY go.mod go.sum ./
RUN go install gotest.tools/gotestsum@latest
RUN go install github.com/go-bindata/go-bindata/...@latest

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

CMD ["tail", "-f", "/dev/null"]
