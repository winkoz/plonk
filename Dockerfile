FROM golang:1.15.6

ENV GOPATH "/tmp/go"
RUN mkdir -p $GOPATH

RUN go get -u gotest.tools/gotestsum
RUN go get -u github.com/go-bindata/go-bindata/...

ENV PATH="${PATH}:${GOPATH}/bin"

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

CMD ["tail", "-f", "/dev/null"]
